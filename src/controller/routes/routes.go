package routes

import (
	"github.com/LaviqueDias/api-crud-go/src/controller/user"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.RouterGroup){

	r.GET("/", user.FindAllUsers)
	r.GET("/id/:userId", user.FindUserById)
	r.GET("/email/:userEmail", user.FindUserByEmail)
	r.POST("/", user.CreateUser)
	r.PUT("/:userId", user.UpdateUser)
	r.DELETE("/:userId", user.DeleteUser)

}