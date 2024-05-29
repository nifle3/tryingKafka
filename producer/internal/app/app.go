package app

import (
	"context"
	"os"
	"os/signal"
	"producer/internal/api/bybit"
	"producer/internal/kafka"
	"producer/internal/producer"
	"syscall"

	"producer/internal/api/binance"
	"producer/internal/cfg"
	"producer/internal/entities"
	"producer/internal/sender"
)

func Start() {
	config := cfg.New()

	binanceAPI := binance.New(config.CurrencyNames)
	bybitAPI := bybit.New(config.CurrencyNames)

	api := sender.New(config.Timeout, binanceAPI, bybitAPI)
	info := make(chan entities.Message)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	go api.Start(ctx, info)

	kaf := kafka.New()
	defer kaf.Close()

	prod := producer.New(kaf)
	prod.Start(ctx, info)
}
