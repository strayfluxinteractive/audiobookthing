package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

//Index Generates the index
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

//AudiobookIndex returns a json structure of all audiobooks
func AudiobookIndex(w http.ResponseWriter, r *http.Request) {
	audiobooks := Audiobooks{
		Audiobook{Name: "Book1", Path: "/"},
		Audiobook{Name: "Book2", Path: "/test"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(audiobooks); err != nil {
		panic(err)
	}

}
