package sender

import (
	"context"
	"time"

	"producer/internal/entities"
)

type Exchange interface {
	Get(ctx context.Context) ([]entities.Currency, error)
}

type Sender struct {
	timeout   time.Duration
	exchanges []Exchange
}

func New(timeout time.Duration, exchanges ...Exchange) *Sender {
	return &Sender{
		timeout:   timeout,
		exchanges: exchanges,
	}
}

func (s Sender) Start(ctx context.Context, infoChan chan<- []entities.Currency) {
	for {
		select {
		case <-time.After(s.timeout):
			result := s.getCurrency(ctx)

			infoChan <- result
		case <-ctx.Done():
			return
		}
	}
}

func (s Sender) getCurrency(ctx context.Context) []entities.Currency {
	result := make([]entities.Currency, 0)

	for _, exchange := range s.exchanges {
		currencies, err := exchange.Get(ctx)
		if err != nil {
			continue
		}

		result = append(result, currencies...)
	}

	return result
}
