package controller

import (
	"net/http"

	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/LaviqueDias/api-crud-go/src/configuration/validation"
	"github.com/LaviqueDias/api-crud-go/src/user/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary Create a new user
// @Description Creates a new user with the provided information and returns a list of created users
// @Tags Users
// @Accept json
// @Produce json
// @Param request body model.UserRequest true "User data to be created"
// @Success 200 {array} model.UserResponse
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /user [post]
func (uc *userControllerInterface) CreateUser(c *gin.Context){
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
	var userRequest model.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil{
		logger.Error("Error trying to validate user info", err,
		zap.String("journey", "createUser"),
	)
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	user := model.UserRequestToUser(userRequest)

	users, err := uc.service.CreateUser(user)
	if err != nil{
		logger.Error("Error trying to call CreateUser service", err,
		zap.String("journey", "createUser"),
	)
		c.JSON(err.Code, err)
	}

	responses := make([]model.UserResponse, len(users))
    for i, u := range users {
		responses[i] = model.UserToUserResponse(u)
    }

	logger.Info("CreateUser controller executed succesfully", 
		zap.String("journey", "createUser"),
	)

	c.JSON(http.StatusOK, responses)

}