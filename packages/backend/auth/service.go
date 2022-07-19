package auth

import users "product-feedback/user"

type AuthService interface {
}

type authService struct {
	repo users.UserRepository
}

func NewAuthService(userRepo users.UserRepository) AuthService {
	return &authService{userRepo}
}
