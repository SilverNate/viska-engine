package authentication

//go:generate mockgen -destination=mocks/mock.go -package=mocks -source=interface.go
type AuthService interface {
	GetUserByEmail(email string) (*User, error)
	GenerateJWT(user *User) (string, error)
	ValidateJWT(tokenStr string) (string, error)
}

type AuthRepository interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(user *User) error
}
