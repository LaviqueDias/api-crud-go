package routes

import (
	"github.com/LaviqueDias/api-crud-go/src/user/controller"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface){

	r.GET("/", userController.GetAllUsers)
	// r.GET("/id/:userId", userController.FindUserById)
	// r.GET("/email/:userEmail", userController.FindUserByEmail)
	r.POST("/", userController.CreateUser)
	r.PUT("/:userId", userController.UpdateUser)
	r.DELETE("/:userId", userController.DeleteUserById )

}