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
type LoadBalancer struct {
	Backends []Backend
	Current  int
}

func (lb *LoadBalancer) NextBackend() Backend {
	current := lb.Current

	lb.Current = (current + 1) % len(lb.Backends)
	return lb.Backends[current]
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	backend := lb.NextBackend()

	proxy := httputil.NewSingleHostReverseProxy(backend.URL)

	proxy.ServeHTTP(w, r)
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

	lb := LoadBalancer{
		Backends: backends,
		Current:  0,
	}

	http.ListenAndServe(":8080", &lb)

}
