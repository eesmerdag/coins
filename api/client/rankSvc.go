package client

import (
	"coins/common"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type RankSvcClientI interface {
	GetRankingList(limit string) (*common.SvcResponse, error)
}

type RankSvcClient struct {
	Client common.HttpClient
	url    string
}

func NewRankSvcClient(c *http.Client, url string) *RankSvcClient {
	return &RankSvcClient{
		Client: c,
		url:    url,
	}
}

func (svc RankSvcClient) GetRankingList(limit string) (*common.SvcResponse, error) {
	req, err := http.NewRequest(http.MethodGet, svc.url, nil)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Add("limit", limit)
	req.Header.Set("Accepts", "application/json")
	req.URL.RawQuery = q.Encode()

	resp, err := svc.Client.Do(req)
	svcResponse := &common.SvcResponse{}
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(body, &svcResponse)
		if err != nil {
			return nil, err
		}
		return svcResponse, nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &svcResponse)
	if err != nil {
		return nil, err
	}
	return svcResponse, nil
}
