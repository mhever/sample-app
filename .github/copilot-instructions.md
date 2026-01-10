## Copilot / AI Agent Instructions for this repository

Summary
- Repository contains a minimal Go HTTP API using Gin, with Prometheus metrics, designed for Kubernetes deployment.
- Uses PostgreSQL-compatible database for readiness checks.
- Built with Go 1.25, includes Dockerfile, K8s manifests, and Makefile.

Primary goal for an agent
- Maintain and extend this simple microservice: add features, fix bugs, improve observability, or enhance deployment configs.
- Follow Go best practices: proper error handling, logging, and testing.

How to get started (short checklist)
- Inspect the code: `main.go` (entry point), `handlers.go` (endpoints), `config.go` (env vars), `db.go` (DB connection), `metrics.go` (Prometheus).
- Build locally: `make build` or `go build -o bin/sample-app ./`
- Run locally: `make run` (needs Postgres running or adjust DB_DSN)
- Test endpoints: curl http://localhost:8080/healthz, /, /readyz, etc.
- Deploy to K8s: apply manifests in `k8s/` (needs Postgres service)

Scaffold examples (if extending)
- Add new endpoint: update `handlers.go` and register in `main.go`
- Add config: update `Config` struct and `loadConfig()`
- Add metric: use `reqCounter` or `reqLatency` in `metrics.go`

Repository-specific conventions (discoverable now)
- Default branch: `main`
- Go version: 1.25 (update in `go.mod` and `Dockerfile`)
- Dependencies: Gin for HTTP, pgx for Postgres, Prometheus client
- Env vars: SERVICE_NAME, VERSION, DB_DSN
- Build: `make build` → `bin/sample-app`
- Docker: `make docker-build` → `sample-app:local`
- K8s: Deployment with liveness/readiness probes, ConfigMap for env

When modifying or adding files
- Keep minimal: focus on core functionality
- Update README if adding new endpoints or configs
- Test builds: run `make build` after changes
- For K8s changes: ensure probes and resources match app behavior

Files to check if added later (key touchpoints)
- `main.go` — router setup and server start
- `handlers.go` — all HTTP endpoint logic
- `config.go` — environment variable loading
- `db.go` — database connection and health checks
- `metrics.go` — Prometheus setup and middleware
- `k8s/deployment.yaml` — pod spec, probes, env
- `k8s/configmap.yaml` — app configuration
- `Dockerfile` — build and runtime image
- `Makefile` — build shortcuts
- `README.md` — usage and deployment instructions

If you need to implement functionality
- For new features: add to `handlers.go`, register route in `main.go`
- For DB queries: add functions in `db.go`, call from handlers
- For metrics: use existing counters/histograms or add new ones
- For config: add to `Config` struct, load in `loadConfig()`

When asking the user for clarification
- What specific endpoint or feature to add?
- Any particular DB schema or queries needed?
- Deployment environment details (K8s version, ingress, etc.)?

Errors and testing
- Build errors: check imports and Go version compatibility
- Runtime errors: check env vars (DB_DSN especially) and DB connectivity
- Test manually: curl endpoints, check logs
- No automated tests yet; add if needed in future

Questions to ask the user (first round)
1. What new feature or endpoint should I add?
2. Any specific DB schema or queries to implement?
3. Deployment tweaks needed (e.g., add ingress, secrets)?

If something here is unclear, ask me which part of the app to focus on next.
