package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "up")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "you requested %s", r.URL.Path)
}