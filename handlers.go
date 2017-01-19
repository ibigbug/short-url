package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var shortService = NewShortService(nil, nil)

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var params ShortenParam
	err := decoder.Decode(&params)
	if err != nil || params.Url == "" {
		http.Error(w, "Malformed JSON body", http.StatusBadRequest)
		return
	}

	shorten := shortService.Short(params.Url)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"short": "%s/%s"}`, SITE_URL, shorten)
}

func OriginalHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var params OriginalParam
	err := decoder.Decode(&params)
	if err != nil || params.Short == "" {
		http.Error(w, "Malformed JSON body: "+err.Error(), http.StatusBadRequest)
		panic(params.Short)
		return
	}

	u, err := url.ParseRequestURI(params.Short)
	if err != nil {
		http.Error(w, "Invalid URL", http.StatusPreconditionFailed)
	}
	original := shortService.Original(u.Path[1:])
	if original == "" {
		http.Error(w, "No such record", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"original": "%s"}`, original)
}
