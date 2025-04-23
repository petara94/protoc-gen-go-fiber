PROTO_SRC=api
OUT_PB=gen/pb
OUT_GW=gen/gw

OUT=bin

generate:
	go tool buf generate
.PHONY: generate

update:
	go tool buf dep update
.PHONY: update

install_tools:
	go get -tool github.com/bufbuild/buf/cmd/buf@v1.53.0 \
 				github.com/envoyproxy/protoc-gen-validate@v1.2.1 \
	 			github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.1.1 \
				google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1 \
				github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.26.3

	go install tool
.PHONY: buildinstall_tools

build:
	go build -o $(OUT)/protoc-gen-fiber main.go
.PHONY: build
