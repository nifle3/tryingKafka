package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"producer/internal/entities"
)

const (
	baseURL    string = "https://api.binance.com/api/v3/ticker/tradingDay?symbols="
	marketName string = "Binance"
)

type Binance struct {
	url string
}

func New(currency []string) Binance {
	if len(currency) < 0 {
		return Binance{}
	}

	urlBuilder := strings.Builder{}
	urlBuilder.WriteString("[")
	urlBuilder.WriteString("\"")
	urlBuilder.WriteString(currency[0])
	urlBuilder.WriteString("\"")
	for i := 1; i < len(currency); i++ {
		tmp := fmt.Sprintf(",\"%s\"", currency[i])
		if _, err := urlBuilder.WriteString(tmp); err != nil {
			return Binance{}
		}
	}
	urlBuilder.WriteString("]")

	url := baseURL + urlBuilder.String()

	return Binance{
		url: url,
	}
}

func (b Binance) Get(_ context.Context) (result []entities.Currency, err error) {
	resp, err := http.Get(b.url)
	if err != nil {
		return
	}

	defer func() {
		err = resp.Body.Close()
	}()

	if resp.StatusCode != 200 {
		err = fmt.Errorf("status code is not 200. %d", resp.StatusCode)
		return
	}

	binanceResult := make([]Currency, 0)

	err = json.NewDecoder(resp.Body).Decode(&binanceResult)
	if err != nil {
		return
	}

	result = make([]entities.Currency, 0, len(binanceResult))
	for _, value := range binanceResult {
		tmp := entities.Currency{
			Market:      marketName,
			Symbol:      value.Symbol,
			OpenPrice:   value.OpenPrice,
			HighPrice:   value.HighPrice,
			LowPrice:    value.LowPrice,
			LastPrice:   value.LastPrice,
			Volume:      value.Volume,
			QuoteVolume: value.QuoteVolume,
			Count:       value.Count,
		}

		result = append(result, tmp)
	}

	log.Printf("%#v", result)
	return
}
