.PHONY: build test docker-build up down rebuild

run:
	go run cmd/main.go

db: 
	go run scripts/db.go

build:
	go build -o bin/customer-list-service cmd/main.go

test:
	go test ./...

docker-build:
	docker build -t customer-list-service:latest .
