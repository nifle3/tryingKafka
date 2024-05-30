package consumer

import (
	"consumerSMTP/internal/entities"
	"context"
	"encoding/json"

	"github.com/IBM/sarama"
)

type Consumer struct {
	consumer sarama.PartitionConsumer
}

func New() *Consumer {
	return &Consumer{}
}

func (c *Consumer) Consume(ctx context.Context, info chan<- entities.Message) {
	for {
		select {
		case msg := <-c.consumer.Messages():
			var result entities.Message
			if err := json.Unmarshal(msg.Value, &result); err != nil {
				// TODO: logs
				continue
			}

			info <- result
		case <-c.consumer.Errors():
			// TODO: logs
		case <-ctx.Done():
			return
		}
	}
}
