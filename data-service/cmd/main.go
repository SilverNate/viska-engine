package main

import (
	"log"
	"viska/data-service/config"
	"viska/data-service/internal/database"
	"viska/data-service/internal/router"
)

func main() {
	cfg := config.Load()

	db := database.InitDB(cfg.DBDSN)

	r, grpcConn := router.InitRouter(db)
	if err := r.Run(":8081"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
	defer grpcConn.Close()

	r.Run(":8081")
}
