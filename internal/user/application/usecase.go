package application_user

import "pruebas/internal/user/domain"

type UseCase interface {
	CreateUser(dto CreateUserDTO) (*domain_user.User, error)
	ListUsers(page, limit int) ([]*domain_user.User, error)
	GetUserByID(id int64) (*domain_user.User, error)
	UpdateUser(u *domain_user.User) error
	DeleteUser(id int64) error
}

type useCaseImpl struct {
	repo    domain_user.Repository
	service domain_user.Service
}

func NewUseCase(repo domain_user.Repository, service domain_user.Service) UseCase {
	return &useCaseImpl{
		repo:    repo,
		service: service,
	}
}
