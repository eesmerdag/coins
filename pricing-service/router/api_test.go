package router

import (
	"coins/pricing-service/client"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPriceLister(t *testing.T) {
	tests := []struct {
		name            string
		priceListClient client.PriceListClientI
		req             *http.Request
		code            int
	}{
		{
			name: "not allowed method",
			req: &http.Request{
				Method: http.MethodPost,
			},
			code: http.StatusMethodNotAllowed,
		},
		{
			name: "error response",
			req: &http.Request{
				Method: http.MethodGet,
			},
			priceListClient: client.NewMockPriceListClientErrorCase(),
			code:            http.StatusInternalServerError,
		},
		{
			name: "success response",
			req: &http.Request{
				Method: http.MethodGet,
			},
			priceListClient: client.NewMockPriceListClientSuccessCase(),
			code:            http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			router, _ := NewRouter(&tt.priceListClient)
			res := http.HandlerFunc(router.ListPrices)
			rr := httptest.NewRecorder()
			res.ServeHTTP(rr, tt.req)
			assert.True(t, rr.Code == tt.code)
		})
	}
}
