GOBIN=$(PWD)/bin

IMAGE:=xakepp35/xprotoc-gen:latest


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

copy_lib_from_docker:
	docker create --name temp-xprotoc-gen $(IMAGE)
	rm -rf lib
	docker cp temp-xprotoc-gen:/app/lib lib
	docker rm temp-xprotoc-gen
.PHONY: copy_lib_from_docker