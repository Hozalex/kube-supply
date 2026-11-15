# KubeSupply architecture

## MVP scope

Cluster connections, discovery, and application views with installations across environments. Installation, upgrades, and adopting existing installations are outside the initial scope.

## Data flow

Configuration -> cluster connections -> discovery adapters -> normalization and grouping -> in-memory snapshot -> HTTP API -> React UI.

UI requests do not trigger full discovery runs. An unavailable source does not imply that its installations were deleted; the last known state is marked as stale.

## Planned modules

- clusters: cluster connections, identity, and availability.
- discovery: Kubernetes, Helm, and Argo CD adapters; normalized observations.
- applications: logical applications and installation mappings.
- sources: Helm repositories, OCI, and Git.
- versions: version identification and comparison.
- operations: future state-changing operations.

Packages are added as behavior is implemented. Separate services or interfaces are not created for each module in advance.

## Model and constraints

A logical application groups multiple installations. An installation has a cluster, namespace, name, source, and observed state. Chart versions, appVersion, and image versions are distinct data. Unknown versions must remain explicitly unknown.

Argo CD typically renders Helm charts without creating Helm releases, so Helm discovery alone is insufficient. Chart sources may require manual mapping. A missing Git reference does not imply default values.

## Open decisions

- Update policy: the latest minor within the current major, or the latest patch within the current minor. The UI must show whether an installation is up to date and identify the available version.
- Credential renewal mechanism.
- Configuration schema and normalized installation model.
- Storage and coordination for future installation operations.

## Implementation sequence

1. HTTP server scaffold.
2. Server lifecycle and configuration.
3. A single cluster connection and access diagnostics.
4. An initial discovery pass and normalized observations.
5. Application catalog and detail view, followed by additional sources and multiple clusters.
