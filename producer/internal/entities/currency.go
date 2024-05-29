package entities

type Currency struct {
	Market    string `json:"market"`
	Symbol    string `json:"symbol"`
	OpenPrice string `json:"openPrice"`
	HighPrice string `json:"highPrice"`
	LowPrice  string `json:"lowPrice"`
	LastPrice string `json:"lastPrice"`
	Volume    string `json:"volume"`
	Count     int    `json:"count"`
}
