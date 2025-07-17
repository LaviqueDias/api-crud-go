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

	stmt, err := ur.databaseConnection.Prepare("insert into users(name, email, password, role, active) values(?, ?, ?, ?, ?)") 
	if err != nil{
		logger.Error("Error trying to create user",
			err,
			zap.String("journey", "createUser"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Password, user.Role, user.Active)
	if err != nil{
		logger.Error("Error trying to create user",
			err,
			zap.String("journey", "createUser"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	logger.Info(
		"CreateUser repository executed successfully",
		zap.String("journey", "createUser"))

	return user, nil
}