package controller

import (
	"net/http"

	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/LaviqueDias/api-crud-go/src/user/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary Logout user
// @Description Logs out the currently authenticated user by clearing the authentication cookie
// @Tags Authentication
// @Produce json
// @Success 200 {object} map[string]string
// @Router /user/logout [post]
func (uc *userControllerInterface) LogoutUser(c *gin.Context) {
	logger.Info("Init LogoutUser controller",
        zap.String("journey", "logoutUser"),
    )

	user, _ := c.Get("user")

	logger.Info("User logged out", zap.Any("user", model.UserToUserResponse(user.(*model.User))), zap.String("journey", "logoutUser"))

	// Invalida o cookie
	c.SetCookie("Authorization", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})

	logger.Info("LogoutUser controller executed successfully",
        zap.String("journey", "logoutUser"),
    )
}
