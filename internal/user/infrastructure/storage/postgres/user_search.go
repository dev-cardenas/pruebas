package postgres

import (
	"context"
	"pruebas/internal/user/domain"
)

func (r *UserRepository) Search(name, lastname string, offset, limit int) ([]*domain_user.User, error) {
	query := `SELECT id, name, lastname, email, age FROM users WHERE 
						name ILIKE '%' || $1 || '%' AND 
						lastname ILIKE '%' || $2 || '%' 
						LIMIT $3 OFFSET $4`

	rows, err := r.db.Query(context.Background(), query, name, lastname, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain_user.User
	for rows.Next() {
		var u domain_user.User
		err := rows.Scan(&u.ID, &u.Name, &u.LastName, &u.Email, &u.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}

	return users, nil
}
