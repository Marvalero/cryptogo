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

func calculateExchange(exc Exchange, client HttpClient) {
	for {
		url := fmt.Sprint("https://min-api.cryptocompare.com/data/price?fsym=", exc.Currency, "&tsyms=USD,GBP,EUR&tryConversion=false")
		resp, err := client.Get(url)
		if err != nil {
			fmt.Println("Error calling cryptocompare:", err)
			time.Sleep(5 * time.Second)
			continue
		}
		defer resp.Body.Close()
		writeResponse(resp.Body, exc)
		time.Sleep(10 * time.Second)
	}
}

func writeResponse(Body io.Reader, exc Exchange) {
	body, _ := ioutil.ReadAll(Body)
	var dat map[string]float64
	if err := json.Unmarshal(body, &dat); err != nil {
		fmt.Println("Error calling Unmarshal")
		return
	}
	exc.WriteCurrentValue <- dat["GBP"]
}
