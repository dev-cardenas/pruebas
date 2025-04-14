package infrastructure_user_gin

import (
	"net/http"
	"pruebas/internal/user/application"
	"pruebas/internal/user/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

func updateUserHandler(uc application_user.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		var u domain_user.User
		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		u.ID = id
		err := uc.UpdateUser(&u)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusOK)
	}
}
