FROM golang:1.12-alpine AS build_base

WORKDIR /app

RUN apk add bash ca-certificates git gcc g++ libc-dev

COPY go.mod .
COPY go.sum .

RUN go mod download

FROM build_base AS server_builder

COPY . /app

RUN go build -o slink-core -ldflags '-s -w' ./core

FROM alpine:latest AS base_app

COPY --from=server_builder /app/slink-core /bin/slink-core

EXPOSE 8080

ENTRYPOINT ["/bin/slink-core"]
