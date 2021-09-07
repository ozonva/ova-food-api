package saver

import (
	"context"
	"errors"
	"time"

	"github.com/ozonva/ova-food-api/internal/logger"

	"github.com/ozonva/ova-food-api/internal/flusher"
	"github.com/ozonva/ova-food-api/internal/food"
)

type saver struct {
	flusher flusher.Flusher
	data    []food.Food
	ticker  *time.Ticker
	stop    chan struct{}
}

func (s *saver) Save(ctx context.Context, food food.Food) error {
	if len(s.data) >= cap(s.data) {
		err := s.flush(ctx)
		if err != nil {
			return err
		}
	}
	s.data = append(s.data, food)
	return nil
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
					err := s.flush(ctx)
					if err != nil {
						logger.GlobalLogger.Warn().Msg("internal db error while Ticker flush")
					}
				} else {
					logger.GlobalLogger.Warn().Msg("Ticker channel was closed")
				}
			case <-s.stop:
				break
			}

		}
	}()
}

func (s *saver) Close(ctx context.Context) error {
	err := s.flush(ctx)
	if err != nil {
		return err
	}
	s.stop <- struct{}{}
	close(s.stop)
	s.ticker.Stop()
	return nil
}

func (s *saver) flush(ctx context.Context) error {
	res := s.flusher.Flush(ctx, s.data)
	if res != nil {
		s.stop <- struct{}{}
		close(s.stop)
		s.ticker.Stop()
		logger.GlobalLogger.Warn().Msg("Internal repo error, cant save")
		return errors.New("Internal repo error, cant save")
	}
	s.data = s.data[:0]
	return nil
}
