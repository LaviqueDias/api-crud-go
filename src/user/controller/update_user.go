package controller

import (
	"net/http"
	"strconv"

	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/LaviqueDias/api-crud-go/src/configuration/rest_err"
	"github.com/LaviqueDias/api-crud-go/src/configuration/validation"
	"github.com/LaviqueDias/api-crud-go/src/user/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
    userIdParam := c.Param("userId")
    logger.Info("Init UpdateUser controller",
        zap.String("journey", "updateUser"),
        zap.String("param_userId", userIdParam),
    )

    id, errConv := strconv.Atoi(userIdParam)
    if errConv != nil {
        restErr := rest_err.NewBadRequestError("invalid user id parameter")
        logger.Error("Invalid userId parameter", errConv,
            zap.String("journey", "updateUser"),
        )
        c.JSON(restErr.Code, restErr)
        return
    }

    var userRequest model.UserRequest
    if errBind := c.ShouldBindJSON(&userRequest); errBind != nil {
        logger.Error("Error trying to validate user info", errBind,
            zap.String("journey", "updateUser"),
        )
        errRest := validation.ValidateUserError(errBind)
        c.JSON(errRest.Code, errRest)
        return
    }

    user := model.UserRequestToUser(userRequest)
    user.ID = id

    users, restErr := uc.service.UpdateUser(id, user)
    if restErr != nil {
        logger.Error("Error calling UpdateUser service", restErr,
            zap.String("journey", "updateUser"),
        )
        c.JSON(restErr.Code, restErr)
        return
    }

    responses := make([]model.UserResponse, len(users))
    for i, u := range users {
        responses[i] = model.UserToUserResponse(u)
    }

    // 7) Log de sucesso e retorno
    logger.Info("UpdateUser controller executed successfully",
        zap.String("journey", "updateUser"),
        zap.Int("user_id", id),
    )
    c.JSON(http.StatusOK, responses)
}