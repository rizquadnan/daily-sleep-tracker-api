package db

import (
	"log"
	"time"

	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/models"
	"gorm.io/datatypes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Sleep{})

	user := models.User{
		Name:         "Budi",
		Email:        "budi@gmail.com",
		PasswordHash: "somepasswordhash",
		Sleeps: []models.Sleep{
			{
				DATE:        datatypes.Date(time.Now()),
				SLEEP_START: datatypes.NewTime(1, 2, 3, 0),
				SLEEP_END:   datatypes.NewTime(1, 2, 4, 0),
			},
		},
	}

	db.Create(&user)

	db.Save(&user)

	return db
}
