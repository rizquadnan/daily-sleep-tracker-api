package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/auth"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/config"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/db"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/middlewares"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/sleeps"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/users"
)

func main() {
	config.InitConfig()

	config := config.GetConfig()

	port := config.PORT
	
	router := gin.Default()
	router.Use(middlewares.CorsMiddleware())
	apiRoutes := router.Group("/api/v1")

	db.Setup()
	dbHandler := db.GetDB()

	auth.RegisterRoutes(apiRoutes, dbHandler)

	apiRoutes.Use(middlewares.JwtAuthMiddleware())

	users.RegisterRoutes(apiRoutes, dbHandler)
	sleeps.RegisterRoutes(apiRoutes, dbHandler)

	router.Run(":" + port)
}
