package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/SergioVenicio/grpc_gtw/config"
	usersGRPC "github.com/SergioVenicio/grpc_gtw/grpc"
	"github.com/SergioVenicio/grpc_gtw/models"
	"github.com/SergioVenicio/grpc_gtw/repositories"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Server struct {
	usersGRPC.UnimplementedUserServiceServer
	users *repositories.Users
}

func NewServer(cfg *config.Config) *Server {
	return &Server{users: repositories.NewUserRepository(cfg)}
}

func (s *Server) CreateUser(ctx context.Context, req *usersGRPC.CreateUserRequest) (*usersGRPC.CreateUserResponse, error) {
	user := req.GetUser()
	err := s.users.Save(models.User{ID: user.Id, Name: user.Name, Email: user.Email})
	if err != nil {
		return nil, err
	}
	return &usersGRPC.CreateUserResponse{User: user}, nil
}

func (s *Server) GetUser(ctx context.Context, req *usersGRPC.GetUserRequest) (*usersGRPC.GetUserResponse, error) {
	user, err := s.users.Get(req.GetId())
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return &usersGRPC.GetUserResponse{User: &usersGRPC.User{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}}, nil
}

func (s *Server) UpdateUser(ctx context.Context, req *usersGRPC.UpdateUserRequest) (*usersGRPC.UpdateUserResponse, error) {
	user := req.GetUser()
	err := s.users.Update(models.User{
		ID:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	})
	if err != nil {
		return nil, err
	}
	return &usersGRPC.UpdateUserResponse{User: user}, nil
}

func (s *Server) DeleteUser(ctx context.Context, req *usersGRPC.DeleteUserRequest) (*usersGRPC.DeleteUserResponse, error) {
	s.users.Delete(req.GetId())
	return &usersGRPC.DeleteUserResponse{Success: true}, nil
}

func setStatus(ctx context.Context, w http.ResponseWriter, m protoreflect.ProtoMessage) error {
	switch m.(type) {
	case *usersGRPC.CreateUserResponse:
		w.WriteHeader(http.StatusCreated)
	case *usersGRPC.DeleteUserRequest:
		w.WriteHeader(http.StatusNoContent)
	case *usersGRPC.UpdateUserResponse:
		w.WriteHeader(http.StatusAccepted)
	}
	return nil
}

func RunGRPCServer(cfg *config.Config) {
	grpcServer := grpc.NewServer()
	usersGRPC.RegisterUserServiceServer(grpcServer, NewServer(cfg))
	reflection.Register(grpcServer)
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())

	lis, err := net.Listen("tcp", cfg.GRPCServerPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("gRPC server listening on port", cfg.GRPCServerPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
