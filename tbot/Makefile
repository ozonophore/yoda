BINARY_NAME=app
build: clean
	go build -o ./build/${BINARY_NAME}-darwim cmd/main.go
	GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o ./build/${BINARY_NAME}-amd64-linux cmd/main.go
run:
	go run cmd/main.go -c config.yml
#go mod tidy
test:
	go test -v ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
clean:
	go clean
	rm -f ./build/${BINARY_NAME}-darwin
	rm -f ./build/${BINARY_NAME}-amd64-linux