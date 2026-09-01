package balancer

import "time"

type BackendStatus struct {
	ID      int    `json:"id"`
	URL     string `json:"url"`
	Healthy bool   `json:"healthy"`
}

type Status struct {
	TotalRequests   int             `json:"total_requests"`
	HealthyBackends int             `json:"healthy_backends"`
	TotalBackends   int             `json:"total_backends"`
	Algorithm       string          `json:"algorithm"`
	Uptime          string          `json:"uptime"`
	Backends        []BackendStatus `json:"backends"`
}

func (lb *LoadBalancer) GetStatus() Status {

	lb.mu.RLock()
	totalrequest := lb.totalRequests
	totalbackend := len(lb.Backends)
	alive := 0
	status := make([]BackendStatus, totalbackend)

	for i := range lb.Backends {
		if lb.Backends[i].Alive {
			alive++
		}
		snapshot := BackendStatus{
			ID:      i,
			URL:     lb.Backends[i].URL.String(),
			Healthy: lb.Backends[i].Alive,
		}

		status[i] = snapshot
	}

	lb.mu.RUnlock()

	totalsnapshot := Status{
		TotalRequests:   totalrequest,
		HealthyBackends: alive,
		TotalBackends:   totalbackend,
		Algorithm:       "round_robin",
		Uptime:          time.Since(lb.startTime).Round(time.Second).String(),
		Backends:        status,
	}

	return totalsnapshot
}
