BINARY_NAME=app
build:
	go build -o ./bin/${BINARY_NAME} main.go

generate:
	oapi-codegen --package=wbclient -generate=types,client -o pkg/wbclient/api.go ../openapi/api.wb.yml
#go mod tidy