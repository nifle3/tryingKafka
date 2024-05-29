package bybit

type Response struct {
	Result `json:"result"`
}

type Result struct {
	Result []Currency `json:"list"`
}

type Currency struct {
	Symbol       string `json:"symbol"`
	Ask1Price    string `json:"ask1Price"`
	LastPrice    string `json:"lastPrice"`
	HighPrice24H string `json:"highPrice24h"`
	LowPrice24H  string `json:"lowPrice24h"`
	Volume24H    string `json:"volume24h"`
}
