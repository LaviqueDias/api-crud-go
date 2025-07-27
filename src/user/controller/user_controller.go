package controller

import (
	"github.com/LaviqueDias/api-crud-go/src/user/service"
	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(service service.UserServiceInterface) UserControllerInterface {
	return &userControllerInterface{
		service: service,
	}
}

type userControllerInterface struct {
	service service.UserServiceInterface
}

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	GetAllUsers(c *gin.Context)
	DeleteUserById(c *gin.Context)
	UpdateUser(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	LoginUser(c *gin.Context)
	ValidateUser(c *gin.Context)

	GetUserService() service.UserServiceInterface
}

func (uc *userControllerInterface) GetUserService() service.UserServiceInterface{
	return uc.service
}