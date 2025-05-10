GOBIN=$(PWD)/bin

generate:
	bin/buf generate
.PHONY: generate

format:
	bin/buf format -w
.PHONY: format

update:
	bin/buf dep update
.PHONY: update

install_tools:
	GOBIN=$(GOBIN) go install github.com/bufbuild/buf/cmd/buf@latest
	GOBIN=$(GOBIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(GOBIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(GOBIN) go install github.com/envoyproxy/protoc-gen-validate@latest
.PHONY: install_tools

build_fiber_tool:
	cd .. && GOBIN=$(GOBIN) go install
.PHONY: build_fiber_tool
