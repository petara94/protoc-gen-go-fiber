#PATH:=$(PATH)$(CURDIR)/../bin

generate:
	go tool buf generate
.PHONY: generate

format:
	go tool buf format -w
.PHONY: format

update:
	go tool buf dep update
.PHONY: update

install_tools:
	go get -tool github.com/bufbuild/buf/cmd/buf@latest \
				google.golang.org/protobuf/cmd/protoc-gen-go@latest \
				google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
				github.com/envoyproxy/protoc-gen-validate@latest
	go install tool
.PHONY: install_tools

build_fiber_tool:
	cd .. && go build -o bin/protoc-gen-go-fiber
.PHONY: build_fiber_tool
