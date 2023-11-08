FROM node as build-ui

WORKDIR /app

RUN npm i tar
RUN npm install -g create-react-app
RUN npm install -g openapi

COPY website/src /app/src
COPY website/package.json /app
COPY website/package-lock.json /app
COPY website/public /app/public
COPY website/tsconfig.json /app

COPY web/openapi/openapi.yml /app/openapi.yml
RUN npm run build

## Build
FROM golang:1.20 AS build-backend

WORKDIR /app

COPY web/cmd ./cmd
COPY web/internal ./internal
COPY web/openapi ./openapi
COPY web/go.mod ./go.mod
COPY web/go.sum ./go.sum
COPY web/Makefile ./Makefile

RUN go mod download

RUN make build

## Deploy
FROM centos:7

WORKDIR /app

COPY web/openapi /app/openapi
COPY --from=build-ui /app/build /app/public

VOLUME /tmp

EXPOSE 8080

COPY --from=build-backend /app/build/web-amd64-linux /app/web-amd64-linux

ENTRYPOINT ["/app/web-amd64-linux"]