package app

import (
	"context"
	"net/smtp"
	"os"
	"os/signal"
	"syscall"

	"consumerSMTP/internal/cfg"
	"consumerSMTP/internal/consumer"
	"consumerSMTP/internal/entities"
	"consumerSMTP/internal/sender"
)

func Run() {
	config := cfg.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	info := make(chan entities.Message)
	consumerImplementation := consumer.New(config.Topic, config.Partition, config.Broker)
	defer consumerImplementation.Close()

	auth := smtp.PlainAuth("", config.Username, config.Password, config.Host)
	senderImplementation := sender.New(auth, config.SMTP.Addr, config.Username)

	go consumerImplementation.Consume(ctx, info)

	senderImplementation.Start(ctx, info)
}
