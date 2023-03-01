all: .image

IMAGE=ghcr.io/duglin/keda-rabbit

.image: send receive Dockerfile
	docker build -t $(IMAGE) .
	docker push $(IMAGE)
	touch .image

keda:
	kubectl apply -f https://github.com/kedacore/keda/releases/download/v2.9.2/keda-2.9.2.yaml

rabbit:
	helm repo add bitnami https://charts.bitnami.com/bitnami
	helm install rabbitmq --set auth.username=user \
		--set auth.password=PASSWORD bitnami/rabbitmq --wait

consumer:
	kubectl apply -f consumer.yaml
	kubectl apply -f scaler.yaml

producer:
	@kubectl run -tiq keda-producer --image $(IMAGE) --restart=Never --rm=true \
		send amqp://user:PASSWORD@rabbitmq.default.svc.cluster.local:5672 300

watch:
	watch -n 1 -t ./stats

clean:
	-kubectl delete -f scaler.yaml
	-kubectl delete -f consumer.yaml
	-kubectl delete pod keda-producer

clean-all: clean
	-helm uninstall rabbitmq
	# -kubectl delete ns keda
