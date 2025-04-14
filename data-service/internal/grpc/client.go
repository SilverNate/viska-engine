package grpc

import (
	"context"
	authpb "github.com/SilverNate/viska-proto/gen/auth"
	"log"

	"google.golang.org/grpc"
)

type AuthClient interface {
	ValidateToken(ctx context.Context, token string) (bool, error)
}

type AuthGRPCClient struct {
	client authpb.AuthServiceClient
}

func NewAuthGRPCClient(conn *grpc.ClientConn) *AuthGRPCClient {
	return &AuthGRPCClient{
		client: authpb.NewAuthServiceClient(conn),
	}
}

func (a *AuthGRPCClient) ValidateToken(ctx context.Context, token string) (bool, error) {
	res, err := a.client.ValidateToken(ctx, &authpb.TokenRequest{Token: token})
	if err != nil {
		log.Println("gRPC error:", err)
		return false, err
	}
	return res.Valid, nil
}
