package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/models"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/utils"
)

func (h handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil {
		utils.SetStatusNotFoundJSON(c, "")
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&user)

	c.Status(http.StatusOK)
}