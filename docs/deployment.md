# Deployment Guide

This document explains how to deploy and run **Atlas** in different environments. The recommended approach is Docker Compose, which provisions the load balancer, backend servers, Prometheus, and Grafana with a single command.

---

# Deployment Options

Atlas currently supports the following deployment methods:

- Local Development (Go)
- Docker
- Docker Compose (Recommended)

Future releases will include Kubernetes deployment support.

---

# Prerequisites

Ensure the following software is installed before deployment.

| Tool | Version |
|------|---------|
| Go | 1.25+ |
| Docker | Latest |
| Docker Compose | Latest |
| Git | Latest |

Verify your installation:

```bash
go version
docker --version
docker compose version
```

---

# Local Development

Clone the repository.

```bash
git clone https://github.com/Utkarsh0uchiha/atlas.git
cd atlas
```

Start the backend servers.

```bash
go run ./cmd/server1
```

```bash
go run ./cmd/server2
```

```bash
go run ./cmd/server3
```

Start Atlas.

```bash
go run ./cmd/proxy
```

Atlas will be available at:

```
http://localhost:8080
```

---

# Docker Deployment

Build the Docker images.

```bash
docker build -f Dockerfile.proxy -t atlas-proxy .
```

```bash
docker build -f Dockerfile.server1 -t atlas-server1 .
```

```bash
docker build -f Dockerfile.server2 -t atlas-server2 .
```

```bash
docker build -f Dockerfile.server3 -t atlas-server3 .
```

Run the containers individually if required.

---

# Docker Compose Deployment

The recommended deployment method is Docker Compose.

Build and start all services.

```bash
docker compose up --build
```

Run in detached mode.

```bash
docker compose up -d --build
```

Stop the deployment.

```bash
docker compose down
```

Remove containers and volumes.

```bash
docker compose down -v
```

---

# Services

Once deployed, the following services are available.

| Service | URL |
|----------|-----|
| Atlas | http://localhost:8080 |
| Backend Server 1 | http://localhost:8081 |
| Backend Server 2 | http://localhost:8082 |
| Backend Server 3 | http://localhost:8083 |
| Prometheus | http://localhost:9090 |
| Grafana | http://localhost:3000 |

### Docker Compose Services

![Docker Compose Services](https://github.com/Utkarsh0uchiha/go-load-balancer/blob/main/docs/image/docker-compose-services.png?raw=true)
---

# Docker Compose Architecture

```text
                    Docker Network
                           │
        ┌──────────────────┼──────────────────┐
        │                  │                  │
        ▼                  ▼                  ▼
     Atlas             Prometheus         Grafana
        │
        ├──────────────┬──────────────┐
        ▼              ▼              ▼
   Backend 1      Backend 2      Backend 3
```

All containers communicate through Docker's internal bridge network using service names as hostnames.

---

# Environment Configuration

Atlas supports configuration through environment variables.

Example:

```env
PORT=8080
HEALTH_CHECK_INTERVAL=5s
```

Future versions may include:

- Backend configuration via environment variables
- Configuration files (YAML/TOML)
- Dynamic service discovery

---

# Deployment Verification

After deployment, verify that all services are running.

Check Docker containers.

```bash
docker compose ps
```

Verify Atlas.

```bash
curl http://localhost:8080
```

Verify Prometheus.

```
http://localhost:9090
```

Navigate to:

```
Status → Targets
```

Ensure the Atlas target is **UP**.

Verify Grafana.

```
http://localhost:3000
```

Open the Atlas dashboard and confirm that metrics are updating.

---

# Simulating Backend Failure

Stop one backend.

```bash
docker compose stop server2
```

Atlas should automatically:

- Detect the failure
- Remove the backend from rotation
- Continue serving traffic through healthy backends

Restart the backend.

```bash
docker compose start server2
```

Once health checks succeed, Atlas automatically restores the backend.

---

# Production Considerations

Atlas is intended as a production-inspired educational project.

For production deployments, consider implementing:

- HTTPS/TLS termination
- Authentication and authorization
- Rate limiting
- Request timeouts
- Graceful shutdown
- Structured logging
- Centralized configuration
- Dynamic service discovery
- Kubernetes deployment
- CI/CD pipeline

---

# Troubleshooting

## Docker Compose fails to start

Rebuild the images.

```bash
docker compose down
docker compose up --build
```

---

## Prometheus target is DOWN

Verify:

- Atlas is running.
- The `/metrics` endpoint is reachable.
- `prometheus.yml` contains the correct target.

---

## Grafana shows "No Data"

Verify:

- Prometheus is running.
- Grafana datasource points to Prometheus.
- Atlas has received traffic.

Generate sample traffic.

```bash
curl http://localhost:8080
```

or

```bash
hey -n 1000 -c 50 http://localhost:8080
```

---

# Future Improvements

- Kubernetes manifests
- Helm charts
- Multi-node deployment
- Horizontal scaling
- Rolling updates
- Zero-downtime deployments
- Automated infrastructure provisioning

