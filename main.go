package main

import (
	"net/http"
)

func main() {
	http.Handle("/shorten", jsonMiddleware(http.HandlerFunc(ShortenHandler)))
}
