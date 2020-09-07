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
		resp *common.RankInfo
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
				client: common.NewMockClientWithRankInfo(&common.RankInfo{
					Data: []common.RankData{
						{
							CoinInfo: common.CoinInfo{
								Name:"ETH",
							},
						},
						{
							CoinInfo: common.CoinInfo{
								Name:"BTC",
							},
						},
					},
				}),
			},
			resp: &common.RankInfo{
				Data: []common.RankData{
					{
						CoinInfo: common.CoinInfo{
							Name:"ETH",
						},
					},
					{
						CoinInfo: common.CoinInfo{
							Name:"BTC",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rlC := &RankListClient{
				Client: tt.fields.client,
			}
			resp, err := rlC.GetRankList("50")
			if err != nil {
				assert.True(t, tt.err)
			} else {
				assert.True(t, reflect.DeepEqual(resp, tt.resp))
			}
		})
	}
}
