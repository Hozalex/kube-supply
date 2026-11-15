# ADR 0001: Initial architecture

Date: 2026-11-13. Status: Accepted.

## Context

The portal must provide an overview of infrastructure applications and their versions across Kubernetes clusters. Argo CD is already in use. The portal is expected to support application installation and upgrades in the future.

## Decision

Use one repository with a Go backend and a React/TypeScript frontend, structured as a modular monolith. For the MVP, configuration is file-based and observations are held in memory. Do not introduce a database, Redis, or a message broker.

The backend connects directly to Kubernetes APIs. All clusters are reachable over the company network. No agents are required. Each cluster has a separate client and credentials; in-cluster authentication is an option for the local cluster.

Each read-only backend replica collects its own observations. The API reports observation timestamps and source errors. Snapshots may differ between replicas. Leader election and shared memory are not required. This approach does not automatically extend to future state-changing operations.

## Consequences

Development and deployment remain simple; snapshots are rebuilt after restarts. Source requests grow with the number of replicas, so concurrency limits and polling jitter are required. Historical observations are not persisted.

Components can become separate processes within the same repository. This requires defining how state is exchanged; moving code into a separate command is not enough. Storage and coordination for future operations will be covered by a separate ADR.
