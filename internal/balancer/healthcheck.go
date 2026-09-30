package balancer

import (
	"time"

	"github.com/Utkarsh0uchiha/go-load-balancer/internal/events"
)

func (lb *LoadBalancer) StartHealthChecker(interval time.Duration, broker *events.Broker) {

	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for range ticker.C {
			lb.HealthCheck(broker)
		}
	}()
}
