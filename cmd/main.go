package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/db"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/sleeps"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/users"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	router := gin.Default()
	dbHandler := db.Init(dbUrl)

	users.RegisterRoutes(router, dbHandler)
	sleeps.RegisterRoutes(router, dbHandler)

	router.Run(port)
}