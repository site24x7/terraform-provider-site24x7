.DEFAULT_GOAL := help

BINARY     := terraform-provider-site24x7
TEST_FLAGS ?= -race
PKGS       ?= $(shell go list ./... | grep -v /vendor/)

.PHONY: help
help:
	@grep -E '^[a-zA-Z-]+:.*?## .*$$' Makefile | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "[32m%-12s[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## build terraform-provider-site24x7
	go build \
		-gcflags "all=-N -l" \
		-o $(BINARY) \
		main.go

.PHONY: install
install: build ## install terraform-provider-site24x7
	mkdir -p $(HOME)/.terraform.d/plugins
	cp $(BINARY) $(HOME)/.terraform.d/plugins

.PHONY: test
test: ## run tests
	go test $(TEST_FLAGS) $(PKGS)

.PHONY: vet
vet: ## run go vet
	go vet $(PKGS)

.PHONY: coverage
coverage: ## generate code coverage
	go test $(TEST_FLAGS) -covermode=atomic -coverprofile=coverage.txt $(PKGS)
	go tool cover -func=coverage.txt

.PHONY: lint
lint: ## run golangci-lint
	golangci-lint run
