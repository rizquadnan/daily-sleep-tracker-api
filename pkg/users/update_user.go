package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/models"
)

type UpdateUserRequestBody struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func (h handler) UpdateUser (c *gin.Context) {
	id := c.Param("id")
	body := UpdateUserRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	user.Name = body.Name
	user.Email = body.Email
	user.PasswordHash = body.Password

	h.DB.Save(&user)

	c.JSON(http.StatusOK, &user)
}