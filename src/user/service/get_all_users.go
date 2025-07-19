package service

import (
	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/LaviqueDias/api-crud-go/src/configuration/rest_err"
	"github.com/LaviqueDias/api-crud-go/src/user/model"
	"go.uber.org/zap"
)

func (us *userServiceInterface) GetAllUsers() (*[]model.User, *rest_err.RestErr) {
	logger.Info("Init GetAllUsers service",
		zap.String("journey", "getAllUsers"),
	)

	users, err := us.repository.GetAllUsers()
	if err != nil {
		logger.Error("Error getting all users from repository",
		err,
		zap.String("journey", "getAllUsers"),
		)
		return nil, rest_err.NewInternalServerError("error getting users from repository")
	}

	logger.Info("GetAllUsers service executed succesfully", 
			zap.String("journey", "getAllUsers"),
	)

	return users, nil
}