package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/constant"
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
		c.JSON(http.StatusBadRequest, utils.GenerateErrorResponse(constant.GENERIC_NOT_VALID_PAYLOAD, http.StatusBadRequest))
		return
	}

	var user models.User

	user.Name = body.Name
	user.Email = body.Email

	hashedPassword, err := decryptString(body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.GenerateErrorResponse(constant.GENERIC_INTERNAL_SERVER_ERROR, http.StatusInternalServerError))
	}
	user.PasswordHash = string(hashedPassword)

	if result := h.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, utils.GenerateErrorResponse(constant.GENERIC_INTERNAL_SERVER_ERROR, http.StatusInternalServerError))
		return
	}

	c.JSON(http.StatusOK, users.UserToUserResponse(user))
}