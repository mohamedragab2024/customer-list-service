# Build stage
FROM golang AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=1 go build -o customer-list-service cmd/main.go

# Run stage
FROM ubuntu
WORKDIR /root/
COPY --from=builder /app/customer-list-service .
COPY --from=builder /app/customers.db .
EXPOSE 8080
CMD ["./customer-list-service"]
