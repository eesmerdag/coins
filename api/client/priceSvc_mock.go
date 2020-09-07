package client

import (
	"coins/common"
	"errors"
)

type MockPriceSvcClient struct {
	ListPricesFunc func() (*common.SvcResponse, error)
}

func (m MockPriceSvcClient) ListPrices() (*common.SvcResponse, error) {
	if m.ListPricesFunc != nil {
		return m.ListPricesFunc()
	}
	return &common.SvcResponse{}, nil
}

func NewMockPriceSvcClientErrorCase() *MockPriceSvcClient {
	return &MockPriceSvcClient{
		ListPricesFunc: func() (*common.SvcResponse, error) {
			return nil, errors.New("dummy error")
		},
	}
}

func NewMockPriceSvcClientSuccessCase() *MockPriceSvcClient {
	return &MockPriceSvcClient{
		ListPricesFunc: func() (*common.SvcResponse, error) {
			return &common.SvcResponse{
				PriceInfo: &common.PriceInfo{
					Data: []common.PriceData{
						{
							Name:   "Bitcoin",
							Symbol: "BTC",
							QuoteInfo: common.Quote{
								Usd: common.USD{
									Price: 25.5,
								},
							},
						},
					},
				},
			}, nil
		},
	}
}
