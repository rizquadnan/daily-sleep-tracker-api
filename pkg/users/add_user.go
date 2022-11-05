package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/models"
)

type AddUserRequestBody struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func (h handler) AddUser(c *gin.Context) {
	body := AddUserRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User

	user.Name = body.Name
	user.Email = body.Email
	user.PasswordHash = body.Password

	if result := h.DB.Create(&user); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, UserToUserResponse(user))
}