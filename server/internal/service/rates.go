package service

import (
	"AlexLitovchenko/TestBinance/server/internal/errs"
	"AlexLitovchenko/TestBinance/server/internal/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const binanceApiUrl = "https://api.binance.com/api/v3/ticker/price"

func MakeRatesResp(pairs []string) (map[string]float32, *errs.Error) {
	binanceResp, err := binanceRequestPrice()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	mapPair := make(map[string]string, len(binanceResp))
	for i := range binanceResp {
		mapPair[binanceResp[i].Symbol] = binanceResp[i].Price
	}

	rates := make(map[string]float32)
	for i, pair := range pairs {
		if value, ok := mapPair[strings.ReplaceAll(pairs[i], "-", "")]; ok {
			price, err := strconv.ParseFloat(value, 32)
			if err != nil {
				fmt.Println(err)
				return nil, errs.Err(errs.ErrParsePrice, err.Error())
			}
			rates[pair] = float32(price)
		}
	}

	if len(rates) == 0 {
		fmt.Println(errs.Err(errs.ErrPairsNotFound))
		return nil, errs.Err(errs.ErrPairsNotFound)
	}
	return rates, nil
}

func binanceRequestPrice() ([]model.BinanceResp, *errs.Error) {

	var respBody []model.BinanceResp
	resp, err := http.Get(binanceApiUrl)
	if err != nil {
		fmt.Println(errs.Err(errs.ErrGetBinance, err.Error()))
		return nil, errs.Err(errs.ErrGetBinance, err.Error())
	}
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		fmt.Println(errs.Err(errs.ErrUnmarshall, err.Error()))
		return nil, errs.Err(errs.ErrUnmarshall, err.Error())
	}

	return respBody, nil
}
