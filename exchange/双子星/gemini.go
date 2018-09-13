package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/henrylee2cn/pholcus/common/simplejson"
)

//双子星api解析
func main() {
	var coin_base_vol string
	var last string
	var high string
	var low string
	var base_coin_name string
	var coin_name string
	var symbol string

	sym_sz := [6]string{"btcusd", "ethbtc", "ethusd", "zecusd", "zecbtc", "zeceth"}
	for i := 0; i < len(sym_sz); i++ {
		url := "https://api.gemini.com/v1/pubticker/" + sym_sz[i]
		req, _ := http.NewRequest("GET", url, nil)
		res, _ := http.DefaultClient.Do(req)
		body, _ := ioutil.ReadAll(res.Body)
		json, _ := simplejson.NewJson(body)
		var nodes = make(map[string]interface{})
		nodes, _ = json.Map()
		coin_name = strings.ToUpper(string([]byte(sym_sz[i])[:3]))
		fmt.Println("coin_name:" + coin_name)
		base_coin_name = strings.Replace(strings.ToUpper(sym_sz[i]), coin_name, "", -1)
		fmt.Println("base_coin_name:" + base_coin_name)
		symbol = coin_name + "_" + base_coin_name
		fmt.Println(symbol)
		for k, v := range nodes {
			// symbol = strings.ToUpper(sym_sz[i])
			// fmt.Println("symbol:" + symbol)
			if k == "bid" {
				low = v.(string)
				fmt.Println(low)
			}
			if k == "ask" {
				high = v.(string)
				fmt.Println(high)
			}
			if k == "last" {
				last = v.(string)
				fmt.Println(last)
			}
			if k == "volume" {
				var node2 = v.(map[string]interface{})
				for k1, v1 := range node2 {
					if k1 == base_coin_name {
						coin_base_vol = v1.(string)
						fmt.Println(coin_base_vol)
					}
				}
			}
		}
		//insert(db, symbol, exchange_id, high, low, last, coin_base_vol, "", "", "", coin_name, base_coin_name, create_time)
	}
}
