package main

import (
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/Utkarsh0uchiha/go-load-balancer/internal/backend"
	"github.com/Utkarsh0uchiha/go-load-balancer/internal/balancer"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	backend1 := os.Getenv("BACKEND_1")
	if backend1 == "" {
		backend1 = "http://localhost:8081"
	}
	backend2 := os.Getenv("BACKEND_2")
	if backend2 == "" {
		backend2 = "http://localhost:8082"
	}
	b1url, err := url.Parse(backend1)
	if err != nil {
		panic(err)
	}
	b2url, err := url.Parse(backend2)
	if err != nil {
		panic(err)
	}
	backends := []backend.Backend{
		{URL: b1url, Alive: true},
		{URL: b2url, Alive: true},
	}

	lb := balancer.New(backends)
	lb.HealthCheck()
	lb.StartHealthChecker(5 * time.Second)
	http.Handle("/", lb)
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
