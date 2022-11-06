package sleeps

import "github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/models"

func SleepToSleepResponse (sleep models.Sleep) map[string]any {
	return map[string]any{
		"id":     sleep.ID,
		"date":   sleep.DATE,
		"sleepStart":  sleep.SLEEP_START,
		"sleepEnd": sleep.SLEEP_END,
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