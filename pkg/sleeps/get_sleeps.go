package sleeps

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/models"
)

func (h handler) GetSleeps(c *gin.Context) {
	var sleeps []models.Sleep

	if result := h.DB.Find(&sleeps); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, SleepsToSleepsResponse(sleeps))
}