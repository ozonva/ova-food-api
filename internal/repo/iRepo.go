package repo

import (
	"github.com/ozonva/ova-food-api/internal/food"
)

// Repo - интерфейс хранилища для сущности Food
type Repo interface {
	AddEntities(entities []food.Food) error
	ListEntities(limit, offset uint64) ([]food.Food, error)
	DescribeEntity(entityId uint64) (*food.Food, error)
}

func NewRepo(filepath string) Repo {
	return &repoFile{
		filepath: filepath,
	}
}
