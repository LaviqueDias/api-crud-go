package repository

import (
	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/LaviqueDias/api-crud-go/src/configuration/rest_err"
	"github.com/LaviqueDias/api-crud-go/src/user/model"
	"go.uber.org/zap"
)

func (ur *userRepositoryInterface) CreateUser(user *model.User) (*model.User, *rest_err.RestErr) {
	logger.Info("Init CreateUser repository",
		zap.String("journey", "createUser"),
	)

	return nil, nil
}