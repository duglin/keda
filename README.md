# Some KEDA stuff

In most cases there should be a `demo` script in each one to run a demo.

## Running in K3D

To make HPA check for metrics faster you should set the
`horizontal-pod-autoscaler-sync-period` flag to something smaller than the
default (`15s`). It makes for a much better demo.

To do this in k3d you can do something like:
```
k3d cluster create rabbit --k3s-arg '--kube-controller-manager-arg=horizontal-pod-autoscaler-sync-period=1s@server:*'
```

Not all managed K8s providers give you access to this flag.

