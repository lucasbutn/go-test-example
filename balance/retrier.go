package balance

//go:generate mockgen -destination=../balance/mocks/mock_retrier.go -package=mocks -source=retrier.go Retrier

import (
	"time"
)

type Retrier interface {
	Run(f func() error) error
}

type delayedRetrier struct {
	retries, delay int
}

func (n *delayedRetrier) Run(f func() error) error {
	err := f()
	for retry := 0; retry < n.retries && err != nil; retry++ {
		time.Sleep(time.Duration(n.delay) * time.Millisecond)
		err = f()
	}
	return err
}

func NewDelayedRetrier(retries, delay int) Retrier {
	return &delayedRetrier{retries, delay}
}
