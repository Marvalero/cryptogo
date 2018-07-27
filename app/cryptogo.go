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
	readChan := exchange_calculator.StartExchangeCalculator(&client)

	server.Run(readChan)
	fmt.Println("Killing cryptogo App")

	return nil
}
