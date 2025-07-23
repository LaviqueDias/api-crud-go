package service

import (
	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/LaviqueDias/api-crud-go/src/configuration/rest_err"
	"github.com/LaviqueDias/api-crud-go/src/user/model"
	"go.uber.org/zap"
)

func (us *userServiceInterface) FindUserByEmail(email string) (*model.User, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail service",
		zap.String("journey", "findUserByEmail"),
	)

	user, err := us.repository.FindUserByEmail(email)
	if err != nil {
		logger.Error("Error finding user by email",
			err,
			zap.String("journey", "findUserByEmail"),
		)
		return nil, err
	}

	logger.Info("FindUserByEmail service executed successfully",
		zap.String("journey", "findUserByEmail"),
	)

	return user, nil
}