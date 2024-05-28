package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"producer/internal/entities"
)

const baseURL string = "https://api.binance.com/api/v3/ticker/tradingDay?symbols="

type Binance struct {
	url string
}

func New() {

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

	err = json.NewDecoder(resp.Body).Decode(&result)
	log.Printf("%#v", result)
	return
}
