package producer

import (
	"context"

	"producer/internal/entities"
)

type Sender interface {
	Send(ctx context.Context, currencies entities.Message) error
}

type Producer struct {
	sender Sender
}

func New(sender Sender) *Producer {
	return &Producer{
		sender: sender,
	}
}

func (p *Producer) Start(ctx context.Context, info <-chan []entities.Currency) {
	for {
		select {
		case result := <-info:
			if err := p.sender.Send(ctx, result); err != nil {
				continue
			}
		case <-ctx.Done():
			return
		}
	}
}
