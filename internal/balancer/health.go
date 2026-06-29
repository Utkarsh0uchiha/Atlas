package balancer

import (
	"net/http"
)

func (lb *LoadBalancer) HealthCheck() {

	for i := range lb.Backends {
		resp, err := http.Get(lb.Backends[i].URL.String())

		Alive := false
		if err != nil {
			Alive = resp.StatusCode == http.StatusOK

			resp.Body.Close()
		}

		lb.mu.Lock()
		lb.Backends[i].Alive = Alive
		lb.mu.Unlock()

	}
}
