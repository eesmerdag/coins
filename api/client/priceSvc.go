package client

import (
	"coins/common"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type PriceSvcClientI interface {
	ListPrices() (*common.SvcResponse, error)
}

type PriceSvcClient struct {
	Client common.HttpClient
	url    string
}

func NewPriceListClient(c *http.Client, url string) *PriceSvcClient {
	return &PriceSvcClient{
		Client: c,
		url:    url,
	}
}

func (svc PriceSvcClient) ListPrices() (*common.SvcResponse, error) {
	req, err := http.NewRequest(http.MethodGet, svc.url, nil)
	if err != nil {
		return nil, err
	}

	res, err := svc.Client.Do(req)
	svcResponse := &common.SvcResponse{}
	if res.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(body, &svcResponse)
		if err != nil {
			return nil, err
		}
		return svcResponse, nil
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &svcResponse)
	if err != nil {
		return nil, err
	}
	return svcResponse, nil
}
