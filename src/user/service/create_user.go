package service

import (
	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/LaviqueDias/api-crud-go/src/configuration/rest_err"
	"github.com/LaviqueDias/api-crud-go/src/user/model"
	"go.uber.org/zap"
)

func (us *userServiceInterface) CreateUser(user *model.User) (*model.User, *rest_err.RestErr){
	logger.Info("Init CreateUser service",
		zap.String("journey", "createUser"),
	)

	user.EncryptPassword()

	logger.Info("CreateUser service executed succesfully", 
		zap.String("journey", "createUser"),
	)
	
	return user, nil
}