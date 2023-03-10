package sleeps

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	h := handler{
		DB: db,
	}

	routes := r.Group("/sleeps")
	routes.POST("", h.AddSleep)
	routes.GET("", h.GetSleeps)
	routes.GET(":id", h.GetSleep)
	routes.PATCH(":id", h.UpdateSleep)
	routes.DELETE(":id", h.DeleteSleep)
}