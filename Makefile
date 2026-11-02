BIN_DIR = bin
PACKAGE=github.com/atselvan/go-grpc
SERVER_DIR = server
CLIENT_DIR = client

PROJECTS = greet

.DEFAULT_GOAL := help

define build-proto
	protoc -I$(1)/proto --go_out=. --go_opt=module=${PACKAGE} --go-grpc_out=. --go-grpc_opt=module=${PACKAGE} $(1)/proto/*.proto
	go build -o ${BIN_DIR}/$(1)/${CLIENT_DIR} ./$(1)/${CLIENT_DIR}
	go build -o ${BIN_DIR}/$(1)/${SERVER_DIR} ./$(1)/${SERVER_DIR}
endef

.PHONY: deps all $(PROJECTS)

deps:
	go get -u -t ./...
	go mod tidy

all: $(PROJECTS) ## Generate Pbs and build for all projects

greet: ## Generate Pbs and build for greet
	$(call build-proto,greet)

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?\#\# .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?\#\# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
