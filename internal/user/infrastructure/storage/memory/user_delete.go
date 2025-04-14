package memory_user

import (
	"errors"
	domain "pruebas/internal/user/domain"
)

func (r *InMemoryUserRepository) Delete(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, u := range r.users {
		if u.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

func (r *InMemoryUserRepository) List() ([]*domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.users, nil
}
