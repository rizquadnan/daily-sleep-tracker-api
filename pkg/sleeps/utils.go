package sleeps

import (
	"strings"

	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/models"
	"gorm.io/datatypes"
)

func trimQuotes (str string) string {
	return str[1 : len(str)-1]
}

func timeToHHMM(time datatypes.Time) string {
	timeArray := strings.Split(time.String(), ":");

	return timeArray[0] + ":" + timeArray[1]
}

func dateToDDMMYY(date datatypes.Date) (string, error){
	dateVal, err := date.MarshalJSON()
	if (err != nil) {
		return "", err
	}

	dateStr := strings.Split(trimQuotes(string(dateVal[:])), "T")[0]
	dateArray := strings.Split(dateStr, "-")

	YY := dateArray[0]
	MM := dateArray[1]
	DD := dateArray[2]

	return (DD + "-" + MM + "-" + YY), err
}

func SleepToSleepResponse (sleep models.Sleep) map[string]any {
	date, _ := dateToDDMMYY(sleep.DATE);

	return map[string]any{
		"id":     sleep.ID,
		"date":   date,
		"sleepStart":  timeToHHMM(sleep.SLEEP_START),
		"sleepEnd": timeToHHMM(sleep.SLEEP_END),
		"sleepDuration": sleep.SLEEP_DURATION,
		"userId": sleep.UserID,
	}
}

func SleepsToSleepsResponse (sleeps []models.Sleep) []map[string]any {
	var response []map[string]any

	for _, sleep := range sleeps {
		response = append(response, SleepToSleepResponse(sleep))
	}

	return response
}