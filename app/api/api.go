package api

import (
	"fmt"
	"net/http"
)

type ExchangeApi struct {
	Exchannel chan float64
}

func (exc ExchangeApi) ShowExchange(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	currentExchange := <-exc.Exchannel
	message := fmt.Sprintf("exchange: %.6f", currentExchange)
	w.Write([]byte(message))
}
