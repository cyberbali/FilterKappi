# First prompt for the coding agent (scaffold only)

Paste this after `AGENTS.md` is committed.

---

Read `AGENTS.md` first and follow it.

Create the project skeleton for FilterKappi, a Go web app.

**Requirements**
- Go 1.22+ standard `net/http` with `ServeMux` (no web framework)
- Layout: `cmd/filterkappi/`, `internal/{config,httpserver,logging,storage}`, `migrations/`, `test/`
- Config loaded from environment variables with validation. Fail to start on invalid config. Provide `.env.example` with no real secrets
- Structured logging with `log/slog` (JSON output), request ID middleware
- `/healthz` endpoint
- HTTP server with ReadHeaderTimeout, ReadTimeout, WriteTimeout, IdleTimeout, MaxHeaderBytes, and graceful shutdown on SIGINT/SIGTERM
- Request body size limit middleware
- Makefile targets: `build`, `test`, `lint`, `vuln`
- Multi-stage Dockerfile, final image non-root and minimal
- `docker-compose.yml` with the app and a database service (Postgres), DB not exposed to the host
- GitHub Actions workflow running `go vet`, `go test`, `staticcheck`, `govulncheck`, and `gosec`

**Must not**
- Implement any business feature, authentication, or database schema
- Add third-party dependencies without justification in the PR description
- Commit any secrets

**Acceptance criteria**
- `make build test lint vuln` passes
- `docker compose up` starts the app and `/healthz` returns 200
- Container runs as non-root
- One PR, with the PR template filled in
