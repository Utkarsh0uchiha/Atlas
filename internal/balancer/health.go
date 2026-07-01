package balancer

import (
	"net/http"
)

func (lb *LoadBalancer) HealthCheck() {

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

	}
}
