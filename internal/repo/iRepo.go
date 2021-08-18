package repo

import f "github.com/ozonva/ova-food-api/pkg/food"

// Repo - интерфейс хранилища для сущности Food
type Repo interface {
	AddEntities(entities []f.Food) error
	ListEntities(limit, offset uint64) ([]f.Food, error)
	DescribeEntity(entityId uint64) (*f.Food, error)
}

func NewRepo(filepath string) Repo {
	return &repoFile{
		filepath: filepath,
	}
}
