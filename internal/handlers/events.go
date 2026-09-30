package handlers

import (
	"fmt"
	"net/http"

	"github.com/Utkarsh0uchiha/go-load-balancer/internal/events"
)

func NewSSEHandler(broker *events.Broker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ch := broker.Subscribe()

		w.Header().Set("Content-Type", "text/event-stream")

		flusher, ok := w.(http.Flusher)

		if !ok {
			broker.Unsubscribe(ch)
			http.Error(w, "Streaming Unsupported", http.StatusInternalServerError)
			return
		}

		for {
			select {
			case data := <-ch:

				fmt.Fprintf(w, "data: %s\n\n", data)
				flusher.Flush()
			case <-r.Context().Done():
				broker.Unsubscribe(ch)
				return
			}
		}

	}
}
