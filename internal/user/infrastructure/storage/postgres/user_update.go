package postgres

import (
	"context"
	"pruebas/internal/user/domain"
)

func (r *UserRepository) Update(u *domain_user.User) error {
	query := `UPDATE users SET name = $1, lastname = $2, email = $3, age = $4 WHERE id = $5`
	_, err := r.db.Exec(context.Background(), query, u.Name, u.LastName, u.Email, u.Age, u.ID)
	return err
}
