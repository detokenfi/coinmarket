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

type AutoGenera struct {
	Date    string `json:"date"`
	Tickers []struct {
		Symbol string `json:"symbol"`
		High   string `json:"high"`
		Vol    string `json:"vol"`
		Last   string `json:"last"`
		Low    string `json:"low"`
		Buy    string `json:"buy"`
		Sell   string `json:"sell"`
	} `json:"tickers"`
}

//okex的api解析
func main() {
	client := createHTTPClient()
	exchange_apiurl := "https://www.okex.com/api/v1/tickers.do"
	request, _ := http.NewRequest("GET", exchange_apiurl, nil)
	request.Header.Set("Accept-Charst", "UTF-8")
	response, _ := client.Do(request)
	body, _ := ioutil.ReadAll(response.Body)
	bodyStr := string(body)
	auto := AutoGenera{}
	err := json.Unmarshal([]byte(bodyStr), &auto)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range auto.Tickers {
		//交易对
		symbol := strings.ToUpper(v.Symbol)
		fmt.Println(symbol)
		//24h最高成交价
		high := v.High
		fmt.Println(high)
		//24h交易对的交易量
		coin_vol := v.Vol
		fmt.Println(coin_vol)
		//24h最新成交价
		last := v.Last
		fmt.Println(last)
		//24h最低成交价
		low := v.Low
		fmt.Println(low)
		//买一价
		buy := v.Buy
		fmt.Println(buy)
		//卖一价
		sell := v.Sell
		fmt.Println(sell)
		sz := strings.Split(symbol, "_")
		var coin_name string
		var base_coin_name string
		for i := 0; i < 2; i++ {
			coin_name = sz[0]
			fmt.Println(coin_name)
			base_coin_name = sz[1]
			fmt.Println(base_coin_name)
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
