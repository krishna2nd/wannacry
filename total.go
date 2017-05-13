package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Rate struct {
	Code string  `json:"code"`
	Name string  `json:"name"`
	Rate float64 `json:"rate"`
}

func main() {
	totalReceived := getWalletReceived()
	rate := getConversionRates()
	totalUsd := float64(totalReceived) * float64(rate)
	log.Println("Total USD: ", totalUsd)
}

func getWalletReceived() int {
	blockChain := "https://blockchain.info/q/getreceivedbyaddress/"
	chains := [...]string{"13AM4VW2dhxYgXeQepoHkHSQuy6NgaEb94", "12t9YDPgwueZ9NyMgw519p7AA8isjr6SMw", "115p7UMMngoj1pMvkpHijcRdfJNXj6LrLn"}
	var total int
	for _, element := range chains {
		res, err := http.Get(blockChain + element)
		mustNot(err)

		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)

		i, err := strconv.Atoi(buf.String())
		mustNot(err)
		total += i
	}
	return total / 100000000
}

func getConversionRates() float64 {
	bitpay := "https://bitpay.com/api/rates"
	res, err := http.Get(bitpay)
	mustNot(err)

	rates := make([]Rate, 0)
	json.NewDecoder(res.Body).Decode(&rates)

	var result float64
	for _, element := range rates {
		if element.Code == "USD" {
			result = element.Rate
		}
	}
	return result
}

func mustNot(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
