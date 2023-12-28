## Build
FROM golang:1.21 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify
COPY . .

RUN make build

## Deploy
FROM centos:7

WORKDIR /app

COPY --from=build /app/openapi /app/openapi
COPY --from=build /app/public /app/public

VOLUME /tmp

EXPOSE 8080

COPY --from=build /app/build/web-amd64-linux /app/web-amd64-linux

ENTRYPOINT ["/app/web-amd64-linux"]