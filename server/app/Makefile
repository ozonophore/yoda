BINARY_NAME=app
build: clean
	go build -o ./build/${BINARY_NAME}-darwim main.go
	GOOS=linux GOARCH=amd64 go build -o ./build/${BINARY_NAME}-amd64-linux main.go
run:
	go run main.go -c config.app.yml
#go mod tidy
test:
	go test -v ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

clean:
	go clean
	rm -f ./build/${BINARY_NAME}-darwin
	rm -f ./build/${BINARY_NAME}-amd64-linux