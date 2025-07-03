package service

import (
	"fmt"

	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/LaviqueDias/api-crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
	"github.com/LaviqueDias/api-crud-go/src/model"
)


func(ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr{
	
	logger.Info("Init createUser model", zap.String("journey","createUser"))

	userDomain.EncryptPassword()
	fmt.Println(ud)
	return nil
}