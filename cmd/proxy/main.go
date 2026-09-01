package main

import (
	"html/template"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/Utkarsh0uchiha/go-load-balancer/internal/backend"
	"github.com/Utkarsh0uchiha/go-load-balancer/internal/balancer"
	"github.com/Utkarsh0uchiha/go-load-balancer/internal/handlers"
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

	backend3 := os.Getenv("BACKEND_3")
	if backend3 == "" {
		backend3 = "http://localhost:8083"
	}
	b1url, err := url.Parse(backend1)
	if err != nil {
		panic(err)
	}
	b2url, err := url.Parse(backend2)
	if err != nil {
		panic(err)
	}

	b3url, err := url.Parse(backend3)
	if err != nil {
		panic(err)
	}
	backends := []backend.Backend{
		{URL: b1url, Alive: true},
		{URL: b2url, Alive: true},
		{URL: b3url, Alive: true},
	}

	tmpl, err := template.ParseFiles("web/templates/dashboard.html")
	if err != nil {
		panic(err)
	}

	lb := balancer.New(backends)
	lb.HealthCheck()
	lb.StartHealthChecker(5 * time.Second)

	statusHandler := handlers.NewStatusHandler(lb)
	filesystem := http.Dir("web/static")

	fileserver := http.FileServer(filesystem)

	handler := http.StripPrefix("/static", fileserver)

	http.Handle("/", lb)
	http.HandleFunc("/dashboard", handlers.NewDashboardHandler(tmpl))
	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/static/", handler)
	http.Handle("/api/status", statusHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
