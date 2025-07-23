package controller

import (
	"net/http"

	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/LaviqueDias/api-crud-go/src/configuration/rest_err"
	"github.com/LaviqueDias/api-crud-go/src/user/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init FindUserByEmail controller",
		zap.String("journey", "findUserByEmail"),
	)

	userEmail := c.Param("userEmail")
	if userEmail == "" {
		restErr := rest_err.NewBadRequestError("invalid user email parameter")
		logger.Error("Invalid userEmail parameter", restErr,
			zap.String("journey", "findUserByEmail"),
		)
		c.JSON(restErr.Code, restErr)
		return
	}

	user, restErr := uc.service.FindUserByEmail(userEmail)
	if restErr != nil {
		logger.Error("Error calling FindUserByEmail service", restErr,
			zap.String("journey", "findUserByEmail"),
		)
		c.JSON(restErr.Code, restErr)
		return
	}

	response := model.UserToUserResponse(user)

	logger.Info("FindUserByEmail controller executed successfully",
		zap.String("journey", "findUserByEmail"),
	)
	c.JSON(http.StatusOK, response)
}
