package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"net"
	"github.com/bitly/go-simplejson"
	"fmt"
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
	var coin_vol string
	var last string
	var sell string
	var buy string
	var high string
	var low string
	var base_coin_name string
	var coin_name string
	var symbol_ok string

	url := "http://api.zb.cn/data/v1/allTicker"
	client := createHTTPClient()
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Accept-Charst", "UTF-8")
	resp, err := client.Do(request)
	if err!=nil {
		fmt.Println(err)
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	json, _ := simplejson.NewJson(body)
	// fmt.Println(json)
	var nodes = make(map[string]interface{})
	nodes, _ = json.Map()

	// ticker := "ticker"
	for k, v := range nodes {
		symbol := strings.ToUpper(k)
		QC_sym := strings.Replace(symbol, "QC", "_QC", 1)
		USDT_sym := strings.Replace(QC_sym, "USDT", "_USDT", 1)
		ZB_sym := strings.Replace(USDT_sym, "ZB", "_ZB", 1)
		BTC_sym := strings.Replace(ZB_sym, "BTC", "_BTC", 2)

		node_data := v.(map[string]interface{})
		for key, v1 := range node_data {
			if key == "low" {
				low = v1.(string)
			}

			if key == "high" {
				high = v1.(string)
			}

			if key == "buy" {
				buy = v1.(string)
			}

			if key == "sell" {
				sell = v1.(string)
			}

			if key == "last" {
				last = v1.(string)
			}

			if key == "vol" {
				coin_vol = v1.(string)
			}
		}
		//fmt.Println(low, high, buy, sell, last, vol)

		// btc_split := strings.Split(BTC_sym, "_")
		c := []byte(BTC_sym)
		if string(c[0]) == "_" {
			// 	//去除下划线在前
			symbol_ok = strings.Replace(BTC_sym, "_", "", 1)
		//	fmt.Println(sym_ok)
		} else {
			btc_spl := strings.Split(BTC_sym, "_")
			if len(btc_spl) > 2 {
				btc_end := strings.HasSuffix(BTC_sym, "BTC")
				usdt_end := strings.HasSuffix(BTC_sym, "USDT")
				qc_end := strings.HasSuffix(BTC_sym, "QC")
				zb_end := strings.HasSuffix(BTC_sym, "ZB")
				//尾部等于BTC
				if btc_end == true {
					btc_end_add := strings.Replace(btc_spl[2], btc_spl[2], "_BTC", 1)
					symbol_ok = btc_spl[0] + btc_spl[1] + btc_end_add
					//尾部等于USDT
				} else if usdt_end == true {
					usdt_end_add := strings.Replace(btc_spl[2], btc_spl[2], "_USDT", 1)
					symbol_ok = btc_spl[0] + btc_spl[1] + usdt_end_add
					//尾部等于QC
				} else if qc_end == true {
					qc_end_add := strings.Replace(btc_spl[2], btc_spl[2], "_QC", 1)
					symbol_ok = btc_spl[0] + btc_spl[1] + qc_end_add
					//尾部等于ZB
				} else if zb_end == true {
					zb_end_add := strings.Replace(btc_spl[2], btc_spl[2], "_ZB", 1)
					symbol_ok = btc_spl[0] + btc_spl[1] + zb_end_add

				}
			} else {
				//symbol
				symbol_ok = BTC_sym
			}
			}
			fmt.Println(symbol_ok)

			coin_name_spl := strings.Split(symbol_ok, "_")

			//coin_name
			coin_name = coin_name_spl[0]
			//base_coin_name
			base_coin_name = coin_name_spl[1]
			fmt.Println(coin_name)
			fmt.Println(base_coin_name)
			fmt.Println(coin_vol,last,sell,buy,high,low)
			
	}
}
