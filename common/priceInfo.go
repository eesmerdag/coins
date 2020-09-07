package common

import (
	"time"
)

type PriceInfo struct {
	Status *Status     `json:"status"`
	Data   []PriceData `json:"data"`
}

type Status struct {
	ErrorCode    int       `json:"error_code"`
	ErrorMessage string    `json:"error_message"`
	Time         time.Time `json:"timestamp"`
}

type PriceData struct {
	Name      string `json:"name"`
	Symbol    string `json:"symbol"`
	QuoteInfo Quote  `json:"quote"`
}

type Quote struct {
	Usd USD `json:"USD"`
}

type USD struct {
	Price float64 `json:"price"`
}
