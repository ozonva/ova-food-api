package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-food-api/internal/food"
)

// Repo - интерфейс хранилища для сущности Food
type Repo interface {
	AddEntities(entities []food.Food) error
	AddEntity(entity food.Food) error
	ListEntities(limit, offset uint64) ([]food.Food, error)
	DescribeEntity(entityId uint64) (*food.Food, error)
	RemoveEntity(entityId uint64) error
}

func NewRepo(database *sqlx.DB) Repo {
	return &repoPostgres{db: *database}
}
