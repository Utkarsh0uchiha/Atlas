package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Backend struct {
	URL   *url.URL
	Alive bool
}

func main() {
	backendURL, err := url.Parse("http://localhost:8081")
	if err != nil {
		panic(err)
	}
	backend := Backend{
		URL:   backendURL,
		Alive: true,
	}
	if !backend.Alive {
		panic("backend is not alive")
	}

	proxy := httputil.NewSingleHostReverseProxy(backend.URL)

	if err := http.ListenAndServe(":8080", proxy); err != nil {
		panic(fmt.Sprintf("failed to start proxy server on port 8080: %v", err))
	}
}
