package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Response from server 1")
}
func healthhandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

}
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/health", healthhandler)
	http.HandleFunc("/kill", func(w http.ResponseWriter, r *http.Request) {
		os.Exit(1)
	})
	http.ListenAndServe(":8081", nil)
}
