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
	var high string
	var low string
	var last string
	var coin_name string
	var base_coin_name string
	var symbol string
	var coin_vol string
	url := "https://api.kraken.com/0/public/AssetPairs"
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
		if k == "result" {
			val := v.(map[string]interface{})
			for key, value := range val {
				sym := key
				//获取交易对
				var node2 = value.(map[string]interface{})
				for k1, v1 := range node2 {
					if k1 == "altname" {
						fmt.Println(v1)
						if strings.Contains(v1.(string), ".d") {
						} else {
							coin_name = string([]byte(v1.(string)[:(len(v1.(string)) - 3)]))
							if coin_name == "XBT" {
								coin_name = "BTC"
								fmt.Println("coin_name:" + coin_name)
								base_coin_name = strings.Replace(v1.(string), "XBT", "", -1)
								if base_coin_name == "XBT" {
									base_coin_name = "BTC"
								}
								fmt.Println("base_coin_name:" + base_coin_name)
								symbol = coin_name + "_" + base_coin_name
							} else {
								fmt.Println("coin_name:" + coin_name)
								base_coin_name = strings.Replace(v1.(string), coin_name, "", -1)
								if base_coin_name == "XBT" {
									base_coin_name = "BTC"
								}
								fmt.Println("base_coin_name:" + base_coin_name)
								symbol = coin_name + "_" + base_coin_name
							}

							fmt.Println("symbol:" + symbol)
						}
					}
				}
				urls := "https://api.kraken.com/0/public/Ticker?pair=" + sym + ""
				reqs, _ := http.NewRequest("GET", urls, nil)
				ress, _ := http.DefaultClient.Do(reqs)
				bodys, _ := ioutil.ReadAll(ress.Body)
				jsons, _ := simplejson.NewJson(bodys)
				var tickdata = make(map[string]interface{})
				tickdata, _ = jsons.Map()
				for keys, vals := range tickdata {
					if keys == "result" {
						ticker_val := vals.(map[string]interface{})
						for _, tick_val := range ticker_val {
							// symbol := tick_key
							// fmt.Println(symbol)

							tick_vals := tick_val.(map[string]interface{})
							for value_k, value_v := range tick_vals {
								if value_k == "v" {
									vol_all := value_v.([]interface{})
									coin_vol = vol_all[1].(string)

								}
								if value_k == "h" {
									high_all := value_v.([]interface{})
									high = high_all[0].(string)
								}
								if value_k == "l" {
									low_all := value_v.([]interface{})
									low = low_all[0].(string)
								}
								if value_k == "c" {
									last_all := value_v.([]interface{})
									last = last_all[0].(string)
								}
							}
						//入库
						fmt.Println(high,low,last,coin_vol,)
						}
					}
				}
			}
		}
	}
}
