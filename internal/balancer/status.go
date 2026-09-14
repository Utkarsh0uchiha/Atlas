package balancer

import "time"

type BackendStatus struct {
	ID             int    `json:"id"`
	URL            string `json:"url"`
	Healthy        bool   `json:"healthy"`
	Requests       int    `json:"requests"`
	AverageLatency string `json:"average_latency"`
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
		average := time.Duration(0)

		requests := lb.Backends[i].Requests
		if requests > 0 {
			average = lb.Backends[i].TotalLatency / time.Duration(requests)
		}
		snapshot := BackendStatus{
			ID:             i,
			URL:            lb.Backends[i].URL.String(),
			Healthy:        lb.Backends[i].Alive,
			Requests:       requests,
			AverageLatency: average.String(),
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
