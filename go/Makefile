branch_name := $(shell git rev-parse --abbrev-ref HEAD)
dynamos := sidecar policy-enforcer orchestrator agent api-gateway
dockerhub_account := javernus

prepare:
	go mod tidy
	go mod download

proto:
	protoc -I ../proto-files --go_out=./pkg/proto --go_opt=paths=source_relative --go-grpc_out=./pkg/proto --go-grpc_opt=paths=source_relative ../proto-files/*.proto

$(dynamos): prepare proto
	cp Dockerfile go.mod go.sum ./cmd/$@
	cp -r pkg ./cmd/$@
	(trap 'rm -f ./cmd/$@/Dockerfile; rm -f ./cmd/$@/go.*; rm -rf ./cmd/$@/pkg' EXIT; \
	docker build --platform linux/amd64 --build-arg NAME='$@' -t $(dockerhub_account)/$@:$(branch_name) -t $(dockerhub_account)/$@:latest ./cmd/$@/)
	docker push $(dockerhub_account)/$@:$(branch_name)
	docker push $(dockerhub_account)/$@:latest


$(sql_microservices): prepare proto
	cp Dockerfile go.mod go.sum ./cmd/$@
	cp -r pkg ./cmd/$@
	(trap 'rm -f ./cmd/$@/Dockerfile; rm -f ./cmd/$@/go.*; rm -rf ./cmd/$@/pkg' EXIT; \
	docker build --platform linux/amd64 --build-arg NAME='$@' -t $(dockerhub_account)/$@:$(branch_name) -t $(dockerhub_account)/$@:latest ./cmd/$@/)
	docker push $(dockerhub_account)/$@:$(branch_name)
	docker push $(dockerhub_account)/$@:latest

all: $(dynamos) $(sql_microservices)
dynamos: $(dynamos)

.PHONY: all prepare proto $(dynamos)
