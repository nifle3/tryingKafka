package cfg

import (
	"os"
	"strings"
	"time"
)

type Cfg struct {
	APICfg
}

type APICfg struct {
	CurrencyNames []string
	Timeout       time.Duration
}

func New() *Cfg {
	cur, ok := os.LookupEnv("CURRENCY")
	if !ok {
		cur = "BTCUSDT"
	}

	curs := strings.Split(cur, ",")

	timeout, ok := os.LookupEnv("TIMEOUT")
	if !ok {
		timeout = "10s"
	}

	parsedTimeout, err := time.ParseDuration(timeout)
	if err != nil {
		parsedTimeout = time.Hour
	}

	return &Cfg{
		APICfg: APICfg{
			CurrencyNames: curs,
			Timeout:       parsedTimeout,
		},
	}
}
