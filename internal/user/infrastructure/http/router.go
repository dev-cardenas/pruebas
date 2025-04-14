package infrastructure_user_gin

import (
	"pruebas/internal/user/infrastructure/bootstrap"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRouter(db *pgxpool.Pool) *gin.Engine {
	r := gin.Default()

	uc := infrastructure_user_bootstrap.SetupUserUseCase(db)

	RegisterRoutes(r, uc)

	return r
}
