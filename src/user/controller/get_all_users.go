package controller

import (
	"net/http"

	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/LaviqueDias/api-crud-go/src/user/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) GetAllUsers(c *gin.Context) {
	logger.Info("Init GetAllUsers controller",
		zap.String("journey", "getAllUsers"),
	)

	users, err := uc.service.GetAllUsers()
	if err != nil {
		logger.Error("Error trying to call GetAllUsers service", err,
			zap.String("journey", "getAllUsers"),
		)
		c.JSON(err.Code, err)
		return
	}
	
	responses := make([]model.UserResponse, len(users))
    for i, u := range users {
		responses[i] = model.UserToUserResponse(u)
    }
	
	logger.Info("GetAllUsers controller executed succesfully", 
		zap.String("journey", "getAllUsers"),
	)

	c.JSON(http.StatusOK, responses)
}