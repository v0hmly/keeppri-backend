package grpc

import (
	"context"

	"github.com/v0hmly/keeppri-backend/internal/grpc/pb"
	"github.com/v0hmly/keeppri-backend/internal/lib/grpc_errors"
	"github.com/v0hmly/keeppri-backend/internal/repository/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (h *Handler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	SessionId, err := h.services.AuthService.Login(req.Email, req.Password)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "h.Auth.Login: %v", err)
	}

	loginResponse := &pb.LoginResponse{
		SessionId: *SessionId,
	}

	return loginResponse, nil
}

func (h *Handler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	UserId, err := h.services.AuthService.Register(&domain.User{
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})
	if err != nil {
		return nil, status.Errorf(codes.AlreadyExists, "h.Auth.Register: %v", err)
	}

	registerResponse := &pb.RegisterResponse{UserId: *UserId}

	return registerResponse, nil
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
