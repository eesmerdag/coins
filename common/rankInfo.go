package common

type RankInfo struct {
	Data []RankData `json:"Data"`
}

type RankData struct {
	CoinInfo CoinInfo `json:"CoinInfo"`
}

type CoinInfo struct {
	Name string `json:"Name"`
}
