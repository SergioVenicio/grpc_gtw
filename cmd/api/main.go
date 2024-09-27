package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	ctx := context.Background()
	ctx, shutdown := context.WithTimeout(ctx, 10*time.Second)

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		shutdown()
	}()
	go server.RunGRPCGWServer(ctx, cfg)
	server.RunGRPCServer(ctx, cfg)

}
