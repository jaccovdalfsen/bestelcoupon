FROM golang:1.24 AS builder
WORKDIR /app
COPY . .
RUN go mod init coupon-service && go mod tidy
RUN go build -o coupon-service

FROM gcr.io/distroless/static:latest@sha256:d9f9472a8f4541368192d714a995eb1a99bab1f7071fc8bde261d7eda3b667d8
COPY --from=builder /app/coupon-service /
ENTRYPOINT ["/coupon-service"]