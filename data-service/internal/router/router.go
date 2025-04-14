package router

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"time"
	authgrpc "viska/data-service/internal/grpc"
	"viska/data-service/internal/middleware"
	"viska/data-service/internal/redis"

	"viska/data-service/internal/library"
)

func InitRouter(db *gorm.DB) (*gin.Engine, *grpc.ClientConn) {
	var conn *grpc.ClientConn
	var err error
	for i := 0; i < 5; i++ {
		conn, err = grpc.Dial("auth-service:50051", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(3*time.Second))
		if err == nil {
			break
		}
		log.Printf("gRPC server not ready yet, retrying in 2s... (%v)", err)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatalf("could not connect to gRPC server: %v", err)
	}

	authClient := authgrpc.NewAuthGRPCClient(conn)

	r := gin.Default()

	redis := redis.NewRedisClient("redis:6379")

	libraryRepo := library.NewLibraryRepository(db)
	libraryService := library.NewLibraryService(libraryRepo)
	libraryHandler := library.NewLibraryHandler(libraryService, redis)
	
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware(authClient))

	author := api.Group("/authors")
	{
		author.POST("", libraryHandler.CreateAuthor)
		author.GET("", libraryHandler.GetAllAuthor)
		author.GET("/:id", libraryHandler.GetAuthorByID)
		author.PUT("/:id", libraryHandler.UpdateAuthor)
		author.DELETE("/:id", libraryHandler.DeleteAuthor)
	}

	publisher := api.Group("/publishers")
	{
		publisher.POST("", libraryHandler.CreatePublisher)
		publisher.GET("", libraryHandler.GetAllPublisher)
		publisher.GET("/:id", libraryHandler.GetPublisherByID)
		publisher.PUT("/:id", libraryHandler.UpdatePublisher)
		publisher.DELETE("/:id", libraryHandler.DeletePublisher)
	}

	book := api.Group("/books")
	{
		book.POST("", libraryHandler.CreateBook)
		book.GET("", libraryHandler.GetAllBook)
		book.GET("/:id", libraryHandler.GetBookByID)
		book.PUT("/:id", libraryHandler.UpdateBook)
		book.DELETE("/:id", libraryHandler.DeleteBook)
	}

	return r, conn
}
