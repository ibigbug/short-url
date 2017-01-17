package main

import (
	"encoding/json"
	"net/http"
)

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	decode := json.NewDecoder(r.Body)
	var params ShortenParam
	err := decode.Decode(&params)
	if err != nil || params.Short == "" {
		http.Error(w, "Malformed JSON body", http.StatusBadRequest)
		return
	}

	return
}
