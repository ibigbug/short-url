package main

import "testing"
import "net/http"
import "strings"
import "net/http/httptest"

func BenchShort(t *testing.B) {
	req, err := http.NewRequest("POST", "/shorten", strings.NewReader(`{"url":"http://a.very.long.url"}`))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ShortenHandler)

	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected: %v, got: %v", http.StatusOK, status)
	}
	expected := `{"short": "http://localhost/1"}`
	if rr.Body.String() != expected {
		t.Errorf("expected: %v, got: %v", expected, rr.Body.String())
	}
}
