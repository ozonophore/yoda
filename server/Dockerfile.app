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
FROM centos:7

WORKDIR /app

COPY --from=build /app/app/build/app-amd64-linux /

ENTRYPOINT ["/webapp-amd64-linux"]
