package user

import (
	"net/http"

	"github.com/gasuhwbab/goAuth/internal/database"
	"github.com/gasuhwbab/goAuth/internal/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// parse email & password
	var data struct {
		Email    string
		Password string
	}
	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to bind",
		})
		return
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// create user
	user := model.User{
		Email:    data.Email,
		Password: string(hash),
	}
	res := database.DB.Create(&user)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}
	// respond
	c.JSON(http.StatusOK, gin.H{
		"message": "User created",
	})
}
