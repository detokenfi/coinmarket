package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/henrylee2cn/pholcus/common/simplejson"
)

//比特儿海外版api解析
func main() {
	var last string
	var sell string
	var buy string
	var low string
	var coin_vol string
	var high string
	exchange_apiurl := "https://data.gateio.io/api2/1/tickers"
	req, _ := http.NewRequest("GET", exchange_apiurl, nil)
	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	json, _ := simplejson.NewJson(body)
	// fmt.Println(json)
	var nodes = make(map[string]interface{})
	nodes, _ = json.Map()

	for k, v := range nodes {
		//symbol
		symbol_upper := strings.ToUpper(k)

		symbol_split := strings.Split(symbol_upper, "_")

		base_coin_name := symbol_split[1]
		fmt.Println(base_coin_name)

		coin_name := symbol_split[0]
		fmt.Println(coin_name)
		node_data := v.(map[string]interface{})
		for key, v1 := range node_data {
			//24h最新成交价
			if key == "last" {
				last = v1.(string)
				fmt.Println(last)
			}
			//24h最高成交价
			if key == "high24hr" {
				high = v1.(string)
				fmt.Println(high)
			}
			//24h最低成交价
			if key == "low24hr" {
				low = v1.(string)
				fmt.Println(low)
			}
			//卖一价
			if key == "lowestAsk" {
				sell = v1.(string)
				fmt.Println(sell)
			}
			//买一价
			if key == "highestBid" {
				buy = v1.(string)
				fmt.Println(buy)
			}
			//24h交易对的交易量
			if key == "baseVolume" {
				coin_vol = v1.(string)
				fmt.Println(coin_vol)
			}
		}
	}
}
