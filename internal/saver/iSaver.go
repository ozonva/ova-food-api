package saver

import (
	"github.com/ozonva/ova-food-api/internal/flusher"
	"github.com/ozonva/ova-food-api/internal/food"
)

type Saver interface {
	Save(food food.Food)
	Init()
	Close()
}

// NewSaver возвращает Saver с поддержкой переодического сохранения
func NewSaver(capacity uint, flusher flusher.Flusher) Saver {
	s := &saver{
		flusher: flusher,
		data:    make([]food.Food, 0, capacity),
	}
	s.Init()
	return s
}
