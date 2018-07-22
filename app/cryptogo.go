package app

import (
	"fmt"
	"time"

	exchange "github.com/Marvalero/cryptogo/app/exchanger"
	"github.com/Marvalero/cryptogo/app/server"
)

// Run creates and executes new kubeadm command
func Run() error {
	// We do not want these flags to show up in --help
	fmt.Println("Starting cryptogo App")
	readChan := exchange.StartExchange()
	go readFromExchange(readChan)

	server.Run()
	fmt.Println("Killing cryptogo App")

	return nil
}

func readFromExchange(readChan chan float64) {
	val := 0.0
	for {
		time.Sleep(10000 * time.Millisecond)
		val = <-readChan
		fmt.Println("OUT read form readChan!!!", val)
	}
}
