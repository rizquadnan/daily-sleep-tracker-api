package sleeps

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/models"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/utils"
)

func (h handler) GetSleep (c *gin.Context) {
	id := c.Param("id")

	var sleep models.Sleep

	if result := h.DB.First(&sleep, id); result.Error != nil {
		utils.SetStatusNotFoundJSON(c, "")
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, SleepToSleepResponse(sleep))
}