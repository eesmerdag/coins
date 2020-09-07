package common

import (
	"encoding/json"
	"net/http"
)

type SvcResponse struct {
	Message      string
	PriceInfo    *PriceInfo
	RankInfo     *RankInfo
	CombinedInfo []*CombinedData
	HttpStatus   int
}

type ApiResponse struct {
	Message      string
	CombinedInfo *[]CombinedData
	HttpStatus   int
}

func GenerateSuccessResponseForCombined(w http.ResponseWriter, res *[]CombinedData) {

	sccObj := ApiResponse{
		Message:      "successful",
		CombinedInfo: res,
		HttpStatus:   http.StatusOK,
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sccObj)
}

func GenerateSuccessResponseForPriceList(w http.ResponseWriter, res *PriceInfo) {
	sccObj := SvcResponse{
		Message:    "successful",
		PriceInfo:  res,
		HttpStatus: http.StatusOK,
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sccObj)
}

func GenerateSuccessResponseForRankList(w http.ResponseWriter, res *RankInfo) {
	sccObj := SvcResponse{
		Message:    "successful",
		RankInfo:   res,
		HttpStatus: http.StatusOK,
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sccObj)
}

func GenerateErrorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	errObj := SvcResponse{
		Message:    message,
		HttpStatus: httpStatusCode,
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(httpStatusCode)
	json.NewEncoder(w).Encode(errObj)
}
