.PHONY: up down run test lint

up:
	docker-compose up  --remove-orphans --build

down:
	docker-compose down

run:
	go build -o app cmd/port-service/main.go && HTTP_ADDR=:8080 ./app

test:
	go test -race ./...

lint:
	golangci-lint run
