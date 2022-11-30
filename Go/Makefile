all: dependencies proto up

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/ConfigService.proto

dependencies:
	go mod download

up:
	go run cmd/config_service.go
