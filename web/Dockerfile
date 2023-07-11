## Build
FROM golang AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN make build

## Deploy
FROM centos:7

WORKDIR /app

COPY --from=build /app/openapi /app/openapi

VOLUME /tmp

EXPOSE 8080

COPY --from=build /app/build/web-amd64-linux /app/web-amd64-linux

ENTRYPOINT ["/app/web-amd64-linux"]