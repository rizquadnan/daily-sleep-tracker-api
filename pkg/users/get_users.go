package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/models"
)

func (h handler) GetUsers(c *gin.Context) {
	var users []models.User

	if result := h.DB.Model(models.User{}).Find(&users); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, UsersToUsersResponse(users))
}
