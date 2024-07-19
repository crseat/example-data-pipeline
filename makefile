.PHONY: build run docker-up docker-down

build:
	go build -o bin/server main.go

run:
	go run main.go

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down
