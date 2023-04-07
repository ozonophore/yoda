## Build
FROM golang:1.20 AS build

WORKDIR /app

COPY ./go.work ./
COPY ./go.work.sum ./
COPY ./common ./common
COPY ./webapp ./webapp
COPY ./app ./app
RUN go mod download

WORKDIR /app/webapp
RUN make build

## Deploy
FROM centos:7

WORKDIR /app

COPY --from=build /app/webapp/build/webapp-amd64-linux /

EXPOSE 8080

ENTRYPOINT ["/webapp-amd64-linux"]
