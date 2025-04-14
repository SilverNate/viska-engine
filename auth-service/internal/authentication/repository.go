package authentication

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(user *User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByEmail(email string) (*User, error) {
	var user User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CreateUser(user *User) error {
	err := r.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
