.PHONY: build,run

build:
	GOOS=linux go build -o ./bin/main ./cmd/ova-food-api

run:
	go run ./cmd/ova-food-api
