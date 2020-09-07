package client

import (
	"coins/common"
	"errors"
)

type MockRankListClient struct {
	GetRankListFunc func() (*common.RankInfo, error)
}

func (m MockRankListClient) GetRankList(limit string) (*common.RankInfo, error) {
	if m.GetRankListFunc != nil {
		return m.GetRankListFunc()
	}
	return &common.RankInfo{}, nil
}

func NewMockRankListClientErrorCase() *MockRankListClient {
	return &MockRankListClient{
		GetRankListFunc: func() (*common.RankInfo, error) {
			return nil, errors.New("dummy error")
		},
	}
}

func NewMockRankListClientSuccessCase() *MockRankListClient {
	return &MockRankListClient{
		GetRankListFunc: func() (*common.RankInfo, error) {
			return &common.RankInfo{}, nil
		},
	}
}
