package service

import "github/com/protocol10/deep-dive-auth-lab/auth/models"

type AuthService interface {
	RegisterUser(req models.UserRegisterRequest) error
	LoginUser(email, password string) error // Change it later
}

type Service struct {
}

func NewAuthService() *Service {
	return &Service{}
}

func (a *Service) RegisterUser(req models.UserRegisterRequest) error {
	return nil
}

func (a *Service) LoginUser(email, password string) error {
	return nil
}
