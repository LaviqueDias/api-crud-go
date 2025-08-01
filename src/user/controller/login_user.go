package controller

import (
	"net/http"

	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/LaviqueDias/api-crud-go/src/configuration/validation"
	"github.com/LaviqueDias/api-crud-go/src/user/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary User login
// @Description Authenticates a user and returns a JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body model.UserLoginRequest true "Login information"
// @Success 200 {object} model.UserLoginResponse
// @Failure 401 {object} rest_err.RestErr
// @Router /user/login [post]
func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init LoginUser controller",
		zap.String("journey", "loginUser"),
	)

	// Bind + validação do JSON
	var userLoginRequest model.UserLoginRequest
	if err := c.ShouldBindJSON(&userLoginRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "loginUser"),
		)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	// Busca o usuário no banco
	userLogin := model.UserLoginRequestToUser(userLoginRequest)

	tokenString, err := uc.service.LoginUser(userLogin)
	if err != nil {
		logger.Error("Error trying to call LoginUser service", err,
			zap.String("journey", "loginUser"),
		)
		c.JSON(err.Code, err)
	}

	// Retorna resposta de sucesso
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	logger.Info("LoginUser controller executed successfully",
		zap.String("journey", "loginUser"),
	)
	// Retorna o token jwt
	c.JSON(http.StatusOK, model.UserLoginResponse{
		Token: tokenString,
	})
}
