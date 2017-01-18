package main

import (
	"log"
	"net/http"
)

const (
	ADDR     = "localhost:8000"
	SITE_URL = "http://localhost"
)

func main() {
	http.Handle("/shorten", LogMiddleware(POST(JsonMiddleware(http.HandlerFunc(ShortenHandler)))))
	http.Handle("/original", LogMiddleware(GET(JsonMiddleware(http.HandlerFunc(OriginalHandler)))))
	log.Println("Server running at:", ADDR)
	log.Fatal(http.ListenAndServe(ADDR, nil))
}
