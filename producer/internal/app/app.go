package app

import (
	"producer/internal/api"
	"producer/internal/cfg"
	"producer/internal/entities"
)

func Start() {
	config := cfg.New()
	api := api.New(config.Timeout, config.CurrencyNames)
	info := make(chan []entities.Currency)
	exit := make(chan interface{})
	go api.Start(info, exit)

	// producer := kafka.New()
	// producer.Start(infoChan)

}
