package main

import (
	"errors"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type Backend struct {
	URL   *url.URL
	Alive bool
}
type LoadBalancer struct {
	Backends []Backend
	Current  int
}

func (lb *LoadBalancer) HealthCheck() {
	for i := range lb.Backends {
		resp, err := http.Get(lb.Backends[i].URL.String())
		if err != nil {
			lb.Backends[i].Alive = false
			continue
		}
		if resp.StatusCode == http.StatusOK {
			lb.Backends[i].Alive = true
		} else {
			lb.Backends[i].Alive = false
		}

		resp.Body.Close()
	}
}

func (lb *LoadBalancer) NextBackend() (Backend, error) {

	n := len(lb.Backends)

	for i := 0; i < n; i++ {

		idx := (lb.Current + i) % n

		if lb.Backends[idx].Alive {
			lb.Current = (idx + 1) % n

			return lb.Backends[idx], nil
		}
	}

	return Backend{}, errors.New("no healthy backends available")

}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	backend, err := lb.NextBackend()
	if err != nil {
		http.Error(w, "No healthy backends available", http.StatusServiceUnavailable)
		return
	}
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

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			lb.HealthCheck()
		}
	}()
	lb.HealthCheck()
	if err := http.ListenAndServe(":8080", &lb); err != nil {
		panic(err)
	}
}
