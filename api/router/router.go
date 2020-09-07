package router

import (
	"coins/api/client"
	"coins/common"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	router         *mux.Router
	priceSvcClient client.PriceSvcClientI
	rankSvcClient  client.RankSvcClientI
}

func NewRouter(priceSvcClient *client.PriceSvcClientI, rankSvcClient *client.RankSvcClientI) (*Router, error) {
	router := mux.NewRouter()

	r := &Router{
		router:         router,
		priceSvcClient: *priceSvcClient,
		rankSvcClient:  *rankSvcClient,
	}

	router.HandleFunc("/list", r.TopCryptoLister).Methods(http.MethodGet)

	return r, nil
}

func (rr *Router) TopCryptoLister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		common.GenerateErrorResponse(w, "unallowed method", http.StatusMethodNotAllowed)
		return
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			common.GenerateErrorResponse(w, "unexpected internal error", http.StatusInternalServerError)
			return
		}
	}()

	limit, err := common.IsLimitValid(r.URL.RawQuery)
	if err != nil {
		common.GenerateErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("calling GetRankingList")
	ranksResp, errRanking := rr.rankSvcClient.GetRankingList(limit)
	if errRanking != nil {
		common.GenerateErrorResponse(w, errRanking.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("calling Price")
	pricesResp, errPricing := rr.priceSvcClient.ListPrices()
	if errPricing != nil {
		common.GenerateErrorResponse(w, errPricing.Error(), http.StatusInternalServerError)
		return
	}
	result := common.CombineResults(*ranksResp.RankInfo, *pricesResp.PriceInfo)
	common.GenerateSuccessResponseForCombined(w, result)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}
