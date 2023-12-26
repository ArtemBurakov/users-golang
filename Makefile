include .env

docker-all: docker-build docker-run

docker-build:
	docker build --build-arg PORT=$(PORT) --build-arg GIN_MODE=$(GIN_MODE) -t users-app .

docker-run:
	docker run -dp $(PORT):$(PORT) --name users-golang users-app

run:
	go run cmd/main.go
