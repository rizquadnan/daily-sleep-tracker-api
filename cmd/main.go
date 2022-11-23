package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/auth"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/db"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/middlewares"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/sleeps"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/users"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	router := gin.Default()
	db.Setup()
	dbHandler := db.GetDB()

	auth.RegisterRoutes(router, dbHandler)

	router.Use(middlewares.JwtAuthMiddleware())
	users.RegisterRoutes(router, dbHandler)
	sleeps.RegisterRoutes(router, dbHandler)

	router.Run(port)
}