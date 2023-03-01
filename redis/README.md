# Redis consumer and sender

A simple docker container that will receive messages from a Redis queue and
scale via KEDA.  The receiver will receive a single message at a time (per
instance), and sleep for 1 second to simulate performing work.  When adding a
massive amount of queue messages, KEDA will drive the container to scale out
according to the event source (Redis).


## Running the demo

### Requirements:
- Docker (https://docs.docker.com/engine/install/ubuntu/)
- `kubectl` (https://kubernetes.io/docs/tasks/tools/)

### Setup
Use k3d so you can configure the HPA config:
```
$ make cluster
```

This will install `k3d`, if not already installed, and then create a k3d
cluster with the right config.

Either way, make sure `kubectl` points to your cluster.

Deploy KEDA and Redis:
```
$ make keda redis
```

You only need to do these once, not on each demo execution.

### Run the demo:
Ideally, in one window run:
```
$ make watch
```
To see the pods as they appear/disappear.

Then in a different window run the demo itself:
```
$ ./demo
```

During the demo:
- press space-bar to unpause (at each step)
- press 'f' to stop the slow-typing
- press 'r' to stop pausing at each step

You should be able to rerun the demo multiple times.

### Cleaning
- `make clean` to clean-up after the demo
- `make clean-cluster` to clean up the entire cluster
- `make clean-all` to stop k3d

