package main

import (
	"net/http"
	"net/url"
	"time"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/Utkarsh0uchiha/go-load-balancer/internal/backend"
	"github.com/Utkarsh0uchiha/go-load-balancer/internal/balancer"
)

func main() {
	backend1, err := url.Parse("http://localhost:8081")
	if err != nil {
		panic(err)
	}
	backend2, err := url.Parse("http://localhost:8082")
	if err != nil {
		panic(err)
	}

	backends := []backend.Backend{
		{URL: backend1, Alive: true},
		{URL: backend2, Alive: true},
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
