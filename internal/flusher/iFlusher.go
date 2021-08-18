package flusher

import (
	"github.com/ozonva/ova-food-api/internal/repo"
	food "github.com/ozonva/ova-food-api/pkg/food"
)

// Flusher - интерфейс для сброса задач в хранилище
type Flusher interface {
	Flush(foods []food.Food) []food.Food
}

// NewFlusher возвращает Flusher с поддержкой батчевого сохранения
func NewFlusher(
	chunkSize int,
	foodRepo repo.Repo,
) Flusher {
	return &flusher{
		chunkSize: chunkSize,
		foodRepo:  foodRepo,
	}
}
