package authentication

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"` // hashed password
	CreatedAt time.Time
	UpdatedAt time.Time
}
