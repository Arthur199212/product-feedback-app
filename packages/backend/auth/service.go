package auth

import users "product-feedback/user"

type AuthService struct {
	userRepo *users.UserRepository
}

func NewAuthService(userRepo *users.UserRepository) *AuthService {
	return &AuthService{userRepo}
}
