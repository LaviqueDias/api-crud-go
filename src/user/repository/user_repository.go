package repository

import (
	"database/sql"

	"github.com/LaviqueDias/api-crud-go/src/configuration/rest_err"
	"github.com/LaviqueDias/api-crud-go/src/user/model"
)

func NewUserReposirotyInterface(databaseConnection *sql.DB) UserRepositoryInterface {
	return &userRepositoryInterface{
		databaseConnection: databaseConnection,
	}
}

type userRepositoryInterface struct {
	databaseConnection *sql.DB
}

type UserRepositoryInterface interface {
	CreateUser(user *model.User) (*model.User, *rest_err.RestErr)
	GetAllUsers() (*[]model.User, *rest_err.RestErr)
}