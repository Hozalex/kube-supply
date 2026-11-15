# KubeSupply development guidelines

## Language

- Use English for all repository content, including file and directory names, documentation, source code, comments, identifiers, configuration examples, tests, log messages, error messages, and UI text.

## Architecture

- Go + React/TypeScript, one repository, modular monolith.
- The MVP is read-only, without agents, a database, or a message broker. Connect directly to Kubernetes APIs.
- Multiple replicas collect observations independently. Do not assume shared memory between Pods.
- cmd/server handles dependency wiring and lifecycle; internal/api contains the HTTP layer.
- Add modules, abstractions, and dependencies when required by implemented behavior.
- Kubernetes/Helm integrations and application grouping must remain independent of HTTP handlers.
- Never commit credentials, log them, or return them through the API.
- Document architectural decisions in docs/adr. Mark unresolved product decisions as open.

## Validation

- For Go changes, run gofmt, go test ./..., and go vet ./... once the first package exists.
- Match validation to the change. Having no packages does not demonstrate implementation correctness.
