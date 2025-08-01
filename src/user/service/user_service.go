package service

import (
	"github.com/LaviqueDias/api-crud-go/src/configuration/rest_err"
	"github.com/LaviqueDias/api-crud-go/src/user/model"
	"github.com/LaviqueDias/api-crud-go/src/user/repository"
)

func NewUserServiceInterface(repository repository.UserRepositoryInterface) UserServiceInterface {
	return &userServiceInterface{
		repository: repository,
	}
}

type userServiceInterface struct {
	repository repository.UserRepositoryInterface
}

type UserServiceInterface interface {
	CreateUser(user *model.User) ([]*model.User, *rest_err.RestErr)
	GetAllUsers() ([]*model.User, *rest_err.RestErr)
	DeleteUserById(int) ([]*model.User, *rest_err.RestErr)
	UpdateUser(id int, user *model.User) ([]*model.User, *rest_err.RestErr)
	FindUserByEmail(email string) (*model.User, *rest_err.RestErr)
	LoginUser(userLogin *model.User) (string, *rest_err.RestErr)
}