.PHONY: build,run,test

build:
	GOOS=linux go build -o ./bin/main ./cmd/ova-food-api

run:
	go run ./cmd/ova-food-api

test:
	go test ./internal/utils -cover