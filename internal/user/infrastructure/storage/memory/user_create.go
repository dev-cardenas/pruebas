package memory_user

import (
	"errors"
	domain "pruebas/internal/user/domain"
)

func (r *InMemoryUserRepository) Create(user *domain.User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("user name and email are required")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.users = append(r.users, user)
	return nil
}
