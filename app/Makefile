BINARY_NAME=app
build:
	go build -o ./bin/${BINARY_NAME} main.go

generate:
	oapi-codegen --package=api -generate=types,client -o pkg/api/api.go ../openapi/api.wb.yml
#go mod tidy
test:
	go test -v ./... -coverprofile=coverage.out