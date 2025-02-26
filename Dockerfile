FROM golang:1.23 AS builder

WORKDIR /app

COPY . .

RUN go build -o prop-filter

# lighter image for execution
FROM debian:latest
WORKDIR /app

COPY --from=builder /app/prop-filter /app/prop-filter

ENTRYPOINT ["/app/prop-filter"]
