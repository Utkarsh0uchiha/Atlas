package balancer

import "errors"

func (lb *LoadBalancer) DisableBackend(id int) error {
	if id < 0 || id >= len(lb.Backends) {
		return errors.New("backend not found")
	}

	lb.mu.Lock()
	lb.Backends[id].Enabled = false
	lb.mu.Unlock()

	return nil
}
func (lb *LoadBalancer) EnableBackend(id int) error {
	if id < 0 || id >= len(lb.Backends) {
		return errors.New("backend not found")
	}

	lb.mu.Lock()
	lb.Backends[id].Enabled = true
	lb.mu.Unlock()

	return nil
}
