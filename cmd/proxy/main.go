package main

import (
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

	proxy := httputil.NewSingleHostReverseProxy(backend.URL)

	if err := http.ListenAndServe(":8080", proxy); err != nil {
		panic(err)
	}
}
