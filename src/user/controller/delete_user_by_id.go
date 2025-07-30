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

// @Summary Delete a user by ID
// @Description Deletes the user with the given ID and returns the remaining users
// @Tags Users
// @Produce json
// @Param userId path int true "ID of the user to delete"
// @Success 200 {array} model.UserResponse
// @Failure 400 {object} rest_err.RestErr
// @Failure 404 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /user/{userId} [delete]
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
