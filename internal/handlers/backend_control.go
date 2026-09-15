package handlers

import (
	"net/http"
	"strconv"

	"github.com/Utkarsh0uchiha/go-load-balancer/internal/balancer"
)

func NewDisableBackendHandler(lb *balancer.LoadBalancer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", "POST")

			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		err = lb.DisableBackend(id)
		if err != nil {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message" : "backend disabled"}`))

	}
}

func NewEnableBackendHandler(lb *balancer.LoadBalancer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", "POST")
			http.Error(w, "Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		idStr := r.PathValue("id")

		id, err := strconv.Atoi(idStr)

		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		err = lb.EnableBackend(id)
		if err != nil {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message" : "backend enabled"}`))

	}
}
