package application_user_test

import (
	application "pruebas/internal/user/application"
	domain "pruebas/internal/user/domain"
	"pruebas/internal/user/infrastructure/storage/memory"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	repo := memory_user.NewInMemoryUserRepository()
	service := domain.NewService(repo)

	useCase := application.NewUseCase(repo, service)

	dto := application.CreateUserDTO{
		Name:     "John",
		LastName: "Doe",
		Email:    "john@example.com",
		Age:      30,
	}

	user, err := useCase.CreateUser(dto)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "John", user.Name)
}

func TestCreateUser_IsAMinor(t *testing.T) {
	repo := memory_user.NewInMemoryUserRepository()
	service := domain.NewService(repo)

	useCase := application.NewUseCase(repo, service)

	dto := application.CreateUserDTO{
		Name:     "",
		LastName: "Doe",
		Email:    "john@example.com",
		Age:      17,
	}

	user, err := useCase.CreateUser(dto)
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.EqualError(t, err, "user must be at least 18 years old")
}

func TestCreateUser_EmailAlreadyExists(t *testing.T) {
	repo := memory_user.NewInMemoryUserRepository()
	service := domain.NewService(repo)

	useCase := application.NewUseCase(repo, service)

	dto1 := application.CreateUserDTO{
		Name:     "John",
		LastName: "Doe",
		Email:    "john@example.com",
		Age:      25,
	}
	user1, err1 := useCase.CreateUser(dto1)
	assert.NoError(t, err1)
	assert.NotNil(t, user1)

	dto2 := application.CreateUserDTO{
		Name:     "Jane",
		LastName: "Smith",
		Email:    "john@example.com",
		Age:      30,
	}
	user2, err2 := useCase.CreateUser(dto2)
	assert.Error(t, err2)
	assert.Nil(t, user2)
	assert.EqualError(t, err2, "email already in use")
}
