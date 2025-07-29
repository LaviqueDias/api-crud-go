package service

import (
	"fmt"

	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/LaviqueDias/api-crud-go/src/configuration/rest_err"
	"github.com/LaviqueDias/api-crud-go/src/user/model"
	"go.uber.org/zap"
)

func (us *userServiceInterface) LoginUser(userLogin *model.User) (string, *rest_err.RestErr){
	user, err := us.repository.FindUserByEmail(userLogin.Email)
    if err != nil {
        logger.Error("Error trying to call FindUserByEmail repository", err,
            zap.String("journey", "loginUser"),
        )
        return "", err
    }

	logger.Info(fmt.Sprintf("user found by email %v", user),
		zap.String("journey", "loginUser"),
	)

    // Compara senha
    if !user.ComparePassword(userLogin.Password) {
        err := rest_err.NewUnauthorizedError("invalid credentials")
        logger.Info("Password does not match",
            zap.String("journey", "loginUser"),
        )
        return "", err
    }

    // Gera token JWT
    tokenString, err := user.GenerateToken()
    if err != nil {
        logger.Error("Error generating JWT token", err,
            zap.String("journey", "loginUser"),
        )
        return "", rest_err.NewInternalServerError("error generating token")
    }

	

	return tokenString, nil
}