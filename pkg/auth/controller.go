package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(apiRoutes *gin.RouterGroup, db *gorm.DB) {
	h := handler{
		DB: db,
	}

	routes := apiRoutes.Group("/auth")
	routes.POST("/register", h.Register)
	routes.POST("/login", h.Login)
}