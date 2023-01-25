SHELL := /bin/bash
VERSION := 1.0
# ==============================================================================
all: load push

azure:
	docker build \
		-f docker/dockerfile.scholar-power-api.basic \
		-t thefueley/scholar-power-api:$(VERSION) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.
	docker push thefueley/scholar-power-api:$(VERSION)

load:
	docker buildx build --load \
		-f docker/dockerfile.scholar-power-api.basic \
		-t thefueley/scholar-power-api:$(VERSION) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.
push:
	docker buildx build --push \
		-f docker/dockerfile.scholar-power-api.basic \
		-t thefueley/scholar-power-api:$(VERSION) \
		--platform linux/amd64,linux/arm64 \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.

# ==============================================================================
run:
	go run cmd/server/main.go
# ==============================================================================
tidy:
	go mod tidy
	go mod vendor
# ==============================================================================
docker-down:
	docker rm -f $(shell docker ps -aq)

docker-clean:
	docker system prune -f	