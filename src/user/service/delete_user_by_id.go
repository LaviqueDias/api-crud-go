package service

import (
	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/LaviqueDias/api-crud-go/src/configuration/rest_err"
	"github.com/LaviqueDias/api-crud-go/src/user/model"
	"go.uber.org/zap"
)

func (us *userServiceInterface) DeleteUserById(id int) ([]*model.User, *rest_err.RestErr) {
	logger.Info("Init DeleteUserById service",
		zap.String("journey", "deleteUserById"),
	)

	users, err := us.repository.DeleteUserById(id)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "deleteUserById"))
		return nil, err
	}

	logger.Info("DeleteUserById service executed succesfully", 
		zap.String("journey", "deleteUserById"),
	)

	return users, nil
}