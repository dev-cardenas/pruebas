package memory_user

import (
	"pruebas/internal/user/domain"
	"sync"
)

type InMemoryUserRepository struct {
	mu    sync.Mutex
	users []*domain_user.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: []*domain_user.User{},
	}
}
