package balancer

import (
	"net/http"

	"github.com/Utkarsh0uchiha/go-load-balancer/internal/metrics"
)

func (lb *LoadBalancer) HealthCheck() {
	healthy := 0
	for i := range lb.Backends {
		resp, err := lb.client.Get(lb.Backends[i].URL.String() + "/health")

		alive := false
		if err == nil {
			alive = resp.StatusCode == http.StatusOK

			resp.Body.Close()
		}

		lb.mu.Lock()
		lb.Backends[i].Alive = alive
		lb.mu.Unlock()

		if alive {
			healthy++
		}

	}

	metrics.HealthyBackends.Set(float64(healthy))
}
