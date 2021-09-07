package saver

import (
	"context"

	"github.com/ozonva/ova-food-api/internal/flusher"
	"github.com/ozonva/ova-food-api/internal/food"
)

type Saver interface {
	Save(ctx context.Context, food food.Food) error
	Init(ctx context.Context)
	Close(ctx context.Context) error
}

// NewSaver возвращает Saver с поддержкой переодического сохранения
func NewSaver(capacity uint, flusher flusher.Flusher) Saver {
	s := &saver{
		flusher: flusher,
		data:    make([]food.Food, 0, capacity),
	}
	return s
}
