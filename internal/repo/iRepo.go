package repo

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-food-api/internal/food"
)

// Repo - интерфейс хранилища для сущности Food
type Repo interface {
	AddEntities(ctx context.Context, entities []food.Food) error
	AddEntity(ctx context.Context, entity food.Food) error
	ListEntities(ctx context.Context, limit, offset uint64) ([]food.Food, error)
	DescribeEntity(ctx context.Context, entityId uint64) (*food.Food, error)
	RemoveEntity(ctx context.Context, entityId uint64) error

	UpdateEntity(ctx context.Context, food food.Food) error
	MultiAddEntity(ctx context.Context, foods [][]food.Food) error
}

func NewRepo(database sqlx.DB) Repo {
	return &repoPostgres{db: database}
}
