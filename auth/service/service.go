package service

import (
	"errors"
	"github/com/protocol10/deep-dive-auth-lab/auth/models"
	"github/com/protocol10/deep-dive-auth-lab/auth/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterUser(req models.UserRegisterRequest) error
	LoginUser(email, password string) error // Change it later
}

type Service struct {
	repo repository.UserRepository
}

const (
	UserAlreadyExistsError = "user already exists"
)

func NewAuthService(repo repository.UserRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) RegisterUser(req models.UserRegisterRequest) error {

	// 1. Check if the user exists
	exists, err := s.repo.ExistsByEmail(req.EmailID)
	if err != nil {
		return err
	}

	// 2. Return an error if they do
	if exists {
		return errors.New(UserAlreadyExistsError)
	}

	// 3. Insert the new user if they don't
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return err
	}
	err = s.repo.CreateUser(models.User{
		Email:        req.EmailID,
		PasswordHash: string(passwordHash),
	})
	return err
}

func (s *Service) LoginUser(email, password string) error {
	return nil
}
