package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type TickersBit struct {
	High      string `json:"high"`
	Last      string `json:"last"`
	Timestamp string `json:"timestamp"`
	Bid       string `json:"bid"`
	Vwap      string `json:"vwap"`
	Volume    string `json:"volume"`
	Low       string `json:"low"`
	Ask       string `json:"ask"`
	Open      string `json:"open"`
}

func main() {
	var symbol string
	var i int
	var arr = [...]string{"btcusd", "btceur", "eurusd", "xrpusd", "xrpeur", "xrpbtc", "ltcusd", "ltceur", "ltcbtc", "ethusd", "etheur", "ethbtc", "bchusd", "bcheur", "bchbtc"}
	for i = 0; i < len(arr); i++ {
		bit_url := "https://www.bitstamp.net/api/v2/ticker/" + arr[i] + ""
		resp, err := http.Get(bit_url)
		if err != nil {
			fmt.Println(err)
		}
		data, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		var data_total TickersBit
		err = json.Unmarshal([]byte(data), &data_total)
		sym := arr[i]
		sym_upp := strings.ToUpper(sym)
		sym_usd := strings.Replace(sym_upp, "USD", "_USD", 1)
		sym_btc := strings.Replace(sym_usd, "BTC", "_BTC", 1)
		sym_eur := strings.Replace(sym_btc, "EUR", "_EUR", 1)
		c := []byte(sym_eur)
		if string(c[0]) == "_" {
			symbol = strings.Replace(sym_eur, "_", "", 1)
		} else {
			symbol = sym_eur
		}
		fmt.Println(symbol)
		//low
		low := data_total.Low
		fmt.Println(low)

		//high
		high := data_total.High
		fmt.Println(high)

		//last
		last := data_total.Last
		fmt.Println(last)

		//vol
		coin_vol := data_total.Volume
		fmt.Println(coin_vol)

		//sell
		sell := data_total.Ask
		fmt.Println(sell)

		//buy
		buy := data_total.Bid
		fmt.Println(buy)

		sym_spl := strings.Split(symbol, "_")

		//base_coin_name
		base_coin_name := sym_spl[1]
		fmt.Println(base_coin_name)

		//coin_name
		coin_name := sym_spl[0]
		fmt.Println(coin_name)
		
	}
}
