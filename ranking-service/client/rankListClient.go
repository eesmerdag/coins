package client

import (
	"coins/common"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type RankListClientI interface {
	GetRankList(limit string) (*common.RankInfo, error)
}

type RankListClient struct {
	Client common.HttpClient
	url    string
}

func NewRankListClient(c *http.Client, url string) *RankListClient {
	return &RankListClient{
		Client: c,
		url:    url,
	}
}

func (rlC RankListClient) GetRankList(limit string) (*common.RankInfo, error) {
	req, err := http.NewRequest(http.MethodGet, rlC.url, nil)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Add("limit", limit)
	q.Add("tsym", "USD")

	req.Header.Set("Accepts", "application/json")
	req.URL.RawQuery = q.Encode()

	resp, err := rlC.Client.Do(req);
	if err != nil {
		return nil, err
	}
	respBody, _ := ioutil.ReadAll(resp.Body)

	response := &common.RankInfo{}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, err
	}
	return response, nil
}
