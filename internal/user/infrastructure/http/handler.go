package infrastructure_user_gin

import (
	"pruebas/internal/user/application"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, uc application_user.UseCase) {
	r.POST("/users", createUserHandler(uc))
	r.GET("/users/:id", getUserByIDHandler(uc))
	r.PUT("/users/:id", updateUserHandler(uc))
	r.DELETE("/users/:id", deleteUserHandler(uc))
}
