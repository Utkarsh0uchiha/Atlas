package balancer

import "time"

func (lb *LoadBalancer) StartHealthChecker(interval time.Duration) {

	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for range ticker.C {
			lb.HealthCheck()
		}
	}()
}
