package consumer

import (
	"context"
	"encoding/json"

	"consumerSMTP/internal/entities"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func New(topicName string, partition int, brokers ...string) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   brokers,
		Topic:     topicName,
		Partition: partition,
		MaxBytes:  10e6, // 10MB
	})

	return &Consumer{
		reader: reader,
	}
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}

func (c *Consumer) Consume(ctx context.Context, info chan<- entities.Message) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			msg, err := c.reader.ReadMessage(context.Background())
			if err != nil {
				// TODO: logs
				continue
			}

			var result entities.Message
			if err := json.Unmarshal(msg.Value, &result); err != nil {
				// TODO: logs
				continue
			}

			info <- result
		}
	}
}
