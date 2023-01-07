package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/models"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/utils"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/users"
)

type RegisterBody struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h handler) Register (c *gin.Context) {
	var body RegisterBody

	if err := c.ShouldBindJSON(&body); err != nil {
		utils.SetBadRequestJSON(c, "")
		return
	}

	var user models.User

	user.Name = body.Name
	user.Email = body.Email

	hashedPassword, err := decryptString(body.Password)
	if err != nil {
		utils.SetInternalServerErrorJSON(c, "")
	}
	user.PasswordHash = string(hashedPassword)

	if result := h.DB.Create(&user); result.Error != nil {
		utils.SetInternalServerErrorJSON(c, "")
		return
	}

	c.JSON(http.StatusOK, users.UserToUserResponse(user))
}