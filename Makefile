.PHONY: run-data
run-data:
	go run data-service/cmd/main/main.go

.PHONY: run-interface
run-interface:
	go run interface-service/cmd/main/main.go

.PHONY: run-bot
run-bot:
	go run bot/cmd/main/main.go

# build app
.PHONY: build
build:
	go mod download \
    && CGO_ENABLED=0 go build -o ./bin/data-service$(shell go env GOEXE) ./data-service/cmd/main/main.go \
    && CGO_ENABLED=0 go build -o ./bin/interface-service$(shell go env GOEXE) ./interface-service/cmd/main/main.go \
    && CGO_ENABLED=0 go build -o ./bin/bot$(shell go env GOEXE) ./bot/cmd/main/main.go

MIGRATIONS_DIR=./migrations
.PHONY: migration
migration:
	goose -dir=${MIGRATIONS_DIR} create $(NAME) sql

LOCAL_BIN:=$(CURDIR)/bin
.PHONY: .deps
.deps:
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
