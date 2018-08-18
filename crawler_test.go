package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCrawl(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Response")
	}))
	defer server.Close()

	_, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}
}
