package users

import (
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/models"
)

func UserToUserResponse(user models.User) map[string]any {
	return map[string]any{
		"id":     user.ID,
		"name":   user.Name,
		"email":  user.Email,
	}
}

func UsersToUsersResponse(users []models.User) []map[string]any {
	var response []map[string]any

	for _, user := range users {
		response = append(response, UserToUserResponse(user))
	}

	return response
}