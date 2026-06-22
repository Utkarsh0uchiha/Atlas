package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Backend struct {
	URL *url.URL
}

func main() {
	backendURL, err := url.Parse("http://localhost:8081")
	if err != nil {
		panic(err)
	}
	backend := Backend{
		URL: backendURL,
	}

	proxy := httputil.NewSingleHostReverseProxy(backend.URL)

	if err := http.ListenAndServe(":8080", proxy); err != nil {
		log.Fatalf("failed to start proxy server on port 8080: %v", err)
	}
}
