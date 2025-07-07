package controller

import (
	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/LaviqueDias/api-crud-go/src/configuration/validation"
	"github.com/LaviqueDias/api-crud-go/src/controller/model/request"
	"github.com/LaviqueDias/api-crud-go/src/model"
	"github.com/LaviqueDias/api-crud-go/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)


func (uc *userControllerInterface) CreateUser(c *gin.Context){
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil{
		logger.Error("Error trying to validate user info", err,
		zap.String("journey", "createUser"),
	)
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Name,
		userRequest.Email,
		userRequest.Password,
		userRequest.Role,
		userRequest.Active,
	)
	if err := uc.service.CreateUser(domain); err != nil{
		logger.Info("Failed in created user",
		zap.String("journey", "createUser"),
	)
		c.JSON(err.Code, err)
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"),
	)

	c.JSON(200, view.ConvertDomainToResponse(domain))

}