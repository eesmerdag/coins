package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombination(t *testing.T) {
	for _, tc := range []struct {
		name      string
		lookFor string
		expectedMoney float64
		rankInfo  *RankInfo
		priceInfo *PriceInfo
		isExist bool
	}{
		{
			name: "not found",
			lookFor: "ETH",
			expectedMoney: -1,
			rankInfo: &RankInfo{
				Data: []RankData{
					{
						CoinInfo: CoinInfo{
							Name: "ETH",
						},
					},
				},
			},
			priceInfo: &PriceInfo{
				Data: []PriceData{
					{
						Name:   "Litecoin",
						Symbol: "LTC",
						QuoteInfo: Quote{
							Usd: USD{
								Price: 333.11,
							},
						},
					},
					{
						Name:   "Bitcoin",
						Symbol: "BTC",
						QuoteInfo: Quote{
							Usd: USD{
								Price: 3.11,
							},
						},
					},
				},
			},
			isExist: false,
		},
		{
			name: "not found",
			lookFor: "LTC",
			expectedMoney: 333.11,
			rankInfo: &RankInfo{
				Data: []RankData{
					{
						CoinInfo: CoinInfo{
							Name: "LTC",
						},
					},
				},
			},
			priceInfo: &PriceInfo{
				Data: []PriceData{
					{
						Name:   "Litecoin",
						Symbol: "LTC",
						QuoteInfo: Quote{
							Usd: USD{
								Price: 333.11,
							},
						},
					},
					{
						Name:   "Bitcoin",
						Symbol: "BTC",
						QuoteInfo: Quote{
							Usd: USD{
								Price: 3.11,
							},
						},
					},
				},
			},
			isExist: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			result := *CombineResults(*tc.rankInfo, *tc.priceInfo)
			for i := 0; i < len(result); i++ {
				if result[i].Symbol == tc.lookFor {
					assert.Equal(t, result[i].Price, tc.expectedMoney)
					break
				}
			}
		})
	}
}
