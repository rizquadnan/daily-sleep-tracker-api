package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/models"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/utils"
)

type UpdateUserRequestBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func prettyPrint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
}

func (h handler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	body := UpdateUserRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		utils.SetBadRequestJSON(c, "")
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	prettyPrint(body)

	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil {
		utils.SetStatusNotFoundJSON(c, "")
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Model(&user).Updates(models.User{Name: body.Name, Email: body.Email, PasswordHash: body.Password})

	c.JSON(http.StatusOK, UserToUserResponse(user))
}
