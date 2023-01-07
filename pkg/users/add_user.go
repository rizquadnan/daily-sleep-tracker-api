package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/models"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/utils"
)

type AddUserRequestBody struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func (h handler) AddUser(c *gin.Context) {
	body := AddUserRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		utils.SetBadRequestJSON(c, "")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User

	user.Name = body.Name
	user.Email = body.Email
	user.PasswordHash = body.Password

	if result := h.DB.Create(&user); result.Error != nil {
		utils.SetInternalServerErrorJSON(c, "")
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusCreated, UserToUserResponse(user))
}