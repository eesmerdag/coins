package common

type InfoCombine struct {
	Data []CombinedData `json:"data"`
}

type CombinedData struct {
	Rank   int     `json:"rank"`
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
}
