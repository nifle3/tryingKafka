package cfg

import (
    "fmt"
    "log/slog"
    "os"
    "strconv"
)

type Cfg struct {
    Kafka
    SMTP
}

type SMTP struct {
    Username string
    Password string
    Host     string
    Addr     string
}

type Kafka struct {
    Broker    string
    Topic     string
    Partition int
}

func New() Cfg {
    broker, ok := os.LookupEnv("KAFKA_BROKER")
    if !ok {
        broker = "localhost:9092"
        slog.Warn(fmt.Sprintf("KAFKA_BROKER is not set. Apply default value %s", broker))
    }

    var topicName string
    topicName, ok = os.LookupEnv("KAFKA_TOPIC")
    if !ok {
        topicName = "test"
        slog.Warn(fmt.Sprintf("KAFKA_TOPIC is not set. Apply default value %s", topicName))
    }

    var partitionString string
    partitionString, ok = os.LookupEnv("KAFKA_PARTITION")
    if !ok {
        partitionString = "0"
        slog.Warn(fmt.Sprintf("KAFKA_PARTITION is not set. Apply default value %s", partitionString))
    }

    partition, err := strconv.Atoi(partitionString)
    if err != nil {
        partition = 0
        slog.Warn(fmt.Sprintf("KAFKA_PARTITION cannot be casts to integer. Apply default value %d", partition))
    }

    var username string
    username, ok = os.LookupEnv("SMTP_USERNAME")
    if !ok {
        username = "smtp@gmail.com"
        slog.Warn(fmt.Sprintf("SMTP_USERNAME is not set. Apply default value %s", username))
    }

    var password string
    password, ok = os.LookupEnv("SMTP_PASSWORD")
    if !ok {
        password = "123"
        slog.Warn(fmt.Sprintf("SMTP_PASSWORD is not set. Apply default value %s", password))
    }

    var host string
    host, ok = os.LookupEnv("SMTP_HOST")
    if !ok {
        host = "smtp.gmail.com"
        slog.Warn(fmt.Sprintf("SMTP_HOST is not set. Apply default value %s", host))
    }

    var addr string
    addr, ok = os.LookupEnv("SMTP_ADDR")
    if !ok {
        addr = "smtp.gmail.com:587"
        slog.Warn(fmt.Sprintf("SMTP_ADDR is not set. Apply default value %s", addr))
    }

    return Cfg{
        Kafka: Kafka{
            Broker: broker,
            Topic:  topicName,
        },
        SMTP: SMTP{
            Username: username,
            Password: password,
            Host:     host,
            Addr:     addr,
        },
    }
}

func parseConfigValue(name string, defaultValue string) string {
    value, ok := os.LookupEnv(name)
    if !ok {
        slog.Warn(fmt.Sprintf("%s is not set. Apply default value %s", name, defaultValue))
        return defaultValue
    }

    return value
}
