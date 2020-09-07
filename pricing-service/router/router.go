package router

import (
	"coins/common"
	"coins/pricing-service/client"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	router          *mux.Router
	priceListClient client.PriceListClientI
}

func NewRouter(priceListClient *client.PriceListClientI) (*Router, error) {
	router := mux.NewRouter()

	r := &Router{
		router:          router,
		priceListClient: *priceListClient,
	}

	router.HandleFunc("/usd-prices", r.ListPrices).Methods(http.MethodGet)

	return r, nil
}

func (re *Router) ListPrices(w http.ResponseWriter, r *http.Request) {
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

	info, err := re.priceListClient.GetUsdPrices()
	if err != nil {
		common.GenerateErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.GenerateSuccessResponseForPriceList(w, info);
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}
