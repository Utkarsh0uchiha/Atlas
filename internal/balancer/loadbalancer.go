package balancer

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/Utkarsh0uchiha/go-load-balancer/internal/backend"
	"github.com/Utkarsh0uchiha/go-load-balancer/internal/events"
)

type LoadBalancer struct {
	Backends      []backend.Backend
	Current       int
	client        *http.Client
	mu            sync.RWMutex
	totalRequests int
	startTime     time.Time
	broker        *events.Broker
}

func New(backend []backend.Backend) *LoadBalancer {
	return &LoadBalancer{
		Backends: backend,
		Current:  0,
		client: &http.Client{
			Timeout: 2 * time.Second,
		},
		totalRequests: 0,
		startTime:     time.Now(),
	}
}

func (lb *LoadBalancer) SetBroker(broker *events.Broker) {
	lb.broker = broker
}

func (lb *LoadBalancer) BroadcastStatus() {
	status := lb.GetStatus()

	data, err := json.Marshal(status)

	if err != nil {
		return
	}

	lb.broker.Broadcast(data)

}
