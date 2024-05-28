package app

import (
	"context"
	"os"
	"os/signal"
	"producer/internal/api"
	"producer/internal/cfg"
	"producer/internal/entities"
	"syscall"
)

func Start() {
	config := cfg.New()
	binance := api.New(config.Timeout, config.CurrencyNames)
	info := make(chan []entities.Currency)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	go binance.Start(ctx, info)
	// producer := kafka.New()
	// producer.Start(infoChan)

}
