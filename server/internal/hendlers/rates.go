package hendlers

import (
	"AlexLitovchenko/TestBinance/server/internal/errs"
	"AlexLitovchenko/TestBinance/server/internal/service"
	"encoding/json"
	"net/http"
	"strings"
)

func GetRatesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		RatesGetRequest(w, r)
	} else if r.Method == http.MethodPost {
		RatesPostRequest(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func RatesGetRequest(w http.ResponseWriter, r *http.Request) {
	pairsParam := r.URL.Query().Get("pairs")
	pairs := strings.Split(pairsParam, ",")

	resp, err := service.MakeRatesResp(pairs)
	if err != nil {
		http.Error(w, err.Msg, err.Key)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), errs.ErrMarshal)
		return
	}
}

func RatesPostRequest(w http.ResponseWriter, r *http.Request) {

	var request struct {
		Pairs []string `json:"pairs"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := service.MakeRatesResp(request.Pairs)
	if err != nil {
		http.Error(w, err.Msg, err.Key)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), errs.ErrMarshal)
		return
	}
}
