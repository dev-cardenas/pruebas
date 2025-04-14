package infrastructure_user_gin

import (
	"net/http"
	"pruebas/internal/user/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

func deleteUserHandler(uc application_user.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		err := uc.DeleteUser(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	}
}
