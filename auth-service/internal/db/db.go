package db

import (
	"log"
	"viska/auth-service/internal/authentication"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&authentication.User{}); err != nil {
		log.Fatalf("failed to migrate user model: %v", err)
	}

	return db
}
