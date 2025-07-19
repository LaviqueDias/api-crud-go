package controller

import (

	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/LaviqueDias/api-crud-go/src/configuration/validation"
	"github.com/LaviqueDias/api-crud-go/src/user/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)


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

	c.JSON(200, responses)

}