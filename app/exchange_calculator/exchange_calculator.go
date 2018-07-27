package exchange_calculator

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpClient interface {
	Get(url string) (*http.Response, error)
}

func calculateExchange(writeChan chan float64, client HttpClient) {
	for {
		resp, err := client.Get("https://min-api.cryptocompare.com/data/price?fsym=ETH&tsyms=USD,GBP,EUR&tryConversion=false")
		if err != nil {
			fmt.Println("Error calling cryptocompare")
			time.Sleep(5 * time.Second)
			continue
		}
		defer resp.Body.Close()
		writeResponse(resp.Body, writeChan)
		time.Sleep(5 * time.Second)
	}
}

func writeResponse(Body io.Reader, writeChan chan float64) {
	body, _ := ioutil.ReadAll(Body)
	var dat map[string]float64
	if err := json.Unmarshal(body, &dat); err != nil {
		fmt.Println("Error calling Unmarshal")
		return
	}
	fmt.Println("Current exchange GBP:", dat["GBP"])
	writeChan <- dat["GBP"]
}

func StartExchangeCalculator(client HttpClient) chan float64 {
	exc := NewExchange()

	go calculateExchange(exc.WriteChan, client)
	return exc.ReadChan
}
