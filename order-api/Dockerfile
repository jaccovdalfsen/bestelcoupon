FROM golang:1.24@sha256:4c0a1814a7c6c65ece28b3bfea14ee3cf83b5e80b81418453f0e9d5255a5d7b8 AS builder
WORKDIR /app
COPY . .
RUN go mod init order-api && go mod tidy
RUN go build -o order-api

FROM gcr.io/distroless/static:latest@sha256:d9f9472a8f4541368192d714a995eb1a99bab1f7071fc8bde261d7eda3b667d8
COPY --from=builder /app/order-api /
ENTRYPOINT ["/order-api"]