package users

import (
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/models"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/sleeps"
)

func UserToUserResponse(user models.User) map[string]any {
	return map[string]any{
		"id":     user.ID,
		"name":   user.Name,
		"email":  user.Email,
		"sleeps": sleeps.SleepsToSleepsResponse(user.Sleeps),
	}
}

func UsersToUsersResponse(users []models.User) []map[string]any {
	var response []map[string]any

	for _, user := range users {
		response = append(response, UserToUserResponse(user))
	}

	return response
}