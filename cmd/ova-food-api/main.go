package main

import (
	"github.com/ozonva/ova-food-api/internal/api"
	ova_food_api "github.com/ozonva/ova-food-api/pkg/ova-food-api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	grpcPort = ":8080"
	grpcServerEndpoint = "localhost:8080"
	)

func main() {
	listen, err := net.Listen("tcp",grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v",err)
	}
	server:= grpc.NewServer()
	ova_food_api.RegisterOvaFoodApiServer(server,api.NewFoodAPI())
	reflection.Register(server)
	if err := server.Serve(listen); err != nil {
		log.Fatalf("failed to serveL %v", err)
	}
}
