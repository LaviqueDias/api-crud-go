package routes

import (
	"github.com/LaviqueDias/api-crud-go/src/user/controller"
	"github.com/gin-gonic/gin"
)

func InitPrivateUserRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	r.GET("/", userController.GetAllUsers)
	// r.GET("/id/:userId", userController.FindUserById)
	r.GET("/email/:userEmail", userController.FindUserByEmail)
	r.GET("/validate", userController.ValidateUser)
	r.POST("/", userController.CreateUser)
	r.PUT("/:userId", userController.UpdateUser)
	r.DELETE("/:userId", userController.DeleteUserById )
	r.POST("/logout", userController.LogoutUser)
}