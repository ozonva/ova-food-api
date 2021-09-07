package saver

import (
	"context"
	"errors"
	"time"

	"github.com/ozonva/ova-food-api/internal/flusher"
	"github.com/ozonva/ova-food-api/internal/food"
)

type saver struct {
	flusher flusher.Flusher
	data    []food.Food
	ticker  *time.Ticker
	stop    chan struct{}
}

func (s *saver) Save(ctx context.Context, food food.Food) {
	if len(s.data) == cap(s.data) {
		s.flush(ctx)
	}
	s.data = append(s.data, food)
}

func (s *saver) Init(ctx context.Context) {
	s.initTimerSaver(ctx, time.Second)
	s.stop = make(chan struct{})
}

func (s *saver) initTimerSaver(ctx context.Context, d time.Duration) {
	s.ticker = time.NewTicker(d)
	go func() {
		for {
			select {
			case _, ok := <-s.ticker.C:
				if ok {
					s.flush(ctx)
				} else {
					panic(errors.New("Ticker channel was closed"))
				}
			case <-s.stop:
				break
			}

		}
	}()
}

func (s *saver) Close(ctx context.Context) {
	s.flush(ctx)
	s.stop <- struct{}{}
	close(s.stop)
	s.ticker.Stop()
}

func (s *saver) flush(ctx context.Context) {
	res := s.flusher.Flush(ctx, s.data)
	if res != nil {
		s.stop <- struct{}{}
		close(s.stop)
		s.ticker.Stop()
		panic(errors.New("Internal repo error, cant save"))
	}
	s.data = s.data[:0]
}
