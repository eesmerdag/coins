package router

import (
	"coins/api/client"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestPriceLister(t *testing.T) {
	tests := []struct {
		name           string
		priceSvcClient client.PriceSvcClientI
		rankSvcClient  client.RankSvcClientI
		req            *http.Request
		code           int
	}{
		{
			name: "not allowed method",
			req: &http.Request{
				Method: http.MethodPost,
			},
			priceSvcClient: client.NewMockPriceSvcClientErrorCase(),
			rankSvcClient:  client.NewMockRankSvcClientErrorCase(),
			code:           http.StatusMethodNotAllowed,
		},
		{
			name: "error response from both service",
			req: &http.Request{
				Method: http.MethodGet,
				URL: &url.URL{
					RawQuery: "limit=10",
				},
			},
			priceSvcClient: client.NewMockPriceSvcClientErrorCase(),
			rankSvcClient:  client.NewMockRankSvcClientErrorCase(),
			code:           http.StatusInternalServerError,
		},
		{
			name: "error from pricing service",
			req: &http.Request{
				Method: http.MethodGet,
				URL: &url.URL{
					RawQuery: "limit=10",
				},
			},
			priceSvcClient: client.NewMockPriceSvcClientErrorCase(),
			rankSvcClient:  client.NewMockRankSvcClientSuccessCase(),
			code:           http.StatusInternalServerError,
		},
		{
			name: "error from Dockerfile_ranking service",
			req: &http.Request{
				Method: http.MethodGet,
				URL: &url.URL{
					RawQuery: "limit=10",
				},
			},
			priceSvcClient: client.NewMockPriceSvcClientSuccessCase(),
			rankSvcClient:  client.NewMockRankSvcClientErrorCase(),
			code:           http.StatusInternalServerError,
		},
		{
			name: "success",
			req: &http.Request{
				Method: http.MethodGet,
				URL: &url.URL{
					RawQuery: "limit=10",
				},
			},
			priceSvcClient: client.NewMockPriceSvcClientSuccessCase(),
			rankSvcClient:  client.NewMockRankSvcClientSuccessCase(),
			code:           http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			router, _ := NewRouter(&tt.priceSvcClient, &tt.rankSvcClient)
			res := http.HandlerFunc(router.TopCryptoLister)
			rr := httptest.NewRecorder()
			res.ServeHTTP(rr, tt.req)
			assert.True(t, rr.Code == tt.code)
		})
	}
}
