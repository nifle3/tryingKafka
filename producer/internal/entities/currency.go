package entities

import "time"

type Currency struct {
	TimeNow     time.Time
	Symbol      string `json:"symbol"`
	OpenPrice   string `json:"openPrice"`
	HighPrice   string `json:"highPrice"`
	LowPrice    string `json:"lowPrice"`
	LastPrice   string `json:"lastPrice"`
	Volume      string `json:"volume"`
	QuoteVolume string `json:"quoteVolume"`
	Count       int    `json:"count"`
}
