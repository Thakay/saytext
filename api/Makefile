build:
	@echo $(shell go env GOPATH)
	protoc -I . say.proto --plugin=protoc-gen-go=/home/kay/go/bin/protoc-gen-go --go_out=. --plugin=protoc-gen-go-grpc=/home/kay/go/bin/protoc-gen-go-grpc --go-grpc_out=.
