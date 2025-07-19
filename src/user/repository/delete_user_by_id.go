package repository

import (
	"fmt"

	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/LaviqueDias/api-crud-go/src/configuration/rest_err"
	"github.com/LaviqueDias/api-crud-go/src/user/model"
	"go.uber.org/zap"
)

func (ur *userRepositoryInterface) DeleteUserById(id int) ([]*model.User, *rest_err.RestErr) {
    logger.Info("Init DeleteUserById repository",
        zap.String("journey", "deleteUserById"),
    )

    result, err := ur.databaseConnection.Exec(
        "DELETE FROM users WHERE id = ?",
        id,
    )
    if err != nil {
        logger.Error("Error executing delete in DeleteUserById",
            err,
            zap.String("journey", "deleteUserById"),
        )
        return nil, rest_err.NewInternalServerError("error executing delete")
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        logger.Error("Error fetching rowsAffected in DeleteUserById",
            err,
            zap.String("journey", "deleteUserById"),
        )
        return nil, rest_err.NewInternalServerError("error checking delete result")
    }
    if rowsAffected == 0 {
        logger.Info("No user deleted (not found) in DeleteUserById",
            zap.String("journey", "deleteUserById"),
        )
        return nil, rest_err.NewNotFoundError(fmt.Sprintf("user with id %d not found", id))
    }

    users, restErr := ur.GetAllUsers()
    if restErr != nil {
        return nil, restErr
    }

    logger.Info("DeleteUserById repository executed successfully",
        zap.String("journey", "deleteUserById"),
        zap.Int("user_id", id),
    )
    return users, nil
}
