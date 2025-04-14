package postgres

import (
	"context"
	"pruebas/internal/user/domain"
)

func (r *UserRepository) FindByEmail(email string) (*domain_user.User, error) {
	query := `SELECT id, name, lastname, email, age FROM users WHERE email = $1`

	row := r.db.QueryRow(context.Background(), query, email)

	var u domain_user.User
	err := row.Scan(&u.ID, &u.Name, &u.LastName, &u.Email, &u.Age)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}

func (r *UserRepository) GetByID(id int64) (*domain_user.User, error) {
	query := `SELECT id, name, lastname, email, age FROM users WHERE id = $1`
	u := &domain_user.User{}
	err := r.db.QueryRow(context.Background(), query, id).Scan(&u.ID, &u.Name, &u.LastName, &u.Email, &u.Age)
	if err != nil {
		return nil, err
	}
	return u, nil
}
