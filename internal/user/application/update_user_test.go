package application_user_test

import (
	application "pruebas/internal/user/application"
	domain "pruebas/internal/user/domain"
	"pruebas/internal/user/infrastructure/storage/memory"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateUser_Success(t *testing.T) {
	repo := memory_user.NewInMemoryUserRepository()
	service := domain.NewService(repo)
	useCase := application.NewUseCase(repo, service)

	createDTO := application.CreateUserDTO{
		Name:     "Alice",
		LastName: "Smith",
		Email:    "alice@example.com",
		Age:      28,
	}
	user, err := useCase.CreateUser(createDTO)
	assert.NoError(t, err)

	user.Name = "Alicia"
	user.Age = 30

	err = useCase.UpdateUser(user)
	assert.NoError(t, err)

	updated, err := useCase.GetUserByID(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Alicia", updated.Name)
	assert.Equal(t, 30, updated.Age)
}

func TestUpdateUser_NotFound(t *testing.T) {
	repo := memory_user.NewInMemoryUserRepository()
	service := domain.NewService(repo)
	useCase := application.NewUseCase(repo, service)

	user := &domain.User{
		ID:       10,
		Name:     "Ghost",
		LastName: "User",
		Email:    "ghost@example.com",
		Age:      40,
	}

	err := useCase.UpdateUser(user)
	assert.Error(t, err)
	assert.EqualError(t, err, "user not found")
}
