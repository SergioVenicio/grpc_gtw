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

func RunGRPCGWServer(ctx context.Context, cfg *config.Config) {
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

	server := http.Server{
		Addr:    cfg.HTTPServerAddr,
		Handler: mux,
	}
	go func() {
		<-ctx.Done()
		log.Println("stoping http server...")
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("failed to shutdown http server: %v", err)
		}
	}()

	log.Println("gRPC-gateway server listening on port", cfg.HTTPServerAddr)
	if err = server.ListenAndServe(); err != nil {
		log.Fatalf("failed to serve http server: %v", err)
	}
}
