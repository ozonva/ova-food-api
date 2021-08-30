package main

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-food-api/internal/repo"
	"github.com/rs/zerolog/log"
)

const (
	grpcPort = ":8080"
	dbHost = "localhost"
	dbPort = "5432"
	dbUser = "postgres"
	dbPassword = "postgres"
	dbName = "postgres"
	dbSslMode = "disable"
	dbDriver = "pgx"
	)

func main() {
	/*listen, err := net.Listen("tcp",grpcPort)
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v",err)
	}
	server:= grpc.NewServer()
*/
	psqlInfo := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSslMode)

	db, err := sqlx.Open(dbDriver, psqlInfo)
	if err != nil {
		log.Error().Err(err).Msgf("failed to create connect to database")
	}

	defer db.Close()
	/*err = db.Ping()
	if err != nil {
		log.Error().Err(err).Msgf("failed to ping to database")
	}*/
	r := repo.NewRepo(*db)
	f, err := r.DescribeEntity(1)
	if err != nil {
		log.Error().Err(err).Msgf("internal error")
	}
	fmt.Print(f)
	/*
	ova_food_api.RegisterOvaFoodApiServer(server,api.NewFoodAPI(r))
	reflection.Register(server)
	if err := server.Serve(listen); err != nil {
		log.Fatal().Msgf("failed to serveL %v", err)
	}*/
}
