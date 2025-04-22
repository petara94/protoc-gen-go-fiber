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

install_tools: export GOBIN=$(shell pwd)/$(OUT)
install_tools:
	go get -tool github.com/bufbuild/buf/cmd/buf@v1.53.0
	go get -tool github.com/envoyproxy/protoc-gen-validate@v1.2.1
	go get -tool github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.1.1
.PHONY: install_tools

build_gen:
	go build -o $(OUT)/gateaway-fiber tools/gateaway-fiber/main.go
.PHONY: build_gen
