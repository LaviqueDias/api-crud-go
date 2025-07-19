package repository

import (
	"github.com/LaviqueDias/api-crud-go/src/configuration/rest_err"
	"github.com/LaviqueDias/api-crud-go/src/user/model"
)

func (ur *userRepositoryInterface) GetAllUsers() (*[]model.User, *rest_err.RestErr) {
	return nil, nil
}