SHELL:=/bin/bash

.PHONY: help
help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

## Nodes bootstrap commands
.node-1: ## Run node 1
	[ -f ./.build/store ] || make -s setup
	./.build/store -node-id=1 -http-port=8085 -replication-addr="127.0.0.1:12000"

.node-2: ## Run node 2
	[ -f ./.build/store ] || make -s setup
	./.build/store -node-id=2 -http-port=8086 -replication-addr="127.0.0.1:12001" -join-addr="localhost:8085"

.node-3: ## Run node 3
	[ -f ./.build/store ] || make -s setup
	./.build/store -node-id=3 -http-port=8087 -replication-addr="127.0.0.1:12002" -join-addr="localhost:8085"

## Cluster bootstrap commands
.PHONY: setup
setup: ## Setup distributed key-value store
	go mod tidy
	[ -d ./.build ] || mkdir .build
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o .build/store cmd/store/main.go

.PHONY: single-node
single-node: ## Run a single node cluster instance
	[ -f ./.build/store ] || make -s setup
	./.build/store -node-id=1 -http-port=8085 -replication-addr="localhost:12000"

.PHONY: master-node
master-node: ## Run a singles master node cluster instance
	@make -s .node-1

.PHONY: slave-node-1
slave-node-1: ## Attach an slave-2 node to the master node
	@make -s .node-2

.PHONY: slave-node-2
slave-node-2: ## Attach an slave-3 node to the master node
	@make -s .node-3

## Acceptance tests commands
.PHONY: run-acceptance-tests
run-acceptance-tests: ## Run acceptance tests over the store
	go test ./tests/store -v

## Benchmark commands
.PHONY: run-benchmark
run-benchmark: ## Run a benchmark tests over the store
	go test ./tests/store -bench=. -count=5 -run=^# -v