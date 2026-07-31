## 🔍 Health Checks

Atlas performs periodic HTTP health checks against every backend server.

When a backend:

- responds successfully → it remains **Healthy**
- becomes unreachable → it is marked **Unhealthy**
- recovers → it automatically rejoins the load-balancing pool

This ensures uninterrupted traffic routing without manual intervention.

---

## 🚦 Failover Demonstration

When a backend server is stopped:

- Atlas detects the failure during health checks.
- The backend is removed from the rotation.
- Incoming traffic is automatically redirected to healthy servers.
- Grafana immediately reflects the updated backend health and request distribution.

Once the backend is restarted and passes health checks, it is automatically added back into the rotation.


