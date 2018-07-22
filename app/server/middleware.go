package server

import (
	"log"
	"net/http"
)

type middleware func(next http.HandlerFunc) http.HandlerFunc

func withLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		log.Printf("Connection from %s path %s %s", r.RemoteAddr, r.Method, r.URL.Path)
	}
}
