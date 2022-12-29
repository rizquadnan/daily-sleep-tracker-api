package sleeps

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/models"
	"gorm.io/datatypes"
)

type AddSleepRequestBody struct {
	DATE string `json:"date"`
	SLEEP_START string `json:"sleepStart"`
	SLEEP_END string `json:"sleepEnd"`
	UserID int `json:"userId"`
}

const dateFormat = "2006-01-02";

func (h handler) AddSleep(c *gin.Context) {
	body := AddSleepRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var sleep models.Sleep

	date, err := time.Parse(dateFormat, body.DATE)
	if (err != nil) {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	sleep.DATE = datatypes.Date(date)

	sleepStartArray := strings.Split(body.SLEEP_START, ":")
	sleepStartHour, err := strconv.Atoi(sleepStartArray[0])
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	sleepStartMinutes, err := strconv.Atoi(sleepStartArray[1])
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	sleep.SLEEP_START = datatypes.NewTime(sleepStartHour, sleepStartMinutes, 0, 0)

	sleepEndArray := strings.Split(body.SLEEP_END, ":")
	sleepEndHour, err := strconv.Atoi(sleepEndArray[0])
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	sleepEndMinutes, err := strconv.Atoi(sleepEndArray[1])
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	sleep.SLEEP_END = datatypes.NewTime(sleepEndHour, sleepEndMinutes, 0, 0)

	sleepStartInRFC := body.DATE + "T" + body.SLEEP_START + ":00" + "Z";
	sleepStartInTime, errInStartParse := time.Parse(time.RFC3339, sleepStartInRFC)
	if (errInStartParse != nil) {
		c.AbortWithError(http.StatusBadRequest, errInStartParse)
	}

	sleepEndDate := date.AddDate(0, 0, 1)
	sleepEndInRFC := sleepEndDate.Format(dateFormat) + "T" + body.SLEEP_END + ":00" + "Z";
	sleepEndInTime, errInEndParse := time.Parse(time.RFC3339, sleepEndInRFC)
	if (errInEndParse != nil) {
		c.AbortWithError(http.StatusBadRequest, errInEndParse)
	}

	sleepTimeDiff := sleepEndInTime.Sub(sleepStartInTime)
	sleep.SLEEP_DURATION = int(sleepTimeDiff.Minutes())

	var user models.User
	if result := h.DB.First(&user, body.UserID); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	sleep.UserID = uint(body.UserID)

	if result := h.DB.Create(&sleep); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &sleep)
}