package postgres

import (
	"context"
	"pruebas/internal/user/domain"
)

func (r *UserRepository) Create(u *domain_user.User) error {
	query := `INSERT INTO users (name, lastname, email, age) VALUES ($1, $2, $3, $4) RETURNING id`
	return r.db.QueryRow(context.Background(), query, u.Name, u.LastName, u.Email, u.Age).Scan(&u.ID)
}
