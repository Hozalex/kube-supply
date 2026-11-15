# KubeSupply

A portal for Kubernetes infrastructure applications. Each application page shows its installations, versions, and health across clusters.

## Current status

A minimal Go HTTP server with Gin routing is implemented. Cluster connections, discovery, and the frontend are not implemented yet.

## Requirements

- Go 1.25 or later.
- Port 8080 available locally.

## Run locally

From the repository root:

```sh
go run ./cmd/server
```

The server listens on `:8080`. If the port is already in use, startup fails with a nonzero exit code.

Gin uses debug mode by default. To disable its debug output:

```sh
GIN_MODE=release go run ./cmd/server
```

## Check the HTTP API

While the server is running, use another terminal:

```sh
curl -i http://localhost:8080/healthz
```

Expected response: HTTP `200` with `Content-Type: application/json; charset=utf-8` and body:

```json
{"status":"ok"}
```

This endpoint reports process liveness; it does not check cluster connectivity.

An unknown path returns HTTP `404`:

```sh
curl -i http://localhost:8080/missing
```

An unsupported method on `/healthz` returns HTTP `405`:

```sh
curl -i -X POST http://localhost:8080/healthz
```

## Development checks

```sh
gofmt -l cmd/server internal/api
go test ./...
go vet ./...
```

`gofmt -l` should produce no output. There are currently no automated tests; `go test ./...` checks package compilation but does not validate HTTP behavior.

## Architecture decisions

- One repository with a Go backend and a React + TypeScript frontend.
- A modular monolith, initially running as one process in one container.
- Direct Kubernetes API connections without agents; clusters are reachable over the company network.
- A read-only MVP with file-based configuration and in-memory observations; no database, Redis, or message broker.
- Multiple backend replicas collect observations independently; snapshot timestamps may differ.
- Application installation and upgrades are planned for later releases.

## Repository layout

```text
cmd/server/       Go entry point, dependency wiring, and lifecycle
internal/api/     HTTP API and handlers
web/              Future React frontend
config/           Configuration documentation and future example
api/              Future OpenAPI contract
docs/adr/         Architecture decision records
```

The clusters, discovery, applications, sources, and versions packages will be added as their behavior is implemented. Separate collector and worker commands are not created in advance.

See [architecture](docs/architecture.md) and the [initial architecture decision](docs/adr/0001-initial-architecture.md).
