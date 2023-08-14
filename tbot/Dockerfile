## Build
FROM golang:1.20 AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN make build

## Deploy
FROM centos:7

WORKDIR /app

VOLUME /tmp

COPY --from=build /app/build/app-amd64-linux /app/app-amd64-linux

ENTRYPOINT ["/app/app-amd64-linux"]