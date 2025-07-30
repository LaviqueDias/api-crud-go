package controller

import (
	"net/http"

	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) LogoutUser(c *gin.Context) {
	logger.Info("Init LogoutUser controller",
        zap.String("journey", "logoutUser"),
    )

	user, _ := c.Get("user")

	logger.Info("User logged out", zap.Any("user", user))

	// Invalida o cookie
	c.SetCookie("Authorization", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})

	logger.Info("LogoutUser controller executed successfully",
        zap.String("journey", "logoutUser"),
    )
}
