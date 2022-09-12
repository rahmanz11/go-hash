ARG GO_VERSION=1.18

FROM golang:${GO_VERSION}-alpine AS builder

RUN mkdir -p /app/build
RUN mkdir -p /app/go-hash

WORKDIR /app

COPY . go-hash/

WORKDIR /app/go-hash
RUN go mod download
RUN go build -o /app/build/ main.go

FROM alpine:latest

RUN mkdir -p /app/build

WORKDIR /app

COPY --from=builder /app/build/* build/

WORKDIR /app/build

EXPOSE 8080

ENTRYPOINT ["./main"]
