# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go run scripts/generate.go
RUN go build -o customer-list-service cmd/main.go

# Run stage
FROM alpine:3.14
WORKDIR /root/
COPY --from=builder /app/customer-list-service .
COPY --from=builder /app/customers.db .
EXPOSE 8080
CMD ["./customer-list-service"]
