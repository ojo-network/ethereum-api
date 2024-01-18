BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
BUILD_DIR ?= $(CURDIR)/build
COMMIT    := $(shell git log -1 --format='%H')

###############################################################################
##                                  Version                                  ##
###############################################################################

ifeq (,$(VERSION))
  VERSION := $(shell git describe --exact-match 2>/dev/null)
  # if VERSION is empty, then populate it with branch's name and raw commit hash
  ifeq (,$(VERSION))
    VERSION := $(BRANCH)-$(COMMIT)
  endif
endif

###############################################################################
##                              Build / Install                              ##
###############################################################################

ldflags = -X github.com/ojo-network/ethereum-api/cmd.Version=$(VERSION) \
		  -X github.com/ojo-network/ethereum-api/cmd.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)'

build: go.sum
	@echo "--> Building..."
	CGO_ENABLED=0 go build -mod=readonly -o $(BUILD_DIR)/ $(BUILD_FLAGS) ./...

install: go.sum
	@echo "--> Installing..."
	CGO_ENABLED=0 go install -mod=readonly $(BUILD_FLAGS) ./...

.PHONY: build install

###############################################################################
##                                  Docker                                   ##
###############################################################################

ethereum_VERSION=20.5.0

docker-build-api:
	@docker build -t ethereum-api .

docker-push-api:
	@docker tag ethereum-api us-west4-docker.pkg.dev/ojo-network/ethereum-api/ethereum-api
	@docker push us-west4-docker.pkg.dev/ojo-network/ethereum-api/ethereum-api

docker-build-ethereumd-e2e:
	@git clone https://github.com/ethereum-labs/ethereum.git
	@cd ethereum; git checkout v${ethereum_VERSION}
	@DOCKER_BUILDKIT=1 docker build -t ghcr.io/ojo-network/ethereumd-e2e -f tests/e2e/docker/ethereum.Dockerfile .
	@rm -rf ethereum

docker-push-ethereumd-e2e:
	@docker push ghcr.io/ojo-network/ethereumd-e2e

.PHONY: docker-build-api docker-build-ethereumd-e2e docker-push-ethereumd-e2e

###############################################################################
##                              Tests & Linting                              ##
###############################################################################

PACKAGES_UNIT=$(shell go list ./... | grep -v -e '/tests/e2e')
PACKAGES_E2E=$(shell go list ./... | grep '/e2e')

test-unit:
	@go test -mod=readonly -race $(PACKAGES_UNIT) -v

test-e2e:
	$(MAKE) docker-build-api
	@go test -mod=readonly -race $(PACKAGES_E2E) -v

.PHONY: test-unit test-e2e

lint:
	@echo "--> Running linter"
	@go run github.com/golangci/golangci-lint/cmd/golangci-lint run --timeout=10m

.PHONY: lint

###############################################################################
##                                  All                                      ##
###############################################################################

all: test-unit install

.PHONY: all
