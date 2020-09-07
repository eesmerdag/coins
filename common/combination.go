package common

func CombineResults(ranks RankInfo, prices PriceInfo) *[]CombinedData {

	pricesMap := make(map[string]USD)
	for i := 0; i < len(prices.Data); i++ {
		pricesMap[prices.Data[i].Symbol] = prices.Data[i].QuoteInfo.Usd
	}

	combinedResults := make([]CombinedData, len(ranks.Data))
	for i := 0; i < len(ranks.Data); i++ {

		data := &CombinedData{
			Rank:   i,
			Symbol: ranks.Data[i].CoinInfo.Name,
			Price:  lookupForPrice(pricesMap, ranks.Data[i].CoinInfo.Name),
		}
		combinedResults[i] = *data
	}

	return &combinedResults
}

func lookupForPrice(pricesMap map[string]USD, name string) float64 {
	if value, ok := pricesMap[name]; ok {
		return value.Price
	}
	return -1
}
