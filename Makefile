.PHONY: go-gen pb-file test-go-all
GO := $(shell which go)

go-gen: 
	@$(GO) generate ./...

test-go-all: 
	@$(GO) test ./...

pb-file:
	@protoc repository/grpc/filepb/file.proto --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=.	

lint: 
	@golint ./...