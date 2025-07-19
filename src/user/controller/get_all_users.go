package controller

import (
	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) GetAllUsers(c *gin.Context) {
	logger.Info("Init GetAllUsers controller",
		zap.String("journey", "getAllUsers"),
	)
}