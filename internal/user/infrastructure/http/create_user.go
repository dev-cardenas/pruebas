package infrastructure_user_gin

import (
	"net/http"
	"pruebas/internal/user/application"

	"github.com/gin-gonic/gin"
)

func createUserHandler(uc application_user.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto application_user.CreateUserDTO
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		u, err := uc.CreateUser(dto)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, u)
	}
}
