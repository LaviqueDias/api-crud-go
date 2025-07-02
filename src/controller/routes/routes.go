package routes

import (
	"github.com/LaviqueDias/api-crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.RouterGroup){

	r.GET("/", controller.FindAllUsers)
	r.GET("/id/:userId", controller.FindUserById)
	r.GET("/email/:userEmail", controller.FindUserByEmail)
	r.POST("/", controller.CreateUser)
	r.PUT("/:userId", controller.UpdateUser)
	r.DELETE("/:userId", controller.DeleteUser)

}