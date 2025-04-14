package grpc

import (
	"context"
	"github.com/SilverNate/viska-proto/gen/auth"
	"viska/auth-service/internal/authentication"
)

type AuthGRPCServer struct {
	auth.UnimplementedAuthServiceServer
	AuthService authentication.AuthService
}

func NewAuthGRPCServer(authService authentication.AuthService) *AuthGRPCServer {
	return &AuthGRPCServer{
		AuthService: authService,
	}
}

func (s *AuthGRPCServer) ValidateToken(ctx context.Context, req *auth.TokenRequest) (*auth.TokenResponse, error) {
	userId, err := s.AuthService.ValidateJWT(req.Token)
	if err != nil {
		return &auth.TokenResponse{
			Valid: false,
			Error: "invalid or expired token",
		}, nil
	}

	return &auth.TokenResponse{
		Valid: true,
		Email: userId,
	}, nil
}
