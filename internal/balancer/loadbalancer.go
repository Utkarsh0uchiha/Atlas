package balancer

import "github.com/Utkarsh0uchiha/go-load-balancer/internal/backend"

type LoadBalancer struct {
	Backends []backend.Backend
	Current  int
}

func New(backend []backend.Backend) *LoadBalancer{
	return &LoadBalancer{
		Backends: backend,
		Current: 0,
	}
}
