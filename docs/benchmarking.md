
## ⚡ Performance

Atlas is designed to provide reliable request distribution while maintaining low latency and high availability.

Performance testing was conducted locally using HTTP benchmarking tools against the Docker Compose deployment.

> **Note**
> Benchmark results may vary depending on hardware, operating system, and network conditions.

### Benchmark Configuration

| Parameter | Value |
|-----------|-------|
| Backend Servers | 3 |
| Algorithm | Round Robin |
| Monitoring | Prometheus + Grafana |
| Deployment | Docker Compose |

### Sample Benchmark

```bash
wrk -t4 -c100 -d30s http://localhost:8080
```

or

```bash
hey -n 10000 -c 100 http://localhost:8080
```
