# Atlas Architecture

## Overview

Atlas is a production-inspired HTTP load balancer written in Go that distributes incoming requests across multiple backend servers using the **Round Robin** scheduling algorithm.

The system is designed with modularity, reliability, and observability in mind. It continuously monitors backend health, automatically reroutes traffic when failures occur, and exposes operational metrics through Prometheus for real-time visualization in Grafana.

---

## High-Level Architecture

```text
                                    ┌─────────────────────┐
                                    │       Client        │
                                    └──────────┬──────────┘
                                               │
                                           HTTP Request
                                               │
                                               ▼
                            ┌────────────────────────────────────┐
                            │               Atlas                │
                            │        Go HTTP Load Balancer       │
                            └────────────────────────────────────┘
                                               │
                 ┌─────────────────────────────┼─────────────────────────────┐
                 │                             │                             │
                 ▼                             ▼                             ▼
        ┌─────────────────┐          ┌─────────────────┐          ┌─────────────────┐
        │   Backend #1    │          │   Backend #2    │          │   Backend #3    │
        │     :8081       │          │     :8082       │          │     :8083       │
        └─────────────────┘          └─────────────────┘          └─────────────────┘
                 ▲                             ▲                             ▲
                 └───────────────┬─────────────┴───────────────┬─────────────┘
                                               │
                                     Active Health Checks (HTTP)
                                               │
                                               ▼
                                   Healthy Backend Registry
                                               │
                                               ▼
                                     Prometheus Metrics
                                               │
                                               ▼
                                            Grafana
```

---

# Core Components

## Client

The client sends HTTP requests to Atlas instead of communicating directly with backend servers.

Atlas acts as the single entry point for all incoming traffic.

---

## Load Balancer

The load balancer is responsible for:

- Receiving client requests
- Selecting a healthy backend
- Forwarding requests using Go's reverse proxy
- Handling backend failures
- Recording operational metrics

---

## Backend Servers

Backend servers provide the actual application responses.

Each backend operates independently and can become healthy or unhealthy without affecting the availability of the remaining servers.

Current implementation:

| Backend | Port |
|---------|------|
| Backend 1 | 8081 |
| Backend 2 | 8082 |
| Backend 3 | 8083 |

---

## Health Checker

Atlas periodically performs HTTP health checks against every backend.

If a backend:

- responds successfully → Healthy
- fails to respond → Unhealthy

Only healthy backends participate in request routing.

---

## Metrics Exporter

Atlas exposes Prometheus-compatible metrics through the `/metrics` endpoint.

Metrics include:

- Total requests
- Backend request count
- Healthy backend count
- Request latency histogram

---

## Prometheus

Prometheus periodically scrapes metrics from Atlas and stores them as time-series data.

---

## Grafana

Grafana connects to Prometheus and visualizes operational metrics through dashboards.

The dashboard provides visibility into:

- Healthy backends
- Request throughput
- Backend traffic distribution
- P95 latency

---

# Request Flow

```text
                         Client
                            │
                            ▼
                  Atlas receives request
                            │
                            ▼
                  Round Robin Scheduler
                            │
                            ▼
                 Select next healthy backend
                            │
                            ▼
               Reverse Proxy forwards request
                            │
                            ▼
                Backend processes request
                            │
                            ▼
               Response returned to client
```

---

# Reverse Proxy Workflow

Atlas uses Go's `httputil.NewSingleHostReverseProxy` to transparently forward client requests.

The reverse proxy:

- preserves HTTP semantics
- forwards request headers
- streams responses back to clients
- handles backend communication

Clients remain unaware of which backend ultimately processes the request.

---

# Backend Selection

Atlas currently uses the **Round Robin** scheduling algorithm.

Each healthy backend receives requests in turn, ensuring an even distribution of traffic.

Example:

```text
Request 1 → Backend 1

Request 2 → Backend 2

Request 3 → Backend 3

Request 4 → Backend 1
```

If a backend becomes unavailable, it is skipped until it successfully passes a future health check.

---

# Fault Tolerance

Atlas is designed to continue serving traffic even when one or more backend servers become unavailable.

Failure handling process:

1. Backend fails.
2. Health checker detects failure.
3. Backend marked unhealthy.
4. Backend removed from rotation.
5. Remaining healthy backends continue serving requests.
6. Backend automatically rejoins after recovery.

No manual intervention is required.

---

# Monitoring Pipeline

```text
Atlas
  │
  ├── /metrics
  ▼
Prometheus
  │
  ▼
Grafana Dashboard
```

Operational metrics are continuously exported and visualized, providing real-time insight into system health and performance.

---

# Design Goals

Atlas was built with the following engineering goals:

- Simplicity
- Reliability
- Extensibility
- Observability
- Containerized deployment
- Production-inspired architecture

Future enhancements such as weighted scheduling, service discovery, and circuit breakers can be added without significant architectural changes.