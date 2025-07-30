package controller

import (
	"net/http"

	"github.com/LaviqueDias/api-crud-go/src/user/model"
	"github.com/gin-gonic/gin"
)

// @Summary Validate user session
// @Description Returns the authenticated user's information based on the current session or token
// @Tags Authentication
// @Produce json
// @Success 200 {object} model.UserResponse
// @Failure 401 {object} rest_err.RestErr
// @Router /user/validate [get]
func (uc *userControllerInterface) ValidateUser(c *gin.Context) {
	user, _ := c.Get("user")

	response := model.UserToUserResponse(user.(*model.User))
	
	c.JSON(http.StatusOK, response)
}