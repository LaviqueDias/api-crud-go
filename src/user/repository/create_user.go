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

	stmt, err := ur.databaseConnection.Prepare("insert into products(id, name, price) values(?, ?, ?)") 
	if err != nil{
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user)
	if err != nil{
		return nil, err
	}

	return nil, nil
}