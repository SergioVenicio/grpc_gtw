package main

import (
	"log"

	"github.com/SergioVenicio/grpc_gtw/config"
	"github.com/SergioVenicio/grpc_gtw/server"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	cfg := config.NewConfig()
	go server.RunGRPCGWServer(cfg)
	server.RunGRPCServer(cfg)
}
