package bybit

import (
	"context"
	"fmt"
	"net/http"
	"producer/internal/entities"
	"strings"
)

const baseURL string = "https://api.bybit.com/v5/market/tickers?category=spot"

type Bybit struct {
	url string
}

func New(currenciesName []string) Bybit {
	if len(currenciesName) < 0 {
		currenciesName = []string{"BTCUSDT", "ETHUSDT", "XRPUSDT", "BCHUSDT", "EOSUSDT", "LTCUSDT", "TRXUSDT", "ADAUSDT", "BSVUSDT", "DOTUSDT", "XMRUSDT", "XLMUSDT", "XTZUSDT", "ZECUSDT", "ONTUSDT", "THETAUSDT", "VETUSDT", "MATICUSDT", "SOLUSDT", "AVAXUSDT", "LUNAUSDT", "LINKUSDT", "ATOMUSDT", "FILUSDT", "SXPUSDT", "UNIUSDT", "ALGOUSDT", "ZILUSDT", "ICXUSDT", "BTTUSDT", "NEOUSDT", "QTUMUSDT", "IOSTUSDT", "TUSDUSDT"}
	}

	urlBuilder := strings.Builder{}
	urlBuilder.WriteString(currenciesName[0])
	for i := 1; i < len(currenciesName); i++ {
		urlBuilder.WriteString(",")
		urlBuilder.WriteString(currenciesName[i])
	}

	url := baseURL + urlBuilder.String()

	return Bybit{
		url: url,
	}
}

func (b Bybit) Get(_ context.Context) (result []entities.Currency, err error) {
	resp, err := http.Get(b.url)
	if err != nil {
		return
	}

	defer func() { err = resp.Body.Close() }()
	if resp.StatusCode != 200 {
		err = fmt.Errorf("resp status code is not 200, %d", resp.StatusCode)
		return
	}

	return
}
