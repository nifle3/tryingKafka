package kafka

import (
	"context"
	"encoding/json"

	"producer/internal/entities"

	"github.com/IBM/sarama"
)

const (
	topicName = ""
	partiion  = 0
)

type Kafka struct {
	producer sarama.SyncProducer
}

func New(addrs ...string) *Kafka {
	producer, err := sarama.NewSyncProducer(addrs, nil)
	if err != nil {
		return nil
	}

	return &Kafka{
		producer: producer,
	}
}

func (k *Kafka) Close() error {
	return k.producer.Close()
}

func (k *Kafka) Send(ctx context.Context, msg entities.Message) error {
	result, err := json.Marshal(msg.Currencies)
	if err != nil {
		return err
	}

	kafkaMsg := &sarama.ProducerMessage{
		Topic:     topicName,
		Partition: partiion,
		Key:       sarama.StringEncoder(msg.Id),
		Value:     sarama.ByteEncoder(result),
	}

	_, _, err = k.producer.SendMessage(kafkaMsg)
	return err
}
