package flusher_test

import (
	"errors"

	"github.com/ozonva/ova-food-api/internal/food"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/ozonva/ova-food-api/internal/flusher"
	"github.com/ozonva/ova-food-api/internal/mocks"
)

var _ = Describe("Flusher", func() {
	var (
		coffee = food.Food{Id: 0, UserId: 0, Type: food.Drinks, Name: "Coffee", PortionSize: 60}
		pizza  = food.Food{Id: 1, UserId: 0, Type: food.Foods, Name: "Pizza", PortionSize: 300}
		slice  = []food.Food{coffee, pizza}

		chunkSize int
		ctrl      *gomock.Controller
		mockRepo  *mocks.MockRepo
		flush     flusher.Flusher
		result    []food.Food
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
	})
	JustBeforeEach(func() {
		flush = flusher.NewFlusher(chunkSize, mockRepo)
		result = flush.Flush(slice)
	})
	AfterEach(func() {
		ctrl.Finish()
	})

	Context("repo save foods, chunks less slice", func() {
		BeforeEach(func() {
			chunkSize = 1
			mockRepo.EXPECT().AddEntities([]food.Food{coffee}).Return(nil).Times(1)
			mockRepo.EXPECT().AddEntities([]food.Food{pizza}).Return(nil).Times(1)
		})
		It("repo save foods, chunks less slice", func() {
			gomega.Expect(result).Should(gomega.BeNil())
		})
	})

	Context("repo save foods, chunks more slice", func() {
		BeforeEach(func() {
			chunkSize = 3
			mockRepo.EXPECT().AddEntities(gomock.Any()).Return(nil).Times(1)
		})
		It("repo save foods, chunks more slice", func() {
			gomega.Expect(result).Should(gomega.BeNil())
		})
	})

	Context("repo not save foods, nil slice", func() {
		BeforeEach(func() {
			chunkSize = 4
			slice = nil
			mockRepo.EXPECT().AddEntities(gomock.Any()).Times(0)
		})
		It("repo not save foods, nil slice", func() {
			gomega.Expect(result).Should(gomega.BeNil())
		})
	})
	Context("repo not save foods, internal errors", func() {
		BeforeEach(func() {
			chunkSize = 2
			slice = []food.Food{coffee, pizza}
			mockRepo.EXPECT().AddEntities(gomock.Any()).Return(errors.New("some internal error")).Times(1)
		})
		It("repo not save foods, internal errors", func() {
			gomega.Expect(result).Should(gomega.BeEquivalentTo(slice))
		})
	})
})
