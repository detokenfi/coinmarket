package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"net"
)
const (
    MaxIdleConns int = 100
    MaxIdleConnsPerHost int = 100
    IdleConnTimeout int = 90
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
            IdleConnTimeout:	 time.Duration(IdleConnTimeout)* time.Second,
        },
	
	Timeout: 20 * time.Second,
    }
    return client
}

type Ticker struct {
	Result []struct {
		ID             int    `json:"id"`
		CoinSymbol     string `json:"coin_symbol"`
		CurrencySymbol string `json:"currency_symbol"`
		Last           string `json:"last"`
		High           string `json:"high"`
		Low            string `json:"low"`
		Change         string `json:"change"`
		Percent        string `json:"percent"`
		Vol24H         string `json:"vol24H"`
		Amount         string `json:"amount"`
		LastCny        string `json:"last_cny"`
		HighCny        string `json:"high_cny"`
		LowCny         string `json:"low_cny"`
		LastUsd        string `json:"last_usd"`
		HighUsd        string `json:"high_usd"`
		LowUsd         string `json:"low_usd"`
	} `json:"result"`
	Cmd string `json:"cmd"`
}

func main() {
	url := "https://api.bibox.com/v1/mdata?cmd=marketAll"
	client := createHTTPClient()
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Accept-Charst", "UTF-8")
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var data_total Ticker
	err = json.Unmarshal([]byte(data), &data_total)
	// fmt.Println(data_total)

	for _, v := range data_total.Result {
		//low
		low := v.Low
		fmt.Println(low)

		//high
		high := v.High
		fmt.Println(high)

		//last
		last := v.Last
		fmt.Println(last)

		//vol
		coin_vol := v.Vol24H
		fmt.Println(coin_vol)

		//base_coin_name
		base_coin_name := v.CurrencySymbol
		fmt.Println(base_coin_name)

		//coin_name
		coin_name := v.CoinSymbol
		fmt.Println(coin_name)

		//symbol
		symbol := coin_name + "_" + base_coin_name
		fmt.Println(symbol)
        //ticker 数据存储在自己的数据源
	}

}
