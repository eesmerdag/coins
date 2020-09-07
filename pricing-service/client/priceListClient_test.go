package client

import (
	"coins/common"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestPriceLister(t *testing.T) {
	type fields struct {
		client common.HttpClient
	}
	tests := []struct {
		name   string
		fields fields
		resp   *common.PriceInfo
		err    bool
	}{
		{
			name: "error response",
			fields: fields{
				client: common.NewMockClientErrorCase(),
			},
			err: true,
		},
		{
			name: "success response",
			fields: fields{
				client: common.NewMockClientWithPriceInfo(&common.PriceInfo{
					Data: []common.PriceData{
						{
							Name:   "Bitcoin",
							Symbol: "BTC",
							QuoteInfo: common.Quote{
								Usd: common.USD{
									Price: 333.11,
								},
							},
						},
						{
							Name:   "Ethereum",
							Symbol: "ETH",
							QuoteInfo: common.Quote{
								Usd: common.USD{
									Price: 2222.4444,
								},
							},
						},
					},
				}),
			},
			resp: &common.PriceInfo{
				Data: []common.PriceData{
					{
						Name:   "Bitcoin",
						Symbol: "BTC",
						QuoteInfo: common.Quote{
							Usd: common.USD{
								Price: 333.11,
							},
						},
					},
					{
						Name:   "Ethereum",
						Symbol: "ETH",
						QuoteInfo: common.Quote{
							Usd: common.USD{
								Price: 2222.4444,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plC := &PriceListClient{
				Client: tt.fields.client,
			}
			resp, err := plC.GetUsdPrices();
			if err != nil {
				assert.True(t, tt.err)
			} else {
				assert.True(t, reflect.DeepEqual(resp, tt.resp))
			}
		})
	}
}
