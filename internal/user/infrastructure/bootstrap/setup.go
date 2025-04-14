package infrastructure_user_bootstrap

import (
	"pruebas/internal/user/application"
	domain_user "pruebas/internal/user/domain"
	"pruebas/internal/user/infrastructure/storage/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupUserUseCase(db *pgxpool.Pool) application_user.UseCase {
	repo := postgres.NewUserRepository(db)
	service := domain_user.NewService(repo)

	return application_user.NewUseCase(repo, service)
}
