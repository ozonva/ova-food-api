package saver_test

import (
	"context"
	"errors"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/ozonva/ova-food-api/internal/flusher"
	"github.com/ozonva/ova-food-api/internal/food"
	"github.com/ozonva/ova-food-api/internal/mocks"
	"github.com/ozonva/ova-food-api/internal/saver"
)

var _ = Describe("Saver check init", func() {
	var (
		ctx           context.Context
		coffee        = food.Food{Id: 0, UserId: 0, Type: food.Drinks, Name: "Coffee", PortionSize: 60}
		chunkSize     int
		ctrl          *gomock.Controller
		mockRepo      *mocks.MockRepo
		flusherEntity flusher.Flusher
		saverEntity   saver.Saver
		capacity      uint
	)

	BeforeEach(func() {
		ctx = context.Background()
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
		flusherEntity = flusher.NewFlusher(chunkSize, mockRepo)
		saverEntity = saver.NewSaver(capacity, flusherEntity)
		saverEntity.Init(ctx)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("saver entity save foods, init ok", func() {
		It("saver entity save foods, init ok", func() {
			chunkSize = 2
			capacity = 2
			gomega.Expect(func() {
				saverEntity.Save(ctx, coffee)
			}).ShouldNot(gomega.Panic())
			saverEntity.Close(ctx)
		})
	})
})
var _ = Describe("Saver internal error", func() {
	var (
		ctx    context.Context
		coffee = food.Food{Id: 0, UserId: 0, Type: food.Drinks, Name: "Coffee", PortionSize: 60}
		pizza  = food.Food{Id: 1, UserId: 0, Type: food.Foods, Name: "Pizza", PortionSize: 300}

		chunkSize     int
		capacity      uint
		ctrl          *gomock.Controller
		mockRepo      *mocks.MockRepo
		flusherEntity flusher.Flusher
		saverEntity   saver.Saver
	)

	BeforeEach(func() {
		ctx = context.Background()
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
	})
	JustBeforeEach(func() {
		flusherEntity = flusher.NewFlusher(chunkSize, mockRepo)
		saverEntity = saver.NewSaver(capacity, flusherEntity)
		saverEntity.Init(ctx)
	})
	AfterEach(func() {
		ctrl.Finish()
	})

	Context("repo not save foods, internal error", func() {
		BeforeEach(func() {
			chunkSize = 1
			capacity = 1
			mockRepo.EXPECT().AddEntities(ctx, gomock.Any()).Return(errors.New("some internal error"))
		})
		It("repo not save foods, internal error", func() {

			gomega.Expect(func() {
				saverEntity.Save(ctx, coffee)
				saverEntity.Save(ctx, pizza)
			}).Should(gomega.PanicWith(errors.New("Internal repo error, cant save")))
		})
	})
})
var _ = Describe("Saver save data", func() {
	var (
		coffee        = food.Food{Id: 0, UserId: 0, Type: food.Drinks, Name: "Coffee", PortionSize: 60}
		ctx           context.Context
		chunkSize     int
		capacity      uint
		ctrl          *gomock.Controller
		mockRepo      *mocks.MockRepo
		flusherEntity flusher.Flusher
		saverEntity   saver.Saver
	)

	BeforeEach(func() {
		ctx = context.Background()
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
	})
	JustBeforeEach(func() {
		flusherEntity = flusher.NewFlusher(chunkSize, mockRepo)
		saverEntity = saver.NewSaver(capacity, flusherEntity)
		saverEntity.Init(ctx)
	})
	AfterEach(func() {
		ctrl.Finish()
	})

	Context("repo save foods by ticker", func() {
		BeforeEach(func() {
			chunkSize = 1
			capacity = 1
			mockRepo.EXPECT().AddEntities(ctx, gomock.Any()).Return(nil).Times(2)
		})
		It("repo save foods by ticker", func() {
			gomega.Expect(func() {
				saverEntity.Save(ctx, coffee) //AddEntities() 1
				time.Sleep(1500 * time.Millisecond)
			}).ShouldNot(gomega.Panic())
			saverEntity.Close(ctx) //AddEntities() 2
		})
	})
	Context("repo save foods by .Save()", func() {
		BeforeEach(func() {
			chunkSize = 1
			capacity = 1
			mockRepo.EXPECT().AddEntities(ctx, gomock.Any()).Return(nil).Times(3)
		})
		It("repo save foods by .Save()", func() {
			gomega.Expect(func() {
				saverEntity.Save(ctx, coffee) //AddEntities() 1
				saverEntity.Save(ctx, coffee) //AddEntities() 2
				saverEntity.Save(ctx, coffee)
			}).ShouldNot(gomega.Panic())
			saverEntity.Close(ctx) //AddEntities() 3
		})
	})
})
