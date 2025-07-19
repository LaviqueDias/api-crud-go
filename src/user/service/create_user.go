package service

import (
	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/LaviqueDias/api-crud-go/src/configuration/rest_err"
	"github.com/LaviqueDias/api-crud-go/src/user/model"
	"go.uber.org/zap"
)

func (us *userServiceInterface) CreateUser(user *model.User) ([]*model.User, *rest_err.RestErr){
	logger.Info("Init CreateUser service",
		zap.String("journey", "createUser"),
	)

	user.EncryptPassword()

	users, err := us.repository.CreateUser(user)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "createUser"))
		return nil, err
	}
	
	logger.Info("CreateUser service executed succesfully", 
		zap.String("journey", "createUser"),
	)
	
	return users, err
}