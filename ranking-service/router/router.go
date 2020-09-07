package router

import (
	"coins/common"
	"coins/ranking-service/client"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	router         *mux.Router
	rankListClient client.RankListClientI
}

func NewRouter(rankListClient *client.RankListClientI) (*Router, error) {
	router := mux.NewRouter()

	r := &Router{
		router:         router,
		rankListClient: *rankListClient,
	}

	router.HandleFunc("/coin-ranks", r.ListRanks).Methods(http.MethodGet)

	return r, nil
}

func (rt Router) ListRanks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		common.GenerateErrorResponse(w, "unallowed method", http.StatusMethodNotAllowed)
		return
	}

	defer func() {
		if err := recover(); err != nil {
			common.GenerateErrorResponse(w, "unexpected internal error", http.StatusInternalServerError)
			return
		}
	}()

	limit, err := common.IsLimitValid(r.URL.RawQuery)
	if err != nil {
		common.GenerateErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	info, err := rt.rankListClient.GetRankList(limit)
	if err != nil {
		common.GenerateErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.GenerateSuccessResponseForRankList(w, info);
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}
