package flusher

import (
	food "github.com/ozonva/ova-food-api/internal/food"
	"github.com/ozonva/ova-food-api/internal/repo"
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
