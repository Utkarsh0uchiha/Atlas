package balancer

import (
	"errors"

	"github.com/Utkarsh0uchiha/go-load-balancer/internal/backend"
)

func (lb *LoadBalancer) NextBackend() (backend.Backend, error) {

	n := len(lb.Backends)

	for i := 0; i < n; i++ {

		idx := (lb.Current + i) % n

		if lb.Backends[idx].Alive {
			lb.Current = (idx + 1) % n

			return lb.Backends[idx], nil
		}
	}

	return backend.Backend{}, errors.New("no healthy backends available")

}
