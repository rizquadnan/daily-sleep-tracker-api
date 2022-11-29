package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/common/models"
	"github.com/rizquadnan/daily-sleep-tracker-api/pkg/users"
	"golang.org/x/crypto/bcrypt"
)

type LoginBody struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h handler) Login (c *gin.Context) {
	var body LoginBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}
	err := h.DB.Model(models.User{}).Where("email = ?", body.Email).Take(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email or password not found"})
		return 
	}

	if verifyPassword(body.Password, user.PasswordHash) == bcrypt.ErrMismatchedHashAndPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email or password not found"})
		return
	}

	token, err := generateToken(user.ID)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Sorry failed to login, please try again"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "user": users.UserToUserResponse(user)})
}