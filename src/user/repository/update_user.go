package repository

import (
	"fmt"

	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/LaviqueDias/api-crud-go/src/configuration/rest_err"
	"github.com/LaviqueDias/api-crud-go/src/user/model"
	"go.uber.org/zap"
)

func (ur *userRepositoryInterface) UpdateUser(id int, user *model.User) ([]*model.User, *rest_err.RestErr) {
    logger.Info("Init UpdateUser repository",
        zap.String("journey", "updateUser"),
    )

    query := `
        UPDATE users
        SET name      = ?,
            email     = ?,
            password  = ?,
            role      = ?,
            active    = ?,
            updated_at = CURRENT_TIMESTAMP
        WHERE id = ?
    `
    result, err := ur.databaseConnection.Exec(
        query,
        user.Name,
        user.Email,
        user.Password,
        user.Role,
        user.Active,
        id,
    )
    if err != nil {
        logger.Error("Error executing update in UpdateUser",
            err,
            zap.String("journey", "updateUser"),
        )
        return nil, rest_err.NewInternalServerError("error executing update")
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        logger.Error("Error fetching rowsAffected in UpdateUser",
            err,
            zap.String("journey", "updateUser"),
        )
        return nil, rest_err.NewInternalServerError("error checking update result")
    }
    if rowsAffected == 0 {
        logger.Info("No user updated (not found) in UpdateUser",
            zap.String("journey", "updateUser"),
        )
        return nil, rest_err.NewNotFoundError(fmt.Sprintf("user with id %d not found", id))
    }

    users, restErr := ur.GetAllUsers()
    if restErr != nil {
        return nil, restErr
    }

    logger.Info("UpdateUser repository executed successfully",
        zap.String("journey", "updateUser"),
    )
    return users, nil
}