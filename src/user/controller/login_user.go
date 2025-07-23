package controller

import (
	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/LaviqueDias/api-crud-go/src/configuration/validation"
	"github.com/LaviqueDias/api-crud-go/src/user/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	//Get the email and pass off req
	logger.Info("Init LoginUser controller",
		zap.String("journey", "loginUser"),
	)

	var userLoginRequest model.UserLoginRequest

	if err := c.ShouldBindJSON(&userLoginRequest); err != nil{
		logger.Error("Error trying to validate user info", err,
		zap.String("journey", "loginUser"),
	)
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	//user := model.UserLoginRequestToUser(userLoginRequest)

	//Look up requested user

	//Compare sent in pass with saved user pass hash

	//Generate a jwt token

	//send it back
}