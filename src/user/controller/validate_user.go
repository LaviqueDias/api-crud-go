package controller

import (
	"net/http"

	"github.com/LaviqueDias/api-crud-go/src/user/model"
	"github.com/gin-gonic/gin"
)

func (uc *userControllerInterface) ValidateUser(c *gin.Context) {
	user, _ := c.Get("user")

	response := model.UserToUserResponse(user.(*model.User))
	
	c.JSON(http.StatusOK, response)
}