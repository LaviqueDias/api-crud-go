package service

import (
	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/LaviqueDias/api-crud-go/src/configuration/rest_err"
	"github.com/LaviqueDias/api-crud-go/src/user/model"
	"go.uber.org/zap"
)

func (us *userServiceInterface) UpdateUser(id int, user *model.User) ([]*model.User, *rest_err.RestErr) {
	logger.Info("Init UpdateUser service",
		zap.String("journey", "updateUser"),
	)

	user.EncryptPassword()

	users, err := us.repository.UpdateUser(id, user)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "updateUser"))
		return nil, err
	}

	logger.Info("UpdateUser service executed successfully",
		zap.String("journey", "updateUser"),
	)

	return users, nil
}