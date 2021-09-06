package main

import (
	"fmt"

	"net"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/ozonva/ova-food-api/internal/logger"

	"github.com/ozonva/ova-food-api/internal/utils"

	"github.com/ozonva/ova-food-api/internal/metrics"
	"github.com/ozonva/ova-food-api/internal/tracer"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/ozonva/ova-food-api/internal/Kafka/consumer"
	"github.com/ozonva/ova-food-api/internal/Kafka/producer"

	"github.com/Shopify/sarama"

	"github.com/ozonva/ova-food-api/internal/api"
	ova_food_api "github.com/ozonva/ova-food-api/pkg/ova-food-api"
	"google.golang.org/grpc/reflection"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-food-api/internal/repo"
	"google.golang.org/grpc"
)

func main() {
	config, err := utils.LoadConfig("./configs/config.yml")
	if err != nil {
		log.Fatal().Msgf("failed to read config")
	}

	logger.InitLogger(config.App.Logfile)

	tracer.InitTracing("food-api tracer")
	producerEx := initKafka(config)
	go initMetrics()

	listen, err := net.Listen("tcp", config.Grpc.GRPCPort)
	if err != nil {
		logger.GlobalLogger.Fatal().Msgf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Database.DBHost, config.Database.DBPort, config.Database.DBUser, config.Database.DBPassword,
		config.Database.DBName, config.Database.DBSslMode)
	db, err := sqlx.Open(config.Database.DBDriver, psqlInfo)
	if err != nil {
		logger.GlobalLogger.Error().Err(err).Msgf("failed to create connect to database")
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		logger.GlobalLogger.Error().Err(err).Msgf("failed to ping to database")
	}
	r := repo.NewRepo(db)

	ova_food_api.RegisterOvaFoodApiServer(server, api.NewFoodAPI(r, config.App.AppChunkSize, *producerEx))
	reflection.Register(server)
	if err := server.Serve(listen); err != nil {
		logger.GlobalLogger.Fatal().Msgf("failed to serveL %v", err)
	}
}

func initKafka(config *utils.Config) *producer.Producer {
	producerEx, err := producer.NewProducer([]string{config.Kafka.KafkaBroker}, config.Kafka.KafkaTopic)
	if err != nil {
		logger.GlobalLogger.Fatal().Msgf("failed to create producer: %v", err)
	}

	consumerEx, err := sarama.NewConsumer([]string{config.Kafka.KafkaBroker}, nil)
	if err != nil {
		logger.GlobalLogger.Fatal().Msgf("failed to create consumer: %v", err)
	}
	consumer.Subscribe(config.Kafka.KafkaTopic, consumerEx)
	return &producerEx
}
func initMetrics() {
	metrics.RegisterMetrics()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
