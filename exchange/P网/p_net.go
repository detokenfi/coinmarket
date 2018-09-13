package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"net"
	"github.com/bitly/go-simplejson"

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

func main() {
	//定义变量
	var coin_base_vol string
	var last string
	//var sell string
	//var buy string
	var high string
	var low string
	var base_coin_name string
	var coin_name string
	//var symbol_ok string

	url := "https://poloniex.com/public?command=returnTicker"
	client := createHTTPClient()
	request, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(request)
	if err!=nil {
		fmt.Println(err)
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	json, _ := simplejson.NewJson(body)
	var nodes = make(map[string]interface{})
	nodes, _ = json.Map()
	for k, v := range nodes {
		//btc eth usdt xmr
		symbol := strings.ToUpper(k)
		sz := strings.Split(symbol, "_")
		base_coin_name = sz[0]
		coin_name = sz[1]
		symbol = coin_name + "_" + base_coin_name
		fmt.Println(symbol)
		node_data := v.(map[string]interface{})
		for key, v1 := range node_data {
			if key == "last" {
				last = v1.(string)
				fmt.Println(last)
			}
			if key == "high24hr" {
				high = v1.(string)
				fmt.Println(high)
			}
			if key == "low24hr" {
				low = v1.(string)
				fmt.Println(low)
			}
			if key == "baseVolume" {
				coin_base_vol = v1.(string)
				fmt.Println(coin_base_vol)
			}
		}
	}
}
