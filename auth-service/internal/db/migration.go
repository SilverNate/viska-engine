package db

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"viska/auth-service/internal/authentication"
)

func MigrateAndSeed(db *gorm.DB) {
	if err := db.AutoMigrate(&authentication.User{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	seedAdminUser(db)
}

func seedAdminUser(db *gorm.DB) {
	var count int64
	db.Model(&authentication.User{}).Where("email = ?", "admin@viska.io").Count(&count)
	if count > 0 {
		log.Println("Admin user already exists, skipping seeding.")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}

	user := authentication.User{
		Email:    "admin@viska.io",
		Password: string(hashedPassword),
	}

	if err := db.Create(&user).Error; err != nil {
		log.Fatalf("Failed to seed admin user: %v", err)
	}

	log.Println("Admin user seeded.")
}
