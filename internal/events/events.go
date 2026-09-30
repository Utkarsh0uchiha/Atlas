package events

import "sync"

type Broker struct {
	clients map[chan []byte]struct{}
	mu      sync.RWMutex
}

func NewBroker() *Broker {
	return &Broker{
		clients: make(map[chan []byte]struct{}),
	}
}

func (b *Broker) Subscribe() chan []byte {
	b.mu.Lock()
	defer b.mu.Unlock()
	client := make(chan []byte, 10)

	b.clients[client] = struct{}{}

	return client
}

func (b *Broker) Unsubscribe(client chan []byte) {
	b.mu.Lock()
	defer b.mu.Unlock()

	delete(b.clients, client)

	return
}

func (b *Broker) Broadcast(data []byte) {

	b.mu.RLock()
	clients := make([]chan []byte, 0, len(b.clients))
	for ch := range b.clients {
		clients = append(clients, ch)
	}
	b.mu.RUnlock()

	for _, ch := range clients {
		ch <- data
	}
}
