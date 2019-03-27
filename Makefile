IMAGE ?= $(shell basename `pwd`)
TAG ?= latest
REGISTRY ?= quay.io/joerx

build:
	docker build -t $(IMAGE):$(TAG) .

publish: build
	docker tag $(IMAGE):$(TAG) $(REGISTRY)/$(IMAGE):$(TAG)
	docker push $(REGISTRY)/$(IMAGE):$(TAG)
