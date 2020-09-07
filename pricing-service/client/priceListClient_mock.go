package client

import (
	"coins/common"
	"errors"
)

type MockPriceListClient struct {
	GetUsdPricesFunc func() (*common.PriceInfo, error)
}

func (m MockPriceListClient) GetUsdPrices() (*common.PriceInfo, error) {
	if m.GetUsdPricesFunc != nil {
		return m.GetUsdPricesFunc()
	}
	return &common.PriceInfo{}, nil
}

func NewMockPriceListClientErrorCase() *MockPriceListClient {
	return &MockPriceListClient{
		GetUsdPricesFunc: func() (*common.PriceInfo, error) {
			return nil, errors.New("dummy error")
		},
	}
}

func NewMockPriceListClientSuccessCase() *MockPriceListClient {
	return &MockPriceListClient{
		GetUsdPricesFunc: func() (*common.PriceInfo, error) {
			return &common.PriceInfo{}, nil
		},
	}
}
