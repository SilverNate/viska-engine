package authentication

import (
	"fmt"
	jwt "github.com/golang-jwt/jwt/v5"
	"time"
	"viska/auth-service/config"
)

type AuthServiceImpl struct {
	config   config.Config
	repoUser UserRepository
}

func NewAuthService(config config.Config, repoUser UserRepository) *AuthServiceImpl {
	return &AuthServiceImpl{config: config, repoUser: repoUser}
}

func (s *AuthServiceImpl) GetUserByEmail(email string) (*User, error) {
	return s.repoUser.GetUserByEmail(email)
}

func (s *AuthServiceImpl) GenerateJWT(user *User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.config.JWTSecret))
}

func (s *AuthServiceImpl) ValidateJWT(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.config.JWTSecret), nil
	})

	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token claims")
	}

	userID := fmt.Sprintf("%v", claims["user_id"])
	if !ok {
		return "", fmt.Errorf("user_id not found in token")
	}

	return userID, nil
}
