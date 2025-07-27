package controller

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/LaviqueDias/api-crud-go/src/configuration/rest_err"
	"github.com/LaviqueDias/api-crud-go/src/configuration/validation"
	"github.com/LaviqueDias/api-crud-go/src/user/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
    logger.Info("Init LoginUser controller",
        zap.String("journey", "loginUser"),
    )

    // 1) Bind + validação do JSON
    var userLoginRequest model.UserLoginRequest
    if err := c.ShouldBindJSON(&userLoginRequest); err != nil {
        logger.Error("Error trying to validate user info", err,
            zap.String("journey", "loginUser"),
        )
        errRest := validation.ValidateUserError(err)
        c.JSON(errRest.Code, errRest)
        return
    }

    // 2) Busca o usuário no banco
    userLogin := model.UserLoginRequestToUser(userLoginRequest)
    user, restErr := uc.service.FindUserByEmail(userLogin.Email)
    if restErr != nil {
        logger.Error("Error trying to find user in DB", restErr,
            zap.String("journey", "loginUser"),
        )
        c.JSON(restErr.Code, restErr)
        return
    }

	logger.Info(fmt.Sprintf("user found by email %v", user))

    // 3) Compara senha
    if !user.ComparePassword(userLogin.Password) {
        restErr := rest_err.NewUnauthorizedError("invalid credentials")
        logger.Info("Password does not match",
            zap.String("journey", "loginUser"),
        )
        c.JSON(restErr.Code, restErr)
        return
    }

    // 4) Gera token JWT
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Email,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		logger.Error("Error signing JWT token", err,
			zap.String("journey", "loginUser"),
		)
		restErr := rest_err.NewInternalServerError("error generating token")
		c.JSON(restErr.Code, restErr)
		return
	}

    // 5) Retorna resposta de sucesso
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

    logger.Info("LoginUser controller executed successfully",
        zap.String("journey", "loginUser"),
        zap.String("email", user.Email),
    )
    c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
