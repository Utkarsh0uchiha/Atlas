package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Response from server 1")
}
func healthhandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(10 * time.Second)
	w.WriteHeader(http.StatusOK)

}
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/health", healthhandler)
	http.ListenAndServe(":8081", nil)
}
