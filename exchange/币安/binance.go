package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

type BinanceTicker []struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	BidQty             string `json:"bidQty"`
	AskPrice           string `json:"askPrice"`
	AskQty             string `json:"askQty"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstID            int    `json:"firstId"`
	LastID             int    `json:"lastId"`
	Count              int    `json:"count"`
}

//币安网api解析
func main() {
	exchange_apiurl := "https://api.binance.com/api/v1/ticker/24hr"
	client := createHTTPClient()
	request, _ := http.NewRequest("GET", exchange_apiurl, nil)
	request.Header.Set("Accept-Charst", "UTF-8")
	resp, _ := client.Do(request)
	data, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var data_total BinanceTicker
	err := json.Unmarshal([]byte(data), &data_total)
	if err != nil {
	}
	for _, v := range data_total {
		//24h最高成交价
		high := v.HighPrice
		fmt.Println(high)

		//24h最低成交价
		low := v.LowPrice
		fmt.Println(low)

		//24h最新成交价
		last := v.LastPrice
		fmt.Println(last)

		//24h交易对的交易量（基础币）
		coin_base_vol := v.QuoteVolume
		fmt.Println(coin_base_vol)

		//symbol
		symbol := v.Symbol
		btc_sym := strings.Replace(symbol, "BTC", "_BTC", 1)
		eth_sym := strings.Replace(btc_sym, "ETH", "_ETH", 1)
		usdt_sym := strings.Replace(eth_sym, "USDT", "_USDT", 1)
		bnb_syn := strings.Replace(usdt_sym, "BNB", "_BNB", 1)
		c := []byte(bnb_syn)
		if string(c[0]) == "_" {
			sym_ok := strings.Replace(bnb_syn, "_", "", 1)
			base_coin := strings.FieldsFunc(sym_ok, split)
			base_coin_name := base_coin[1]
			fmt.Println(base_coin_name)
			coin_name := base_coin[0]
			fmt.Println(coin_name)
		} else {
			sym_ok := bnb_syn
			base_coin := strings.FieldsFunc(sym_ok, split)
			base_coin_name := base_coin[1]
			fmt.Println(base_coin_name)
			coin_name := base_coin[0]
			fmt.Println(coin_name)
		}

	}

}
func split(s rune) bool {
	if s == '_' {
		return true
	}
	return false
}

const (
	MaxIdleConns        int = 100
	MaxIdleConnsPerHost int = 100
	IdleConnTimeout     int = 90
)

func createHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        MaxIdleConns,
			MaxIdleConnsPerHost: MaxIdleConnsPerHost,
			IdleConnTimeout:     time.Duration(IdleConnTimeout) * time.Second,
		},

		Timeout: 20 * time.Second,
	}
	return client
}
