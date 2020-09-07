package client

import (
	"coins/common"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const maxLimit = "5000"

type PriceListClientI interface {
	 GetUsdPrices() (*common.PriceInfo, error)
}

type PriceListClient struct {
	Client common.HttpClient
	url string
}

func NewPriceListClient(c *http.Client, url string) *PriceListClient {
	return &PriceListClient{
		Client: c,
		url : url,
	}
}

// see https://coinmarketcap.com/api/documentation/v1/#operation/getV1CryptocurrencyListingsLatest
func (c PriceListClient) GetUsdPrices() (*common.PriceInfo, error) {

	req, err := http.NewRequest(http.MethodGet, c.url, nil)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Add("limit", maxLimit)
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c")
	req.URL.RawQuery = q.Encode()

	resp, err := c.Client.Do(req);
	if err != nil {
		return nil, err
	}
	respBody, _ := ioutil.ReadAll(resp.Body)

	response := &common.PriceInfo{}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, err
	}

	return response, nil
}