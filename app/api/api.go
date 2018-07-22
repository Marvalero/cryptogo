package api

import (
	"net/http"
)

func ShowExchange(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("exchange"))
}
