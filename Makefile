HOST_ARCH := $(shell file `which docker` | awk '{print $$NF}')
ifeq ($(HOST_ARCH), arm64)
	DOCKER_PLATFORM := --platform linux/amd64
endif


.PHONY: build
build: ## Build service docker image.
	go build -mod vendor -o go-app ./main.go

.PHONY: image
image:
	DOCKER_BUILDKIT=1 docker build $(DOCKER_PLATFORM) -f Dockerfile . -t naive
