
## 📊 Monitoring & Observability

Atlas includes a complete observability stack powered by **Prometheus** and **Grafana**, enabling real-time monitoring of system health, request traffic, and backend performance.

### Monitoring Stack

```text
             Atlas
               │
               │ /metrics
               │
               ▼
           Prometheus
               │
               ▼
        Grafana Dashboard
```

Prometheus periodically scrapes metrics exposed by Atlas and stores them as time-series data. Grafana connects to Prometheus to visualize operational metrics through interactive dashboards.

---

## 📈 Exported Metrics

Atlas exposes Prometheus metrics at:

```text
http://localhost:8080/metrics
```

| Metric | Type | Description |
|--------|------|-------------|
| `load_balancer_requests_total` | Counter | Total number of requests handled by Atlas |
| `load_balancer_backend_requests_total` | Counter | Requests forwarded to each backend |
| `load_balancer_healthy_backends` | Gauge | Number of currently healthy backends |
| `load_balancer_request_duration_seconds` | Histogram | Request latency distribution |

---

## 📉 Grafana Dashboard

The included Grafana dashboard provides real-time insights into system behavior.

### Dashboard Panels

- 🟢 Healthy Backends
- 📈 Request Rate
- ⚖️ Backend Request Distribution
- ⏱ P95 Request Latency

---

### Healthy Backends

Displays the number of healthy backend servers currently participating in request routing.

> ![Healthy Backends](docs/images/healthy-backends.jpg)

---

### Request Rate

Shows the incoming request throughput (requests/second).

> ![Request Rate](docs/images/request-rate.jpg)

---

### Backend Request Distribution

Visualizes how requests are distributed across backend servers.

> ![Backend Request Distribution](docs/images/backend-request-distribution.jpg)

---

### P95 Request Latency

Tracks the 95th percentile request latency, helping identify performance regressions.

> ![P95 Request Latency](docs/images/p95-latency.jpg)
