package service

import (
	"context"
	"errors"
	"github/com/protocol10/deep-dive-auth-lab/auth/models"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterUser(req models.UserRegisterRequest) error
	LoginUser(email, password string) error // Change it later
}

type Service struct {
	db  *pgxpool.Pool
	ctx context.Context
}

func NewAuthService(db *pgxpool.Pool, ctx context.Context) *Service {
	return &Service{
		db:  db,
		ctx: ctx,
	}
}

func (s *Service) RegisterUser(req models.UserRegisterRequest) error {
	var exists bool

	// 1. Check if the user exists
	err := s.db.QueryRow(s.ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", req.EmailID).Scan(&exists)
	if err != nil {
		return err
	}

	// 2. Return an error if they do
	if exists {
		return errors.New("user already exists")
	}

	// 3. Insert the new user if they don't
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return err
	}

	_, err = s.db.Exec(s.ctx, "INSERT INTO users (email, password_hash) VALUES ($1, $2)", req.EmailID, passwordHash)
	return err
}

func (s *Service) LoginUser(email, password string) error {
	return nil
}
