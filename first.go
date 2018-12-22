package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
)

const BaseUrl = "https://1token.trade/api/v1/quote"

func Get_quote(endpoint string) {
	fmt.Println(BaseUrl + endpoint)
	resp, err := http.Get(BaseUrl + endpoint)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	js, err := simplejson.NewJson(body)
	fmt.Println(js)

}

func main() {
	Get_quote("/single-tick/okex/btc.usdt")
}
