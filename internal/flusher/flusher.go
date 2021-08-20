package flusher

import (
	food "github.com/ozonva/ova-food-api/internal/food"
	"github.com/ozonva/ova-food-api/internal/repo"
	"github.com/ozonva/ova-food-api/internal/utils"
)

type flusher struct {
	chunkSize int
	foodRepo  repo.Repo
}

func (f *flusher) Flush(foodsIn []food.Food) []food.Food {
	splittedFoods := utils.SplitToBulks(foodsIn, f.chunkSize)
	var errFoods []food.Food
	for _, foodSlice := range splittedFoods {
		err := f.foodRepo.AddEntities(foodSlice)
		if err != nil {
			errFoods = append(errFoods, foodSlice...)
		}
	}
	return errFoods
}
