package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Utkarsh0uchiha/go-load-balancer/internal/balancer"
)

func NewStatusHandler(lb *balancer.LoadBalancer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := lb.GetStatus()

		w.Header().Set("Content-Type", "application/json")
		
		err := json.NewEncoder(w).Encode(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}
