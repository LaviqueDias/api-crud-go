package controller

import (
	"net/http"
	"strconv"

	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/LaviqueDias/api-crud-go/src/configuration/rest_err"
	"github.com/LaviqueDias/api-crud-go/src/user/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUserById(c *gin.Context) {
	logger.Info("Init DeleteUserById controller",
		zap.String("journey", "deleteUserById"),
	)

	userId := c.Param("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		restErr := rest_err.NewBadRequestError("invalid user id parameter")
		logger.Error("Invalid userId parameter", err,
			zap.String("journey", "deleteUserById"),
		)
		c.JSON(restErr.Code, restErr)
		return
	}

	// 3) Chama o service
	users, restErr := uc.service.DeleteUserById(id)
	if restErr != nil {
		logger.Error("Error deleting user by id", err,
			zap.String("journey", "deleteUserById"),
		)
		c.JSON(restErr.Code, restErr)
		return
	}

	responses := make([]model.UserResponse, len(users))
	for i, u := range users {
		responses[i] = model.UserToUserResponse(u)
	}

	logger.Info("DeleteUserById controller executed successfully",
		zap.String("journey", "deleteUserById"),
		zap.Int("user_id", id),
	)

	c.JSON(http.StatusOK, responses)
}
