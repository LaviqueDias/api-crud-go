package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *userControllerInterface) ValidateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message" : "I'm logged in"})
}