BINARY_NAME=webapp
build: clean
	go build -o ./build/${BINARY_NAME}-darwim main.go
	GOOS=linux GOARCH=amd64 go build -o ./build/${BINARY_NAME}-amd64-linux main.go
run:
	go run main.go
generate:
	oapi-codegen --package=api -generate=types,gorilla -o pkg/api/server_api.go ../openapi/server.yml
generateProto:
	protoc --go_out=pkg ../grpc/client.proto
clean:
	go clean
	rm -f ./build/${BINARY_NAME}-darwin
	rm -f ./build/${BINARY_NAME}-amd64-linux