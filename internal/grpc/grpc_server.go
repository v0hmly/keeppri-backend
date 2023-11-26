package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/v0hmly/keeppri-backend/internal/grpc/pb"
	"github.com/v0hmly/keeppri-backend/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type (
	AuthService interface {
		Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error)
		Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error)
		Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error)
	}

	Handler struct {
		services *services.Services
	}
)

func NewGrpcHandler(services *services.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Run(address string) error {
	op := "grpc.Run"

	var err error
	listen, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("%s: listen error: %w", op, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServer(grpcServer, h)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listen); err != nil {
		return fmt.Errorf("%s: serve error: %w", op, err)
	}
	return nil
}
