package client

import (
	"coins/common"
	"errors"
)

type MockRankSvcClient struct {
	GetRankingListFunc func(limit string) (*common.SvcResponse, error)
}

func (m MockRankSvcClient) GetRankingList(limit string) (*common.SvcResponse, error) {
	if m.GetRankingListFunc != nil {
		return m.GetRankingListFunc(limit)
	}
	return &common.SvcResponse{}, nil
}

func NewMockRankSvcClientErrorCase() *MockRankSvcClient {
	return &MockRankSvcClient{
		GetRankingListFunc: func(limit string) (*common.SvcResponse, error) {
			return nil, errors.New("dummy error")
		},
	}
}

func NewMockRankSvcClientSuccessCase() *MockRankSvcClient {
	return &MockRankSvcClient{
		GetRankingListFunc: func(limit string) (*common.SvcResponse, error) {
			return &common.SvcResponse{
				RankInfo: &common.RankInfo{
					Data: []common.RankData{
						{
							CoinInfo: common.CoinInfo{
								Name: "BTC",
							},
						},
					},
				},
			}, nil
		},
	}
}
