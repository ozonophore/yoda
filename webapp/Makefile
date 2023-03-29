generate:
	oapi-codegen --package=api -generate=types,gorilla -o pkg/api/server_api.go ../openapi/server.yml
generateProto:
	protoc --go_out=pkg ../grpc/client.proto