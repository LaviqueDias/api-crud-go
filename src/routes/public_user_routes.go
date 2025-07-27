package routes

import (
	"github.com/LaviqueDias/api-crud-go/src/user/controller"
	"github.com/gin-gonic/gin"
)

func InitPublicUserRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	r.POST("/login", userController.LoginUser)
}