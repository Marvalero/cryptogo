package exchange

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Exchange struct {
	currentValue float64
}

func (exc Exchange) readWriteExchange(readChan chan float64, writeChan chan float64) {
	for {
		select {
		case readChan <- exc.currentValue:
			continue
		case exc.currentValue = <-writeChan:
			continue
		}
	}
}

func calculateExchange(writeChan chan float64) {
	client := http.Client{Timeout: time.Duration(2 * time.Second)}
	for {
		time.Sleep(5 * time.Second)

		resp, err := client.Get("https://min-api.cryptocompare.com/data/price?fsym=ETH&tsyms=USD,GBP,EUR&tryConversion=false")
		if err != nil {
			fmt.Println("Error calling cryptocompare")
			continue
		}
		defer resp.Body.Close()
		writeResponse(resp.Body, writeChan)
	}
}

func writeResponse(Body io.Reader, writeChan chan float64) {
	body, _ := ioutil.ReadAll(Body)
	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		fmt.Println("Error calling Unmarshal")
		return
	}
	fmt.Println("Current exchange GBP:", dat["GBP"])
	writeChan <- dat["GBP"].(float64)
}

func StartExchange() chan float64 {
	writeChan := make(chan float64)
	readChan := make(chan float64)
	exc := Exchange{}

	go exc.readWriteExchange(readChan, writeChan)
	go calculateExchange(writeChan)
	return readChan
}
