package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "load_balancer_requests_total",
			Help: "Total number of requests received by load balancer",
		},
	)
	BackendRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "load_balancer_backend_requests_total",
			Help: "Total number of requests received by each backends",
		},
		[]string{"backend"},
	)

	HealthyBackends = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "load_balancer_healthy_backends",
			Help: "Total number of healthy backends",
		},
	)

	RequestDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "load_balancer_request_duration_seconds",
			Help:    "Duration of HTTP requests handled by the load balancer",
			Buckets: prometheus.DefBuckets,
		},
	)
)

func init() {
	prometheus.MustRegister(
		RequestsTotal,
		BackendRequestsTotal,
		HealthyBackends,
		RequestDuration,
	)
}
