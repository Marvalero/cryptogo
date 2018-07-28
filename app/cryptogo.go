package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Marvalero/cryptogo/app/exchange_calculator"
	"github.com/Marvalero/cryptogo/app/server"
)

// Run creates and executes new kubeadm command
func Run() error {
	// We do not want these flags to show up in --help
	fmt.Println("Starting cryptogo App")
	client := http.Client{Timeout: time.Duration(2 * time.Second)}
	readEthChan := exchange_calculator.StartExchangeCalculator(&client, "ETH", 380.0, 320.0)
	readBtcChan := exchange_calculator.StartExchangeCalculator(&client, "BTC", 6300.0, 6100.0)

	server.Run(readEthChan, readBtcChan)
	fmt.Println("Killing cryptogo App")

	return nil
}
