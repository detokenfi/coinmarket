package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
	// "github.com/bitly/go-simplejson"
)

type HuobiTicker struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   []struct {
		Open   float64 `json:"open"`
		Close  float64 `json:"close"`
		Low    float64 `json:"low"`
		High   float64 `json:"high"`
		Amount float64 `json:"amount"`
		Count  int     `json:"count"`
		Vol    float64 `json:"vol"`
		Symbol string  `json:"symbol"`
	} `json:"data"`
}

//火币api解析
func main() {
	var i int
	var symbol string
	var high string
	var low string
	var last string
	//var vol string
	var coin_name string
	var base_coin_name string
	client := createHTTPClient()
	exchange_apiurl := "https://api.huobipro.com/market/tickers"
	request, _ := http.NewRequest("GET", exchange_apiurl, nil)
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := ioutil.ReadAll(response.Body)
	var data_totals HuobiTicker
	err = json.Unmarshal([]byte(data), &data_totals)
	//基础币
	basecoin := strings.Split("USDT_BTC_ETH_HT", "_")

	for _, v := range data_totals.Data {
		symbol = v.Symbol
		if strings.HasSuffix(symbol, "10") {
		} else {
			for i = 0; i < len(basecoin); i++ {
				if strings.HasSuffix(strings.ToUpper(symbol), basecoin[i]) {
					coin_name = strings.Replace(strings.ToUpper(symbol), basecoin[i], "", 1)
					base_coin_name = basecoin[i]
					symbol = coin_name + "_" + base_coin_name
					fmt.Println(symbol)
				}
			}
			//24h最新成交价
			close := float64(v.Close)
			last = strconv.FormatFloat(close, 'E', -1, 64)
			_ = last
			//24h最高成交价
			highing := float64(v.High)
			high = strconv.FormatFloat(highing, 'E', -1, 64)
			_ = high
			//24h最低成交价
			lowing := float64(v.Low)
			low = strconv.FormatFloat(lowing, 'E', -1, 64)
			_ = low
			//24h交易对的交易量
			vol := float64(v.Amount)
			coin_vol := strconv.FormatFloat(vol, 'E', -1, 64)
			_ = coin_vol
		}
	}
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
