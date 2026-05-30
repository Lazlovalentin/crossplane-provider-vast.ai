# Provider Vast.ai

`provider-vastai` is a [Crossplane](https://crossplane.io/) provider built with
[Upjet](https://github.com/crossplane/upjet). It wraps the
[realnedsanders/vastai](https://registry.terraform.io/providers/realnedsanders/vastai/latest)
Terraform provider to expose Vast.ai GPU rental, networking, storage, and team
resources as Kubernetes Custom Resources.

Based on:
- Crossplane upjet provider template
- Terraform provider: https://github.com/realnedsanders/terraform-provider-vastai

## Status

Scaffold. Resource Kinds are configured (see `config/`), but the generated
Kubernetes types (`apis/<group>/v1alpha1/...`) and controllers
(`internal/controller/...`) must be produced from the Terraform provider
schema using `make generate`.

## Resources

| Group     | Kind                 | Terraform resource             |
|-----------|----------------------|--------------------------------|
| account   | APIKey               | `vastai_api_key`               |
| account   | SSHKey               | `vastai_ssh_key`               |
| account   | Subaccount           | `vastai_subaccount`            |
| account   | Team                 | `vastai_team`                  |
| account   | TeamRole             | `vastai_team_role`             |
| account   | TeamMember           | `vastai_team_member`           |
| account   | EnvironmentVariable  | `vastai_environment_variable`  |
| account   | InstanceTemplate     | `vastai_template`              |
| compute   | Instance             | `vastai_instance`              |
| compute   | Cluster              | `vastai_cluster`               |
| compute   | ClusterMember        | `vastai_cluster_member`        |
| compute   | WorkerGroup          | `vastai_worker_group`          |
| compute   | Endpoint             | `vastai_endpoint`              |
| network   | Overlay              | `vastai_overlay`               |
| network   | OverlayMember        | `vastai_overlay_member`        |
| storage   | Volume               | `vastai_volume`                |
| storage   | NetworkVolume        | `vastai_network_volume`        |

## Developing

### Prerequisites

- Go 1.24+
- Terraform 1.5.x (1.6+ uses BSL and is not permitted)
- Docker (only for `make build` / image push)

### Generate types and controllers

Pulls the Terraform provider schema and runs the upjet generator:

```console
make generate
```

This populates `apis/<group>/v1alpha1/` and `internal/controller/<group>/`.

### Build

```console
go build ./...
```

### Run against a Kubernetes cluster

```console
make run
```

## ProviderConfig

Credentials are read from a Kubernetes Secret containing JSON with the
Vast.ai API key (and optional API URL override):

```json
{
  "api_key": "YOUR_VASTAI_API_KEY",
  "api_url": "https://console.vast.ai"
}
```

Example:

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: vastai-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {"api_key":"YOUR_VASTAI_API_KEY"}
---
apiVersion: vastai.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      namespace: crossplane-system
      name: vastai-creds
      key: credentials
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://gitlab.adva-soft.com/devopsteam/crossplane-vast.ai-provider/-/issues).
