package infrastructure_user_gin

import (
	"net/http"
	"pruebas/internal/user/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getUserByIDHandler(uc application_user.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		u, err := uc.GetUserByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, u)
	}
}
