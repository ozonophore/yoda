## Build
FROM golang:1.20 AS build

WORKDIR /app

COPY server/go.work ./
COPY server/go.work.sum ./
COPY server/common ./common
COPY server/webapp ./webapp
COPY server/app ./app

RUN go mod download

WORKDIR /app/app
RUN make build

## Deploy
FROM ubuntu:latest

RUN apt-get update -y
RUN apt-get install -y tzdata

WORKDIR /app

COPY --from=build /app/app/build/app-amd64-linux /app/app-amd64-linux

ENTRYPOINT ["/app/app-amd64-linux"]
