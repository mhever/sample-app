# Incident: Node DiskPressure caused pod eviction and scheduling halt

## Environment:
Single-node k3s cluster (home lab), 2 TB disk, containerd runtime, with a default settings (10-15% disk free) for Disk Pressure.

## Impact
- pods evicted, or stuck in ContainerStatusUnkown
- new pods not scheduled due to NoSchedule taint

## Detection
- pods in Evicted / ContainerStatusUnkown / Error state
- node with active disk-pressure taint

~~~
sudo k3s kubectl get pods
NAME                          READY   STATUS                   RESTARTS      AGE
postgres-654766c7bb-7v7m7     0/1     ContainerStatusUnknown   1 (21h ago)   34h
postgres-654766c7bb-8sjvj     0/1     ContainerStatusUnknown   0             26m
postgres-654766c7bb-fxkmd     1/1     Running                  0             5m
sample-app-64f57c888c-4kwjj   0/1     ContainerStatusUnknown   0             26m
sample-app-64f57c888c-5hh88   1/1     Running                  0             5m
sample-app-64f57c888c-7vzjx   1/1     Running                  0             5m
sample-app-64f57c888c-99v6s   0/1     Error                    3             33h
sample-app-64f57c888c-j98fs   0/1     ContainerStatusUnknown   0             26m
sample-app-64f57c888c-lbds7   0/1     Error                    0             11h
~~~

~~~
sudo k3s kubectl get node thinkcentre -o jsonpath='{range .spec.taints[*]}{.key}={.value}:{.effect}{"\n"}{end}'
node.kubernetes.io/disk-pressure=:NoSchedule
~~~

check the system with
* sudo k3s kubectl describe node thinkcentre
* sudo k3s kubectl get node thinkcentre

## Root cause
- ephemeral storage usage exceeded kubelet eviction thresholds, triggering DiskPressure=True
- kubelet tainted the node, evicted pods

## Steps to resolve the issue
* fix the underlying issue - free up disk space
* restart the k3s service with "sudo systemctl restart k3s" to force kubelet to reevaluate node conditions
* Disk Pressure condition was cleared automatically, taint removed by kubelet
* "sudo k3s kubectl delete pod --field-selector=status.phase=Failed" to remove the failed pods