package application_user_test

import (
	application "pruebas/internal/user/application"
	domain "pruebas/internal/user/domain"
	"pruebas/internal/user/infrastructure/storage/memory"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListUsers_Success(t *testing.T) {
	repo := memory_user.NewInMemoryUserRepository()
	service := domain.NewService(repo)
	useCase := application.NewUseCase(repo, service)

	dto1 := application.CreateUserDTO{
		Name:     "John",
		LastName: "Doe",
		Email:    "john@example.com",
		Age:      25,
	}
	dto2 := application.CreateUserDTO{
		Name:     "Jane",
		LastName: "Smith",
		Email:    "jane@example.com",
		Age:      30,
	}

	_, err1 := useCase.CreateUser(dto1)
	assert.NoError(t, err1)

	_, err2 := useCase.CreateUser(dto2)
	assert.NoError(t, err2)

	users, err := useCase.ListUsers(1, 10)
	assert.NoError(t, err)
	assert.Len(t, users, 2)

	assert.Equal(t, "John", users[0].Name)
	assert.Equal(t, "Jane", users[1].Name)
}
