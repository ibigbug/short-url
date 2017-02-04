package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	ADDR             = "0.0.0.0"
	PORT             = "80"
	DEFAULT_SITE_URL = "http://localhost"
)

var SITE_URL string

func main() {
	http.Handle("/shorten", LogMiddleware(POST(JsonMiddleware(http.HandlerFunc(ShortenHandler)))))
	http.Handle("/original", LogMiddleware(GET(JsonMiddleware(http.HandlerFunc(OriginalHandler)))))
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ADDR
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = PORT
	}

	SITE_URL = os.Getenv("SITE_URL")
	if SITE_URL == "" {
		SITE_URL = DEFAULT_SITE_URL
	}

	binding := fmt.Sprintf("%s:%s", addr, port)
	log.Println("Server running at:", binding)
	log.Fatal(http.ListenAndServe(binding, nil))
}
