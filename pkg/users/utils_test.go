package users

import (
	"testing"
	"time"

	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func TestUserToUserResponse(t *testing.T) {
	currentTime := time.Now()

	assert.Equal(t,
		UserToUserResponse(models.User{
			Model:        gorm.Model{ID: 11},
			Name:         "testName",
			Email:        "testEmail",
			PasswordHash: "testPassword",
			Sleeps: []models.Sleep{
				{
					UserID:      11,
					DATE:        datatypes.Date(currentTime),
					SLEEP_START: datatypes.NewTime(1, 1, 0, 0),
					SLEEP_END:   datatypes.NewTime(2, 1, 0, 0),
				},
			},
		}),
		map[string]any{
			"id":    uint(11),
			"name":  "testName",
			"email": "testEmail",
			"sleeps": []models.Sleep{
				{
					UserID:      11,
					DATE:        datatypes.Date(currentTime),
					SLEEP_START: datatypes.NewTime(1, 1, 0, 0),
					SLEEP_END:   datatypes.NewTime(2, 1, 0, 0),
				},
			},
		},
		"Wrong return value")
}

func TestUserToUserEmptySleepResponse(t *testing.T) {
	assert.Equal(t,
		UserToUserResponse(models.User{
			Model:        gorm.Model{ID: 11},
			Name:         "testName",
			Email:        "testEmail",
			PasswordHash: "testPassword",
			Sleeps:       []models.Sleep{},
		}),
		map[string]any{
			"id":     uint(11),
			"name":   "testName",
			"email":  "testEmail",
			"sleeps": []models.Sleep{},
		},
		"Wrong return value")
}

func TestUsersToUsersResponse(t *testing.T) {
	currentTime := time.Now()

	var sleep1 models.Sleep = models.Sleep{
		UserID:      1,
		DATE:        datatypes.Date(currentTime),
		SLEEP_START: datatypes.NewTime(1, 1, 0, 0),
		SLEEP_END:   datatypes.NewTime(2, 1, 0, 0),
	}

	var user1 models.User = models.User{
		Model: gorm.Model{
			ID: 1,
		},
		Name:         "name1",
		Email:        "email1",
		PasswordHash: "password1",
		Sleeps:       []models.Sleep{sleep1},
	}

	var sleep2 models.Sleep = models.Sleep{
		UserID:      2,
		DATE:        datatypes.Date(currentTime),
		SLEEP_START: datatypes.NewTime(1, 1, 0, 0),
		SLEEP_END:   datatypes.NewTime(2, 1, 0, 0),
	}

	var user2 models.User = models.User{
		Model: gorm.Model{
			ID: 2,
		},
		Name:         "name2",
		Email:        "email2",
		PasswordHash: "password2",
		Sleeps: []models.Sleep{
			sleep2,
		},
	}

	assert.Equal(
		t,
		UsersToUsersResponse([]models.User{
			user1,
			user2,
		}),
		[]map[string]any{
			{
				"id":     uint(1),
				"name":   "name1",
				"email":  "email1",
				"sleeps": []models.Sleep{sleep1},
			},
			{
				"id":     uint(2),
				"name":   "name2",
				"email":  "email2",
				"sleeps": []models.Sleep{sleep2},
			},
		},
		"Wrong return value",
	)
}
