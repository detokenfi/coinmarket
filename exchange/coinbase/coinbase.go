package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Ticker_co []struct {
	ID             string `json:"id"`
	BaseCurrency   string `json:"base_currency"`
	QuoteCurrency  string `json:"quote_currency"`
	BaseMinSize    string `json:"base_min_size"`
	BaseMaxSize    string `json:"base_max_size"`
	QuoteIncrement string `json:"quote_increment"`
	DisplayName    string `json:"display_name"`
	Status         string `json:"status"`
	MarginEnabled  bool   `json:"margin_enabled"`
	StatusMessage  string `json:"status_message"`
	MinMarketFunds string `json:"min_market_funds"`
	MaxMarketFunds string `json:"max_market_funds"`
	PostOnly       bool   `json:"post_only"`
	LimitOnly      bool   `json:"limit_only"`
	CancelOnly     bool   `json:"cancel_only"`
}

type Tickers_co struct {
	Open        string `json:"open"`
	High        string `json:"high"`
	Low         string `json:"low"`
	Volume      string `json:"volume"`
	Last        string `json:"last"`
	Volume30Day string `json:"volume_30day"`
}

func main() {
	url := "https://api.pro.coinbase.com/products"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var data_total Ticker_co
	err = json.Unmarshal([]byte(data), &data_total)

	for _, v := range data_total {
		sym := v.ID
		sym_spl := strings.Split(sym, "-")
		//base_coin_name
		coin_name := sym_spl[0]
		//fmt.Println(coin_name)

		//coin_name
		base_coin_name := sym_spl[1]
		//fmt.Println(base_coin_name)
		//symbol
		symbol := coin_name + "_" + base_coin_name
		fmt.Println(symbol)

		urls := "https://api.pro.coinbase.com/products/" + v.ID + "/stats"
		resp, err := http.Get(urls)
		if err != nil {
			fmt.Println(err)
		}
		data, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		var data_total Tickers_co
		err = json.Unmarshal([]byte(data), &data_total)

		//high
		high := data_total.High
		fmt.Println(high)

		//low
		low := data_total.Low
		fmt.Println(low)

		//last
		last := data_total.Last
		fmt.Println(last)

		//vol
		coin_vol := data_total.Volume
		fmt.Println(coin_vol)

	}

}
