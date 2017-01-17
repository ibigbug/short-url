package main

import (
	"log"
	"net/http"
	"strings"
)

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Headers: %v\n", r.Header)
		if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
			// application/json; charset=utf-8
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
