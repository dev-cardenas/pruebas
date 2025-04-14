package postgres

import (
	"pruebas/internal/user/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) domain_user.Repository {
	return &UserRepository{db: db}
}
