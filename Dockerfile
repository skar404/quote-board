FROM golang:1.12-alpine AS build_base

WORKDIR /app

RUN apk add bash ca-certificates git gcc g++ libc-dev

COPY go.mod .
COPY go.sum .

RUN go mod download

FROM build_base AS server_builder

COPY . /app

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go install -a -tags netgo -ldflags '-w -extldflags "-static"' ./core

FROM alpine:latest AS base_app

COPY --from=server_builder /go/bin/core /bin/core

EXPOSE 8080

ENTRYPOINT ["/bin/core"]
