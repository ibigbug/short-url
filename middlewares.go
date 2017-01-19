package main

import (
	"log"
	"net/http"
	"strings"
	"time"
)

func JsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Golang will convert the header keys into canonical format
		// Since RFC speaks HTTP header keys are case insensive
		if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
			// application/json; charset=utf-8
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GET(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func POST(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		next.ServeHTTP(w, r)
		// Must wrap a new ResponseWriter to get the status code
		log.Printf("%s %s %s %dms", r.Method, r.RequestURI, r.Proto, time.Now().Sub(t)/time.Millisecond)
	})
}
