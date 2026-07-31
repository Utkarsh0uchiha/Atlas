# Atlas

![Go](https://img.shields.io/badge/Go-1.25+-00ADD8?logo=go)
![Docker](https://img.shields.io/badge/Docker-Compose-2496ED?logo=docker)
![Prometheus](https://img.shields.io/badge/Prometheus-Monitoring-E6522C?logo=prometheus)
![Grafana](https://img.shields.io/badge/Grafana-Dashboard-F46800?logo=grafana)
![License](https://img.shields.io/badge/License-MIT-green)

**Atlas** is a production-inspired HTTP load balancer built in Go that intelligently distributes traffic across backend servers while ensuring high availability through automatic health checks and failover.

Designed to explore production-grade backend engineering concepts including reverse proxying, concurrency, observability, and containerized deployments, Atlas demonstrates how modern load balancers operate in real-world distributed systems.

---

## Why Atlas?

Most tutorials stop after implementing a basic reverse proxy. Atlas goes further by incorporating the engineering practices used in production systems, including health monitoring, automatic failover, metrics collection, and containerized deployment.

The project was built to gain hands-on experience with backend infrastructure, distributed systems concepts, and production-grade tooling rather than simply forwarding HTTP requests.

---

## Features

### ⚖️ Load Balancing

- Reverse proxy built with Go's `net/http/httputil`
- Round Robin request scheduling
- Intelligent backend selection

### 🛡 Reliability

- Active health checks
- Passive health checks
- Automatic retry on backend failure
- Automatic failover

### ⚡ Concurrency

- Thread-safe backend management using `sync.RWMutex`

### 📊 Observability

- Prometheus metrics
- Grafana dashboards
- Request latency histogram
- Backend traffic monitoring
- Health status visualization

### 🐳 DevOps

- Multi-stage Docker builds
- Docker Compose orchestration
- Environment-based configuration

## 📖 Table of Contents

- [Why Atlas?](#-why-atlas)
- [Features](#-features)
- [Architecture](#-architecture)
- [Tech Stack](#-tech-stack)
- [Project Structure](#-project-structure)
- [Core Features](#-core-features)
- [Getting Started](#-getting-started)
- [Monitoring & Metrics](#-monitoring-metrics)
- [Grafana Dashboard](#-grafana-dashboard)
- [Benchmarks](#-benchmarks)
- [Roadmap](#-roadmap)
- [Contributing](#-contributing)
- [License](#-license)

## 🏗️ Architecture

```text
                        ┌──────────────────────┐
                        │       Client         │
                        └──────────┬───────────┘
                                   │
                              HTTP Request
                                   │
                                   ▼
                  ┌────────────────────────────────┐
                  │             Atlas              │
                  │      Go HTTP Load Balancer     │
                  └────────────────────────────────┘
                                   │
                ┌──────────────────┼──────────────────┐
                │                  │                  │
                ▼                  ▼                  ▼
         ┌────────────┐     ┌────────────┐     ┌────────────┐
         │ Backend 1  │     │ Backend 2  │     │ Backend 3  │
         │   :8081    │     │   :8082    │     │   :8083    │
         └────────────┘     └────────────┘     └────────────┘
                ▲                  ▲                  ▲
                └──────────┬───────┴──────────┬───────┘
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

### Request Lifecycle

1. A client sends an HTTP request to Atlas.
2. Atlas selects the next healthy backend using the Round Robin algorithm.
3. The request is forwarded through Go's reverse proxy.
4. If the backend fails to respond, Atlas marks it as unhealthy and retries the request using another healthy backend.
5. Periodic health checks determine when an unhealthy backend becomes available again.
6. Metrics are exported to Prometheus and visualized through Grafana in real time.

## 🛠 Tech Stack

Atlas is built using modern backend technologies and follows production-inspired engineering practices.

| Category | Technology |
|----------|------------|
| **Language** | Go 1.25+ |
| **HTTP Server** | `net/http` |
| **Reverse Proxy** | `net/http/httputil` |
| **Concurrency** | `sync.RWMutex`, Goroutines |
| **Scheduling** | Round Robin |
| **Health Monitoring** | Active & Passive Health Checks |
| **Observability** | Prometheus |
| **Visualization** | Grafana |
| **Containerization** | Docker |
| **Orchestration** | Docker Compose |
| **Version Control** | Git & GitHub |

## 📁 Project Structure

```text
atlas/
├── cmd/
│   ├── proxy/                 # Atlas load balancer
│   ├── server1/               # Backend server 1
│   ├── server2/               # Backend server 2
│   └── server3/               # Backend server 3
│
├── internal/
│   ├── backend/               # Backend representation
│   ├── loadbalancer/          # Round Robin scheduler
│   ├── metrics/               # Prometheus metrics
│   └── proxy/                 # Reverse proxy implementation
│
├── prometheus/
│   └── prometheus.yml         # Prometheus configuration
│
├── grafana/                   # Grafana provisioning (future)
│
├── Dockerfile.proxy
├── Dockerfile.server1
├── Dockerfile.server2
├── Dockerfile.server3
├── docker-compose.yml
│
├── go.mod
├── go.sum
│
└── README.md
```

## ⚙️ Core Features

### ⚖️ Intelligent Request Distribution

Atlas distributes incoming requests using the **Round Robin** scheduling algorithm, ensuring traffic is evenly balanced across all healthy backend servers.

---

### ❤️ Active Health Checks

Atlas continuously probes backend servers at configurable intervals. If a backend fails a health check, it is automatically removed from the rotation until it becomes healthy again.

---

### 🔁 Automatic Failover

When a request to a backend fails unexpectedly, Atlas immediately retries the request on another healthy backend, ensuring high availability with minimal disruption.

---

### 🔒 Thread-Safe Backend Management

Backend state is protected using Go's `sync.RWMutex`, allowing concurrent request handling while safely updating backend health information.

---

### 📊 Built-in Observability

Atlas exports Prometheus metrics for:

- Total requests
- Backend request distribution
- Healthy backend count
- Request latency histogram

These metrics are visualized through Grafana dashboards.

---

### 🐳 Containerized Deployment

Every service runs inside Docker containers and is orchestrated using Docker Compose, making the entire stack reproducible with a single command.

---

## 🚀 Getting Started

### Prerequisites

Before running Atlas, ensure the following tools are installed:

- Go **1.25+**
- Docker
- Docker Compose
- Git

Verify your installation:

```bash
go version
docker --version
docker compose version
```

---

## 📥 Clone the Repository

```bash
git clone https://github.com/Utkarsh0uchiha/atlas.git
cd atlas
```
---

## 🏗️ Build and Run

Start the complete stack using Docker Compose:

```bash
docker compose up --build
```

To run in detached mode:

```bash
docker compose up --build -d
```

Stop all services:

```bash
docker compose down
```

---

## 🌐 Services

Once Atlas is running, the following services are available:

| Service | URL |
|----------|-----|
| Atlas Load Balancer | http://localhost:8080 |
| Backend Server 1 | http://localhost:8081 |
| Backend Server 2 | http://localhost:8082 |
| Backend Server 3 | http://localhost:8083 |
| Prometheus | http://localhost:9090 |
| Grafana | http://localhost:3000 |

---

## 📚 Documentation

Detailed documentation is available in the `docs/` directory.

- [Architecture](docs/architecture.md)
- [Monitoring & Observability](docs/monitoring.md)
- [Deployment Guide](docs/deployment.md)
- [Health Checks](docs/health-checks.md)
- [Benchmarking](docs/benchmarking.md)
---

## 🛣️ Roadmap

Atlas is actively evolving. Planned improvements include:

### Load Balancing

- [ ] Weighted Round Robin
- [ ] Least Connections
- [ ] IP Hash
- [ ] Random Selection

### Reliability

- [ ] Circuit Breaker
- [ ] Graceful Shutdown
- [ ] Request Timeout Configuration
- [ ] Connection Pooling

### Service Discovery

- [ ] Dynamic backend registration
- [ ] Configuration file support
- [ ] DNS-based service discovery

### Dashboard

- [ ] Web management interface
- [ ] Enable/Disable backends
- [ ] Live request monitoring
- [ ] Health status indicators

### Observability

- [ ] Grafana provisioning
- [ ] Alertmanager integration
- [ ] Distributed tracing
- [ ] Structured logging

### DevOps

- [ ] GitHub Actions CI
- [ ] Automated testing
- [ ] Kubernetes deployment
- [ ] Helm Chart

---

## 🤝 Contributing

Contributions are welcome!

If you'd like to contribute:

1. Fork the repository.
2. Create a feature branch.

```bash
git checkout -b feature/amazing-feature
```

3. Commit your changes.

```bash
git commit -m "feat: add amazing feature"
```

4. Push your branch.

```bash
git push origin feature/amazing-feature
```

5. Open a Pull Request.

Please ensure all changes are tested before submitting.

---

## 📜 License

This project is licensed under the MIT License.

See the `LICENSE` file for more information.

---

## 👨‍💻 Author

**Utkarsh Raj**

- GitHub: https://github.com/Utkarsh0uchiha
- LinkedIn: https://www.linkedin.com/in/utkarsh0uchiha/

If you found this project helpful, consider giving it a ⭐ on GitHub.