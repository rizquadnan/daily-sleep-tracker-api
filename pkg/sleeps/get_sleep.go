package sleeps

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/models"
)

func (h handler) GetSleep (c *gin.Context) {
	id := c.Param("id")

	var sleep models.Sleep

	if result := h.DB.First(&sleep, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, SleepToSleepResponse(sleep))
}