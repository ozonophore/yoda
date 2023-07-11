BINARY_NAME=web
build: clean
	go build -o ./build/${BINARY_NAME}-darwim cmd/main.go
	GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o ./build/${BINARY_NAME}-amd64-linux cmd/main.go
run:
	go run cmd/main.go -c config.yml
generate:
	oapi-codegen --package=api -generate=types,server -o internal/api/api.go ./openapi/openapi.yml
clean:
	go clean
	rm -f ./build/${BINARY_NAME}-darwin
	rm -f ./build/${BINARY_NAME}-amd64-linux