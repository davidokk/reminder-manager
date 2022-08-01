.PHONY: run
run:
	go run cmd/bot/main.go cmd/bot/server.go

# build app
.PHONY: build
build:
	go mod download \
    && CGO_ENABLED=0 go build -o ./bin/bot-main$(shell go env GOEXE) ./cmd/bot/main.go ./cmd/bot/server.go

LOCAL_BIN:=$(CURDIR)/bin
.PHONY: .deps
.deps:
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
