package main

import (
	"fmt"
	"net"

	"github.com/ozonva/ova-food-api/internal/api"
	ova_food_api "github.com/ozonva/ova-food-api/pkg/ova-food-api"
	"google.golang.org/grpc/reflection"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-food-api/internal/repo"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

const (
	grpcPort   = ":8080"
	dbHost     = "localhost"
	dbPort     = "5432"
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbName     = "postgres"
	dbSslMode  = "disable"
	dbDriver   = "pgx"

	chunkSize = 2
)

func main() {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSslMode)
	db, err := sqlx.Open(dbDriver, psqlInfo)
	if err != nil {
		log.Error().Err(err).Msgf("failed to create connect to database")
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Error().Err(err).Msgf("failed to ping to database")
	}
	r := repo.NewRepo(db)

	ova_food_api.RegisterOvaFoodApiServer(server, api.NewFoodAPI(r, chunkSize))
	reflection.Register(server)
	if err := server.Serve(listen); err != nil {
		log.Fatal().Msgf("failed to serveL %v", err)
	}
}
