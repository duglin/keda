IMAGE=ghcr.io/duglin/keda-redis

.image: send receive Dockerfile
	docker build -t $(IMAGE) .
	docker push $(IMAGE)
	touch .image

k3d:
	k3d > /dev/null 2>&1 || \
	curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash

cluster: k3d
	k3d cluster list | grep redis || \
	k3d cluster create redis --k3s-arg '--kube-controller-manager-arg=horizontal-pod-autoscaler-sync-period=1s@server:*'

keda:
	kubectl get ns keda || \
	kubectl apply -f https://github.com/kedacore/keda/releases/download/v2.9.2/keda-2.9.2.yaml

redis:
	kubectl run redis --image redis --restart=Always
	kubectl expose pod redis --port 6379

consumer:
	kubectl apply -f consumer.yaml
	kubectl apply -f scaler.yaml

producer:
	@kubectl run -tiq redis-producer --image $(IMAGE) --restart=Never \
		--rm=true send redis.default.svc.cluster.local 300

watch:
	watch -n 1 -t ./stats

clean:
	-kubectl delete -f scaler.yaml
	-kubectl delete -f consumer.yaml
	-kubectl delete pod redis-producer

clean-cluster: clean
	-kubectl delete pod redis
	-kubectl delete service redis
	-kubectl delete ns keda

clean-all:
	-k3d cluster delete redis
