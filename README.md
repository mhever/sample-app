# sample-app — tiny Go HTTP service

This repository contains a minimal Go HTTP API intended for deployment to a small Kubernetes cluster. It uses `gin` and exposes a few endpoints useful for demos, probes and Prometheus scraping.

Endpoints
- GET / — service name, version, hostname, time
- GET /healthz — liveness probe (always 200 if process alive)
- GET /readyz — readiness probe (returns 200 only when DB reachable)
- GET /env — shows a small subset of config values (from `ConfigMap`)
- GET /work?ms=200 — sleeps N ms to simulate latency
- GET /metrics — Prometheus metrics

Build & run (local)

```bash
go mod tidy
make build
SERVICE_NAME=sample-app VERSION=0.1.0 DB_DSN="postgres://postgres:password@localhost:5432/postgres?sslmode=disable" ./bin/sample-app
```

Docker

```bash
make docker-build
docker run --rm -p 8080:8080 -e DB_DSN="postgres://postgres:password@localhost:5432/postgres?sslmode=disable" sample-app:local
```

Kubernetes (example)

```bash
kubectl apply -f k8s/configmap.yaml
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```

Notes
- The app expects a Postgres-compatible database configured via `DB_DSN`. For real clusters put credentials in a `Secret` and reference via env.
- The readiness probe calls `db.Ping()` with a small timeout; if the DB is unreachable the pod will show NotReady.
