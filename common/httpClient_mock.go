package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	if m.DoFunc != nil {
		return m.DoFunc(req)
	}
	return &http.Response{}, nil
}

func NewMockClientWithPriceInfo(priceInfo *PriceInfo) *MockClient {
	return &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(priceInfo)
			if err != nil {
				panic(err)
			}
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(b)),
			}, nil
		},
	}
}

func NewMockClientErrorCase() *MockClient {
	return &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusBadRequest,
			}, errors.New("dummy error")
		},
	}
}

func NewMockClientWithRankInfo(priceInfo *RankInfo) *MockClient {
	return &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(priceInfo)
			if err != nil {
				panic(err)
			}
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBuffer(b)),
			}, nil
		},
	}
}
