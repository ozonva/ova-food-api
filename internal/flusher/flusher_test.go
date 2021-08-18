package flusher_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/ozonva/ova-food-api/internal/flusher"
	"github.com/ozonva/ova-food-api/internal/mocks"
	f "github.com/ozonva/ova-food-api/pkg/food"
)

var _ = Describe("Flusher", func() {
	var (
		coffee = f.Food{Id: 0, UserId: 0, Type: f.Drinks, Name: "Coffee", PortionSize: 60}
		pizza  = f.Food{Id: 1, UserId: 0, Type: f.Foods, Name: "Pizza", PortionSize: 300}
		slice  = []f.Food{coffee, pizza}

		err       error
		chunkSize int
		ctrl      *gomock.Controller
		mockRepo  *mocks.MockRepo
		flush     flusher.Flusher
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)

	})
	JustBeforeEach(func() {
		flush = flusher.NewFlusher(chunkSize, mockRepo)
		flush.Flush(slice)
	})
	AfterEach(func() {
		ctrl.Finish()
	})

	Context("repo save foods, chunks less slice", func() {
		BeforeEach(func() {
			chunkSize = 1
			mockRepo.EXPECT().AddEntities([]f.Food{coffee}).Return(nil).Times(1)
			mockRepo.EXPECT().AddEntities([]f.Food{pizza}).Return(nil).Times(1)
		})
		It("repo save foods, chunks less slice", func() {
			gomega.Expect(err).Should(gomega.BeNil())
		})
	})

	Context("repo save foods, chunks more slice", func() {
		BeforeEach(func() {
			chunkSize = 3
			mockRepo.EXPECT().AddEntities(gomock.Any()).Return(nil).Times(1)
		})
		It("repo save foods, chunks more slice", func() {
			gomega.Expect(err).Should(gomega.BeNil())
		})
	})
	/*
		Context("repo not save foods, nil slice", func() {
			BeforeEach(func() {
				chunkSize = 4
				slice = nil //[]f.Food{{}, {}}
				mockRepo.EXPECT().AddEntities(gomock.Any()).Return().Times(0)
			})
			It("repo not save foods, nil slice", func() {
				gomega.Expect(err).Should(gomega.BeNil())
			})
		})*/
})
