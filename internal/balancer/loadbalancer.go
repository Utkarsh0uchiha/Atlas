package balancer

import (
	"net/http"
	"sync"
	"time"

	"github.com/Utkarsh0uchiha/go-load-balancer/internal/backend"
)

type LoadBalancer struct {
	Backends      []backend.Backend
	Current       int
	client        *http.Client
	mu            sync.RWMutex
	totalRequests int
	startTime     time.Time
}

func New(backend []backend.Backend) *LoadBalancer {
	return &LoadBalancer{
		Backends: backend,
		Current:  0,
		client: &http.Client{
			Timeout: 2 * time.Second,
		},
		totalRequests: 0,
		startTime: time.Now(),
	}
}
