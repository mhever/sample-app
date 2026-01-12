Images initally were published to GHCR with tags sha-<8-char git commit SHA>.
Kustomize overlays must reference this exact tag format - by default it was showing a 7-char SHA value.

Now images are tagged with both, so referencing any will work in the kustomization overlays.