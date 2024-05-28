package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"producer/internal/cfg"
	"producer/internal/entities"
	"producer/internal/sender"
)

func Start() {
	config := cfg.New()
	binance := sender.New(config.Timeout, config.CurrencyNames)
	info := make(chan []entities.Currency)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	go binance.Start(ctx, info)

	// producer := kafka.New()
	// producer.Start(ctx, infoChan)

}
