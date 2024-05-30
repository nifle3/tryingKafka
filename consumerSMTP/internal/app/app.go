package app

import (
    "context"
    "os"
    "os/signal"
    "syscall"
)

func Run() {
    // config := cfg.New()
    // kafkaImplementation := kafka.New()

    ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
    defer cancel()

    // info := make(chan interface{})
    // consumerImplementation := consumer.New(kafkaImplementation)

    // go consumer.Start(ctx, info)

}
