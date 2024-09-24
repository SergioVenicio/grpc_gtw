package main

import (
	"log"

	"github.com/SergioVenicio/grpc_gtw/server"
	"github.com/SergioVenicio/grpc_gtw/settings"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s := settings.NewSettings()
	go server.RunGRPCGWServer(s)
	server.RunGRPCServer(s)
}
