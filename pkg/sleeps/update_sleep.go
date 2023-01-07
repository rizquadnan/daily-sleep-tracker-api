package sleeps

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/models"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/utils"
	"gorm.io/datatypes"
)

type UpdateSleepRequestBody struct {
	DATE        string `json:"date"`
	SLEEP_START string `json:"sleepStart"`
	SLEEP_END   string `json:"sleepEnd"`
}

func (h handler) UpdateSleep(c *gin.Context) {
	id := c.Param("id")
	body := UpdateSleepRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		utils.SetBadRequestJSON(c, "")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var sleep models.Sleep
	if result := h.DB.First(&sleep, id); result.Error != nil {
		utils.SetStatusNotFoundJSON(c, "")
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	var updatedSleep models.Sleep

	if body.DATE != "" {
		date, err := time.Parse(dateFormat, body.DATE)
		if err != nil {
			utils.SetBadRequestJSON(c, "")
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		updatedSleep.DATE = datatypes.Date(date)
	}

	if body.SLEEP_START != "" {
		sleepStartArray := strings.Split(body.SLEEP_START, ":")
		sleepStartHour, err := strconv.Atoi(sleepStartArray[0])
		if err != nil {
			utils.SetBadRequestJSON(c, "")
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		sleepStartMinutes, err := strconv.Atoi(sleepStartArray[1])
		if err != nil {
			utils.SetBadRequestJSON(c, "")
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		updatedSleep.SLEEP_START = datatypes.NewTime(sleepStartHour, sleepStartMinutes, 0, 0)
	}

	if body.SLEEP_END != "" {
		sleepEndArray := strings.Split(body.SLEEP_END, ":")
		sleepEndHour, err := strconv.Atoi(sleepEndArray[0])
		if err != nil {
			utils.SetBadRequestJSON(c, "")
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		sleepEndMinutes, err := strconv.Atoi(sleepEndArray[1])
		if err != nil {
			utils.SetBadRequestJSON(c, "")
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		updatedSleep.SLEEP_END = datatypes.NewTime(sleepEndHour, sleepEndMinutes, 0, 0)
	}

	h.DB.Model(&sleep).Updates(models.Sleep{DATE: updatedSleep.DATE, SLEEP_START: updatedSleep.SLEEP_START, SLEEP_END: updatedSleep.SLEEP_END})

	c.JSON(http.StatusOK, SleepToSleepResponse(sleep))
}
