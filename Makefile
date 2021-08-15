run-server:
	go run server/main.go

run-client:
	go run client/main.go

#protoc:
#	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld/helloworld.proto

protoc:
	protoc --go_out=./pkg --go_opt=paths=source_relative --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative proto/*.proto

tidy:
	go mod tidy

vendor:
	go mod vendor