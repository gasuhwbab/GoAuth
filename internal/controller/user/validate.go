package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "You are successfully logged in",
	})
}
