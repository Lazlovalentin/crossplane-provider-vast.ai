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

Scaffold. Resource Kinds are configured (see `config/`), with generated
Kubernetes types, CRDs, examples, and controllers included in the repository.

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

## Install

Install the provider into a Crossplane-enabled cluster:

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-vastai
spec:
  package: ghcr.io/lazlovalentin/provider-vastai:v0.1.0
```

```console
kubectl apply -f examples/install.yaml
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

## Examples

More ready-to-apply manifests live under [`examples/`](./examples). The snippets
below cover the most common workflows.

### SSH key

Register a public key so you can `ssh` into rented instances.

```yaml
apiVersion: account.vastai.crossplane.io/v1alpha1
kind: SSHKey
metadata:
  name: ops-laptop
spec:
  forProvider:
    sshKey: "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAI... ops@laptop"
  providerConfigRef:
    name: default
```

### On-demand GPU instance

Rent an RTX 4090 offer, attach the SSH key above, install extras at boot, and
write connection details to a Secret.

```yaml
apiVersion: compute.vastai.crossplane.io/v1alpha1
kind: Instance
metadata:
  name: training-rtx4090
spec:
  forProvider:
    # Numeric offerId from `vastai search offers` (or vastai_gpu_offers).
    offerId: 1234567
    image: pytorch/pytorch:2.3.0-cuda12.1-cudnn8-runtime
    diskGb: 60
    label: training-rtx4090
    useSsh: true
    useJupyterLab: true
    onstart: |
      #!/bin/bash
      apt-get update && apt-get install -y htop
    env:
      MODEL_NAME: "llama-3-8b"
      EPOCHS: "3"
    sshKeyIds:
      - "12345"
  providerConfigRef:
    name: default
  writeConnectionSecretToRef:
    name: training-rtx4090-connection
    namespace: crossplane-system
```

### Interruptible (bid) instance

Same shape, but with a max bid price in `$/hour`.

```yaml
apiVersion: compute.vastai.crossplane.io/v1alpha1
kind: Instance
metadata:
  name: batch-spot
spec:
  forProvider:
    offerId: 1234567
    image: nvidia/cuda:12.4.0-runtime-ubuntu22.04
    diskGb: 40
    label: batch-spot
    bidPrice: 0.35
  providerConfigRef:
    name: default
```

### Persistent volume

Provision a volume from a `vastai search volumes` offer.

```yaml
apiVersion: storage.vastai.crossplane.io/v1alpha1
kind: Volume
metadata:
  name: dataset-cache
spec:
  forProvider:
    offerId: 9988
    size: 500
    name: dataset-cache
    disableCompression: false
    # cloneFromId: 1234   # optional: clone an existing volume
  providerConfigRef:
    name: default
```

### Network volume

Multi-host shared storage.

```yaml
apiVersion: storage.vastai.crossplane.io/v1alpha1
kind: NetworkVolume
metadata:
  name: shared-checkpoints
spec:
  forProvider:
    offerId: 7777
    size: 200
    name: shared-checkpoints
  providerConfigRef:
    name: default
```

### Cluster + overlay network

Group rented machines into a private overlay.

```yaml
apiVersion: compute.vastai.crossplane.io/v1alpha1
kind: Cluster
metadata:
  name: training-cluster
spec:
  forProvider:
    name: training-cluster
  providerConfigRef:
    name: default
---
apiVersion: network.vastai.crossplane.io/v1alpha1
kind: Overlay
metadata:
  name: training-overlay
spec:
  forProvider:
    name: training-overlay
    clusterIdRef:
      name: training-cluster
  providerConfigRef:
    name: default
---
apiVersion: network.vastai.crossplane.io/v1alpha1
kind: OverlayMember
metadata:
  name: training-overlay-node-1
spec:
  forProvider:
    overlayIdRef:
      name: training-overlay
    instanceIdRef:
      name: training-rtx4090
  providerConfigRef:
    name: default
```

### Worker group + endpoint (inference serving)

Expose a pool of workers behind a Vast.ai endpoint.

```yaml
apiVersion: compute.vastai.crossplane.io/v1alpha1
kind: WorkerGroup
metadata:
  name: vllm-workers
spec:
  forProvider:
    name: vllm-workers
    image: vllm/vllm-openai:latest
    diskGb: 80
    minWorkers: 1
    maxWorkers: 4
  providerConfigRef:
    name: default
---
apiVersion: compute.vastai.crossplane.io/v1alpha1
kind: Endpoint
metadata:
  name: vllm-endpoint
spec:
  forProvider:
    name: vllm-endpoint
    workerGroupIdRef:
      name: vllm-workers
  providerConfigRef:
    name: default
```

### Reusable instance template

Define an instance template once, then reference it from instances.

```yaml
apiVersion: account.vastai.crossplane.io/v1alpha1
kind: InstanceTemplate
metadata:
  name: pytorch-cuda12
spec:
  forProvider:
    name: pytorch-cuda12
    image: pytorch/pytorch:2.3.0-cuda12.1-cudnn8-runtime
    diskGb: 60
    useSsh: true
    onstart: |
      #!/bin/bash
      pip install -U transformers accelerate
  providerConfigRef:
    name: default
```

### Team, role, member

```yaml
apiVersion: account.vastai.crossplane.io/v1alpha1
kind: Team
metadata:
  name: ml-platform
spec:
  forProvider:
    name: ml-platform
  providerConfigRef:
    name: default
---
apiVersion: account.vastai.crossplane.io/v1alpha1
kind: TeamRole
metadata:
  name: ml-platform-admin
spec:
  forProvider:
    name: admin
    teamIdRef:
      name: ml-platform
    permissions:
      - read
      - write
      - manage_billing
  providerConfigRef:
    name: default
---
apiVersion: account.vastai.crossplane.io/v1alpha1
kind: TeamMember
metadata:
  name: ml-platform-alice
spec:
  forProvider:
    email: alice@example.com
    teamIdRef:
      name: ml-platform
    roleIdRef:
      name: ml-platform-admin
  providerConfigRef:
    name: default
```

### Environment variable

Set an account-wide env var injected into instances.

```yaml
apiVersion: account.vastai.crossplane.io/v1alpha1
kind: EnvironmentVariable
metadata:
  name: wandb-api-key
spec:
  forProvider:
    key: WANDB_API_KEY
    value: "your-wandb-key"
  providerConfigRef:
    name: default
```

### Sub-account + API key

```yaml
apiVersion: account.vastai.crossplane.io/v1alpha1
kind: Subaccount
metadata:
  name: ci-runner
spec:
  forProvider:
    email: ci@example.com
    name: ci-runner
  providerConfigRef:
    name: default
---
apiVersion: account.vastai.crossplane.io/v1alpha1
kind: APIKey
metadata:
  name: ci-runner-key
spec:
  forProvider:
    name: ci-runner-key
    subaccountIdRef:
      name: ci-runner
  providerConfigRef:
    name: default
  writeConnectionSecretToRef:
    name: ci-runner-key
    namespace: crossplane-system
```

Browse [`examples/cluster/`](./examples/cluster) and
[`examples/namespaced/`](./examples/namespaced) for the full set of curated
manifests (cluster-scoped vs namespaced MRs).

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/lazlovalentin/crossplane-provider-vast.ai/issues).
