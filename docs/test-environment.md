# remediator-test Namespace

## Purpose

`remediator-test` is an isolated sandbox namespace that mirrors the `home` ("prod") namespace. It exists to support the [gitops-remediator](https://github.com/mhever/gitops-remediator) project — a Go service that watches a namespace for failure events, diagnoses them with an LLM, and opens GitHub PRs with proposed fixes.

The namespace is safe to break deliberately. Failures are triggered by committing intentionally broken manifests to `k8s/overlays/test/`, FluxCD reconciles them, and the remediator reacts. Live services in `home` are never touched.

---

## What Was Added to This Repo

**`k8s/overlays/test/`** — Kustomize overlay targeting `remediator-test`. Mirrors the `home` overlay but with 1 replica (no need for 3 in a test environment) and no Prometheus annotations patch. Image tag is pinned the same way as `home` and updated here when needed.

**`clusters/home/apps-test.yaml`** — FluxCD `Kustomization` resource that points Flux at `./k8s/overlays/test`. Flux automatically picks up all `.yaml` files in `./clusters/home`, so no further registration is needed.

---

## Secrets

Kubernetes secrets are **namespace-scoped**. The secrets from `home` do not exist in `remediator-test` and must be created manually.

For a test namespace, use simple known values — you'll be injecting intentional failures and don't want to debug auth issues on top of them. Both secrets must use the **same password** because `sample-app-secret` is what the app uses to connect to Postgres, and `postgres-secret` is what Postgres uses to set its own password.

```bash
kubectl create secret generic postgres-secret \
  --from-literal=postgres-password='XXX' \
  --namespace=remediator-test

kubectl create secret generic sample-app-secret \
  --from-literal=DB_PASSWORD='XXX' \
  --namespace=remediator-test
```

If secrets are missing, pods fail with `CreateContainerConfigError`. Check with:

```bash
kubectl get secret -n remediator-test
```

---

## Recreating from Scratch

```bash
# 1. Create the namespace
kubectl create namespace remediator-test

# 2. Create secrets (same password for both — see above)
kubectl create secret generic postgres-secret \
  --from-literal=postgres-password='testpassword' \
  --namespace=remediator-test

kubectl create secret generic sample-app-secret \
  --from-literal=DB_PASSWORD='testpassword' \
  --namespace=remediator-test

# 3. The overlay and FluxCD Kustomization are already in the repo.
#    Flux will deploy the app once it reconciles (up to 1 minute).
```

---

## Triggering Failures

Edit `k8s/overlays/test/kustomization.yaml` or the base manifests and commit to `main`. FluxCD reconciles within 1 minute.

| Failure | How to trigger |
|---|---|
| OOMKilled | Add a patch setting `resources.limits.memory: 1Mi` on the app deployment |
| CrashLoopBackOff | Set a required env var to an invalid value |
| ImagePullBackOff | Change the image tag to `does-not-exist` in `kustomization.yaml` |

Revert the commit to restore the namespace to a healthy state.

---

## Verification

```bash
# FluxCD reconciliation status
kubectl get kustomization -n flux-system

# Pods in the namespace
kubectl get pods -n remediator-test

# Events (useful for diagnosing failures)
kubectl get events -n remediator-test --sort-by='.lastTimestamp'
```

With homelab-mcp in Claude Desktop, use `k8s_pods` and `k8s_events` filtered to `remediator-test` for live monitoring during remediator development.
