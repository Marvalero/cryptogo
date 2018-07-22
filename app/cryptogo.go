package app

import (
	"fmt"

	exchange "github.com/Marvalero/cryptogo/app/exchanger"
	"github.com/Marvalero/cryptogo/app/server"
)

// Run creates and executes new kubeadm command
func Run() error {
	// We do not want these flags to show up in --help
	fmt.Println("Starting cryptogo App")
	readChan := exchange.StartExchange()

	server.Run(readChan)
	fmt.Println("Killing cryptogo App")

	return nil
}
