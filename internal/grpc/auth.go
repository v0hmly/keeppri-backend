package grpc

import (
	"context"
	"errors"

	"github.com/v0hmly/keeppri-backend/internal/grpc/pb"
	"github.com/v0hmly/keeppri-backend/internal/lib/grpc_errors"
	"github.com/v0hmly/keeppri-backend/internal/repository/domain"
	"github.com/v0hmly/keeppri-backend/internal/repository/postgres"
	"github.com/v0hmly/keeppri-backend/internal/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (h *Handler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	SessionId, err := h.services.AuthService.Login(req.Email, req.Password)
	if err != nil {
		if errors.Is(err, services.ErrLoginCredsInvalid) {
			return nil, status.Errorf(codes.NotFound, "invalid login credentials")
		}
		return nil, status.Errorf(codes.Internal, "failed to login user")
	}

	return &pb.LoginResponse{SessionId: *SessionId}, nil
}

func (h *Handler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	if req.FirstName == "" {
		return nil, status.Error(codes.InvalidArgument, "first name is required")
	}

	if req.LastName == "" {
		return nil, status.Error(codes.InvalidArgument, "last name is required")
	}

	UserId, err := h.services.AuthService.Register(&domain.User{
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})
	if err != nil {
		if errors.Is(err, postgres.ErrUserExists) {
			return nil, status.Errorf(codes.AlreadyExists, "user already exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to register user")
	}

	return &pb.RegisterResponse{UserId: *UserId}, nil
}

func (h *Handler) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata.FromIncomingContext: %v", grpc_errors.ErrNoCtxMetaData)
	}

	sessionID := md.Get("session_id")[0]
	if sessionID == "" {
		return nil, status.Errorf(codes.PermissionDenied, "md.Get sessionId: %v", grpc_errors.ErrInvalidSessionId)
	}

	err := h.services.AuthService.Logout(sessionID)
	if err != nil {
		return nil, err
	}

	logoutResponse := &pb.LogoutResponse{}

	return logoutResponse, nil
}
