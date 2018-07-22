package server

import (
	"fmt"
	"net/http"

	"github.com/Marvalero/cryptogo/app/api"
)

func Run() {
	http.HandleFunc("/ping", withLogging(pong))
	http.HandleFunc("/exchange", withLogging(api.ShowExchange))
	http.HandleFunc("/", withLogging(notFound))
	fmt.Println("Serving on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	message := "NotFound path=" + r.URL.Path
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(message))
}

func pong(w http.ResponseWriter, r *http.Request) {
	message := "pong"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}
