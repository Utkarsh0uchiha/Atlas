package balancer

import (
	"net/http"
)
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