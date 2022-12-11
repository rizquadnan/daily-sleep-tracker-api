package sleeps

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/models"
)

func (h handler) GetSleeps(c *gin.Context) {
	var sleeps []models.Sleep
	userId := c.Query("user")

	if (userId == "") {
		if result := h.DB.Find(&sleeps); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}
	}	else {
		userIdInt, err := strconv.Atoi(userId);
		if (err != nil) {
			c.AbortWithError(http.StatusBadRequest, err)
		}

		if result := h.DB.Where(models.Sleep{ UserID: uint(userIdInt)}).Find(&sleeps); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}
	}

	c.JSON(http.StatusOK, SleepsToSleepsResponse(sleeps))
}
