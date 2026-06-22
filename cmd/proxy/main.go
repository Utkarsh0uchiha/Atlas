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
	backend1, err := url.Parse("http://localhost:8081")
	if err != nil {
		panic(err)
	}
	backend2, err := url.Parse("http://localhost:8082")
	if err != nil {
		panic(err)
	}

	backends := []Backend{
		{URL: backend1, Alive: true},
		{URL: backend2, Alive: true},
	}

	target := backends[0].URL
	proxy := httputil.NewSingleHostReverseProxy(target)

	http.ListenAndServe(":8080", proxy)
}
