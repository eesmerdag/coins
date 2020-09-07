package router

import (
	"coins/ranking-service/client"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestRankLister(t *testing.T) {
	tests := []struct {
		name           string
		rankListClient client.RankListClientI
		req            *http.Request
		code           int
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
			rankListClient: client.NewMockRankListClientErrorCase(),
			code:           http.StatusInternalServerError,
		},
		{
			name: "bad request",
			req: &http.Request{
				Method: http.MethodGet,
				URL: &url.URL{
					RawQuery: "limit=test",
				},
			},
			rankListClient: client.NewMockRankListClientSuccessCase(),
			code:           http.StatusBadRequest,
		},
		{
			name: "bad request",
			req: &http.Request{
				Method: http.MethodGet,
				URL: &url.URL{
					RawQuery: "",
				},
			},
			rankListClient: client.NewMockRankListClientSuccessCase(),
			code:           http.StatusBadRequest,
		},
		{
			name: "bad request",
			req: &http.Request{
				Method: http.MethodGet,
				URL: &url.URL{
					RawQuery: "limit=9",
				},
			},
			rankListClient: client.NewMockRankListClientSuccessCase(),
			code:           http.StatusBadRequest,
		},
		{
			name: "bad request",
			req: &http.Request{
				Method: http.MethodGet,
				URL: &url.URL{
					RawQuery: "limit=101",
				},
			},
			rankListClient: client.NewMockRankListClientSuccessCase(),
			code:           http.StatusBadRequest,
		},
		{
			name: "success response",
			req: &http.Request{
				Method: http.MethodGet,
				URL: &url.URL{
					RawQuery: "limit=50",
				},
			},
			rankListClient: client.NewMockRankListClientSuccessCase(),
			code:           http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router, _ := NewRouter(&tt.rankListClient)
			res := http.HandlerFunc(router.ListRanks)
			rr := httptest.NewRecorder()
			res.ServeHTTP(rr, tt.req)
			assert.True(t, rr.Code == tt.code)
		})
	}
}
