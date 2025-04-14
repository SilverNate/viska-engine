package main

import (
	authpb "github.com/SilverNate/viska-proto/gen/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net"
	"viska/auth-service/config"
	"viska/auth-service/internal/authentication"
	"viska/auth-service/internal/db"
	authGrpc "viska/auth-service/internal/grpc"
	"viska/auth-service/internal/middleware"
)

func main() {
	cfg := config.Load()
	dbCon := db.InitDB(cfg.DBDSN)

	db.MigrateAndSeed(dbCon)

	// Wire dependencies
	repo := authentication.NewUserRepository(dbCon)
	authService := authentication.NewAuthService(*cfg, repo)
	authHandler := authentication.NewAuthHandler(authService)

	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()
		authpb.RegisterAuthServiceServer(grpcServer, authGrpc.NewAuthGRPCServer(authService))

		log.Println("gRPC server listening on port 50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	limiter := middleware.NewClientLimiter(1, 5)

	r := gin.Default()
	r.POST("/login", limiter.Middleware(), authHandler.Login)
	r.GET("/health", authHandler.HealthCheck)

	r.Run(":8080")
}
