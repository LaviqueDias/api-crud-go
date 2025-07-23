package repository

import (
	"database/sql"
	"fmt"

	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/LaviqueDias/api-crud-go/src/configuration/rest_err"
	"github.com/LaviqueDias/api-crud-go/src/user/model"
	"go.uber.org/zap"
)

func (ur *userRepositoryInterface) FindUserByEmail(email string) (*model.User, *rest_err.RestErr) {
    logger.Info("Init FindUserByEmail repository",
        zap.String("journey", "findUserByEmail"),
    )

    query := `
        SELECT id, name, email, password, role, active, created_at, updated_at
          FROM users
         WHERE email = ?
    `
    row := ur.databaseConnection.QueryRow(query, email)

    user := &model.User{}
    if err := row.Scan(
        &user.ID,
        &user.Name,
        &user.Email,
        &user.Password,
        &user.Role,
        &user.Active,
        &user.CreatedAt,
        &user.UpdatedAt,
    ); err != nil {
        if err == sql.ErrNoRows {
            logger.Info("User not found by email",
                zap.String("journey", "findUserByEmail"),
            )
            return nil, rest_err.NewNotFoundError(fmt.Sprintf("user with email %s not found", email))
        }
        logger.Error("Error scanning user by email",
            err,
            zap.String("journey", "findUserByEmail"),
        )
        return nil, rest_err.NewInternalServerError("error finding user by email")
    }

    logger.Info("FindUserByEmail repository executed successfully",
        zap.String("journey", "findUserByEmail"),
    )
	
    return user, nil
}