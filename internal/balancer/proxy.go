package balancer

import (
	"log"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/Utkarsh0uchiha/go-load-balancer/internal/metrics"
)

func (lb *LoadBalancer) serveWithRetry(w http.ResponseWriter, r *http.Request, idx int, attempts int) {
	if attempts >= len(lb.Backends) {
		http.Error(w, "503 Service Unavailable", http.StatusServiceUnavailable)
		return
	}
	lb.mu.Lock()
	lb.Backends[idx].Alive = false
	log.Printf("Backend %s marked unhealthy", lb.Backends[idx].URL)
	lb.mu.Unlock()
	backend2, idx2, err := lb.NextBackend()
	if err != nil {
		http.Error(w, "503 Service Unavailable", http.StatusServiceUnavailable)
		return
	}
	failed := false
	log.Printf("Retrying on backend %s", backend2.URL)
	retryProxy := httputil.NewSingleHostReverseProxy(backend2.URL)
	retryProxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, proxyError error) {
		failed = true
		lb.serveWithRetry(w, r, idx2, attempts+1)
	}

	backendStart := time.Now()
	retryProxy.ServeHTTP(w, r)
	backendLatency := time.Since(backendStart)
	if !failed {
		metrics.BackendRequestsTotal.WithLabelValues(backend2.URL.Host).Inc()
		lb.mu.Lock()
		lb.Backends[idx2].Requests++
		lb.Backends[idx2].TotalLatency += backendLatency
		lb.mu.Unlock()

	}
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	defer func() {
		metrics.RequestDuration.Observe(time.Since(start).Seconds())
	}()
	lb.mu.Lock()
	lb.totalRequests++
	lb.mu.Unlock()
	metrics.RequestsTotal.Inc()
	backend, idx, err := lb.NextBackend()
	if err != nil {
		http.Error(w, "503 Service Unavailable", http.StatusServiceUnavailable)
		return
	}
	failed := false
	proxy := httputil.NewSingleHostReverseProxy(backend.URL)
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, proxyError error) {
		failed = true
		lb.serveWithRetry(w, r, idx, 1)
	}

	backendStart := time.Now()
	proxy.ServeHTTP(w, r)
	backendLatency := time.Since(backendStart)
	if !failed {
		metrics.BackendRequestsTotal.WithLabelValues(backend.URL.Host).Inc()
		lb.mu.Lock()
		lb.Backends[idx].Requests++
		lb.Backends[idx].TotalLatency += backendLatency
		lb.mu.Unlock()
	}

}
