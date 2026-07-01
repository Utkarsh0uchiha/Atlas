package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Response from server 2")
}
func healthhandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/health", healthhandler)
	http.ListenAndServe(":8082", nil)
}
