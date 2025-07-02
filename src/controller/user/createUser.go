package user

import (
	"time"

	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/LaviqueDias/api-crud-go/src/configuration/validation"
	"github.com/LaviqueDias/api-crud-go/src/model/request"
	"github.com/LaviqueDias/api-crud-go/src/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context){
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

	userResponse := response.UserResponse{
		ID:        "te",
		Name:      userRequest.Name,
		Email:     userRequest.Email,
		Role:      userRequest.Role,
		Active:    userRequest.Active,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"),
	)

	c.JSON(200, userResponse)

}