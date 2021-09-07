package internal

//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozonva/ova-food-api/internal/repo Repo
//go:generate mockgen -destination=./mocks/producer_mock.go -package=mocks github.com/ozonva/ova-food-api/internal/Kafka/producer Producer
