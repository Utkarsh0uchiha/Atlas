package balancer

import (
	"net/http/httputil"
	"net/http"
)
func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	backend, err := lb.NextBackend()
	if err != nil {
		http.Error(w, "No healthy backends available", http.StatusServiceUnavailable)
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(backend.URL)

	proxy.ServeHTTP(w, r)
}