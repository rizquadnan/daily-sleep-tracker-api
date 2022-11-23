package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/auth"
)

func JwtAuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		if (auth.IsTokenValid(c)) {
			c.Next()
		} else {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
	}
}