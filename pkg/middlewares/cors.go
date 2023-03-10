package middlewares

import (
	"net/http"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowCredentials = true
	corsConfig.AllowOrigins = []string{"http://localhost:3000", "http://localhost:3001", "https://daily-sleep-tracker.vercel.app"}
	corsConfig.AllowHeaders = []string{"Authorization", "Content-Type"}
	corsConfig.AllowMethods = []string{http.MethodPost, http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions}
	return cors.New(corsConfig)
}
