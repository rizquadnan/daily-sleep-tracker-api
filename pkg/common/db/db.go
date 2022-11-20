package db

import (
	"log"

	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func Setup() {
	dbUrl := viper.Get("DB_URL").(string)

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Sleep{})

	DB = db;
}

func GetDB() *gorm.DB {
	return DB
}
