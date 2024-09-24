package settings

import "os"

type Settings struct {
	ScylladbURI        string
	GRPCServerEndpoint string
	HTTPServerAddr     string
	GRPCServerPort     string
}

func NewSettings() *Settings {
	scylladbURI := os.Getenv("SCYLLADB_URI")
	GRPCServerEndpoint := os.Getenv("GRPC_SERVER_ENDPOINT")
	HTTPServerAddr := os.Getenv("HTTP_SERVER_ADDR")
	GRPCServerPort := os.Getenv("GRPC_SERVER_PORT")

	return &Settings{
		ScylladbURI:        scylladbURI,
		GRPCServerEndpoint: GRPCServerEndpoint,
		HTTPServerAddr:     HTTPServerAddr,
		GRPCServerPort:     GRPCServerPort,
	}
}
