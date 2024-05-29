package bybit

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"producer/internal/entities"
)

const (
	baseURL    string = "https://api.bybit.com/v5/market/tickers?category=spot"
	marketName string = "bybit"
)

type Bybit struct {
	url string
}

func New(currency []string) Bybit {
	if len(currency) < 0 {
		currency = []string{"BTCUSDT", "ETHUSDT", "XRPUSDT", "BCHUSDT", "EOSUSDT", "LTCUSDT", "TRXUSDT", "ADAUSDT", "BSVUSDT", "DOTUSDT", "XMRUSDT", "XLMUSDT", "XTZUSDT", "ZECUSDT", "ONTUSDT", "THETAUSDT", "VETUSDT", "MATICUSDT", "SOLUSDT", "AVAXUSDT", "LUNAUSDT", "LINKUSDT", "ATOMUSDT", "FILUSDT", "SXPUSDT", "UNIUSDT", "ALGOUSDT", "ZILUSDT", "ICXUSDT", "BTTUSDT", "NEOUSDT", "QTUMUSDT", "IOSTUSDT", "TUSDUSDT"}
	}

	urlBuilder := strings.Builder{}
	urlBuilder.WriteString(currency[0])
	for i := 1; i < len(currency); i++ {
		urlBuilder.WriteString(",")
		urlBuilder.WriteString(currency[i])
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

	var response Response
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return
	}

	result = make([]entities.Currency, 0, len(response.Result.Result))

	for _, value := range response.Result.Result {
		result = append(result, entities.Currency{
			Market:    marketName,
			Symbol:    value.Symbol,
			OpenPrice: value.Ask1Price,
			HighPrice: value.HighPrice24H,
			LowPrice:  value.LowPrice24H,
			LastPrice: value.LastPrice,
			Volume:    value.Volume24H,
		})
	}

	return
}
