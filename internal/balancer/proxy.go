package balancer

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func (lb *LoadBalancer) proxyToBackend(w http.ResponseWriter, r *http.Request, idx int, attempts int) {
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
	log.Printf("Retrying on backend %s", backend2.URL)
	retryProxy := httputil.NewSingleHostReverseProxy(backend2.URL)
	retryProxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, proxyError error) {
		lb.proxyToBackend(w, r, idx2, attempts+1)
	}
	retryProxy.ServeHTTP(w, r)
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	backend, idx, err := lb.NextBackend()
	if err != nil {
		http.Error(w, "503 Service Unavailable", http.StatusServiceUnavailable)
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(backend.URL)
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, proxyError error) {
		lb.proxyToBackend(w, r, idx, 1)
	}
	proxy.ServeHTTP(w, r)
}
