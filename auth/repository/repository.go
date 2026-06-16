package repository

import (
	"context"
	"github/com/protocol10/deep-dive-auth-lab/auth/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	CreateUser(user models.User) error
	GetUserByEmail(email string) (models.User, error)
	ExistsByEmail(email string) (bool, error)
}

type Repository struct {
	db  *pgxpool.Pool
	ctx context.Context
}

func NewRepository(db *pgxpool.Pool, ctx context.Context) *Repository {
	return &Repository{
		db:  db,
		ctx: ctx,
	}
}

func (r *Repository) CreateUser(user models.User) error {
	_, err := r.db.Exec(r.ctx, "INSERT INTO users (email, password_hash) VALUES ($1, $2)", user.Email, user.PasswordHash)
	return err
}

func (r *Repository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.QueryRow(r.ctx, "SELECT id, email, password_hash FROM users WHERE email = $1", email).
		Scan(&user.ID, &user.Email, &user.PasswordHash)
	return user, err
}

func (r *Repository) ExistsByEmail(email string) (bool, error) {
	var exists bool
	err := r.db.QueryRow(r.ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", email).Scan(&exists)
	return exists, err
}
