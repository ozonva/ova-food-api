package saver

import (
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

func (s *saver) Save(food food.Food) {
	if len(s.data) == cap(s.data) {
		s.flush()
	}
	s.data = append(s.data, food)
}

func (s *saver) Init() {
	s.initTimerSaver(time.Second)
	s.stop = make(chan struct{})
}

func (s *saver) initTimerSaver(d time.Duration) {
	s.ticker = time.NewTicker(d)
	go func() {
		for {
			select {
			case _, ok := <-s.ticker.C:
				if ok {
					s.flush()
				} else {
					panic(errors.New("Ticker channel was closed"))
				}
			case <-s.stop:
				break
			}

		}
	}()
}

func (s *saver) Close() {
	s.flush()
	s.stop <- struct{}{}
	close(s.stop)
	s.ticker.Stop()
}

func (s *saver) flush() {
	res := s.flusher.Flush(s.data)
	if res != nil {
		s.stop <- struct{}{}
		close(s.stop)
		s.ticker.Stop()
		panic(errors.New("Internal repo error, cant save"))
	}
	s.data = s.data[:0]
}
