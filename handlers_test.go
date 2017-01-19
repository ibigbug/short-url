package main

import (
	"fmt"
	"math"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func base62(i int) string {
	if i == 0 {
		return "0"
	}
	// TODO: may overflow
	r := make([]byte, 0, 6)
	for i != 0 {
		mod := int(math.Mod(float64(i), 62.0))
		r = append(r, CHAR_MAPPING[mod])
		i = i / 62
	}
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func TestOriginalHandler(t *testing.T) {

	handler := http.HandlerFunc(ShortenHandler)

	// create 100000 url
	for i := 0; i < 100000; i++ {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/shorten", strings.NewReader(fmt.Sprintf(`{"url":"http://a.very.long.url/%s"}`, strconv.Itoa(i))))
		if err != nil {
			t.Fatal(err)
		}
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("expected: %v, got: %v", http.StatusOK, status)
		}
		expected := fmt.Sprintf(`{"short": "http://localhost/%s"}`, base62(i+1))
		if rr.Body.String() != expected {
			t.Fatalf("expected: %v, got: %v", expected, rr.Body.String())
		}
	}

	/*for i := 0; i < 100000; i++ {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/original", strings.NewReader(fmt.Sprintf(`{"short": "http://localhost/%s"}`, base62(i+1))))
		if err != nil {
			t.Fatal(err)
		}
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("expected: %v, got: %v", http.StatusOK, status)
		}
		expected := fmt.Sprintf(`{"original": "http://a.very.long.url/%s"}`, strconv.Itoa(i+1))
		if rr.Body.String() != expected {
			t.Fatalf("expected: %v, got: %v", expected, rr.Body.String())
		}
	}*/
}
