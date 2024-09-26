package server

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/SergioVenicio/grpc_gtw/config"
	usersGRPC "github.com/SergioVenicio/grpc_gtw/grpc"

	"github.com/flowchartsman/swaggerui"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RunGRPCGWServer(cfg *config.Config) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	rmux := runtime.NewServeMux(
		runtime.WithForwardResponseOption(setStatus),
	)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := usersGRPC.RegisterUserServiceHandlerFromEndpoint(ctx, rmux, cfg.GRPCServerEndpoint, opts)
	if err != nil {
		log.Fatalf("failed to register HTTP handlers: %v", err)
	}

	swaggerFile, err := os.Open("./swagger/user.swagger.json")
	if err != nil {
		log.Fatalf("failed to load swagger json: %v", err)
	}
	spec, err := io.ReadAll(swaggerFile)
	if err != nil {
		log.Fatalf("failed parse swagger data: %v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/swagger/", http.StripPrefix("/swagger", swaggerui.Handler(spec)))
	mux.Handle("/", rmux)

	log.Println("gRPC-gateway server listening on port", `s.HTTPServerAddr`)
	if err = http.ListenAndServe(cfg.HTTPServerAddr, mux); err != nil {
		log.Fatalf("failed to serve http server: %v", err)
	}
}
