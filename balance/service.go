package balance

//go:generate mockgen -destination=../balance/mock_service.go -package=balance -source=service.go Service
import (
	"errors"
)

type Service interface {
	GetBalance(userId string) (*Balance, error)
}

type service struct {
	client  Client
	retrier Retrier
}

func (s *service) GetBalance(userId string) (*Balance, error) {

	var movements []*Movement
	var err error

	err = s.retrier.Run(func() error {
		movements, err = s.client.GetAllMovements(userId)
		return err
	})

	if err != nil {
		return nil, err
	}

	total := getTotal(movements)

	return &Balance{userId, total}, nil
}

func getTotal(movements []*Movement) float64 {
	var sum float64
	for _, movement := range movements {
		sum += movement.Value
	}
	return sum
}

func NewService(rc Client, retrier Retrier) (Service, error) {
	if rc == nil {
		return nil, errors.New("rest client cant be nil")
	}
	if retrier == nil {
		return nil, errors.New("retrier cant be nil")
	}
	return &service{rc, retrier}, nil
}
