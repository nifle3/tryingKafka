package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"producer/internal/entities"
)

const baseURL string = "https://api.binance.com/api/v3/ticker/tradingDay?symbols="

type API struct {
	timeout time.Duration
	URL     string
}

func New(timeout time.Duration, currencyName string) *API {
	URLWithCurrenct := baseURL + currencyName
	return &API{
		timeout: timeout,
		URL:     URLWithCurrenct,
	}
}

func (a API) Start(ctx context.Context, infoChan chan<- []entities.Currency) {
	for {
		select {
		case <-time.After(a.timeout):
			result, err := a.getCurrency(ctx)
			if err != nil {
				continue
			}

			infoChan <- result
		case <-ctx.Done():
			return
		}
	}
}

func (a API) getCurrency(_ context.Context) (result []entities.Currency, err error) {
	resp, err := http.Get(a.URL)
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

	err = json.NewDecoder(resp.Body).Decode(&result)
	return
}
