package repository

import (
	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/LaviqueDias/api-crud-go/src/configuration/rest_err"
	"github.com/LaviqueDias/api-crud-go/src/user/model"
	"go.uber.org/zap"
)

func (ur *userRepositoryInterface) CreateUser(user *model.User) ([]*model.User, *rest_err.RestErr) {
    logger.Info("Init CreateUser repository",
        zap.String("journey", "createUser"),
    )

    // 1) Insere o usuário (sem created_at/updated_at nem ID)
    query := `INSERT INTO users(name, email, password, role, active)
              VALUES(?, ?, ?, ?, ?)`
    result, err := ur.databaseConnection.Exec(
        query,
        user.Name,
        user.Email,
        user.Password,
        user.Role,
        user.Active,
    )
    if err != nil {
        logger.Error("Error trying to insert user",
            err,
            zap.String("journey", "createUser"),
        )
        return nil, rest_err.NewInternalServerError(err.Error())
    }

    // 2) Captura o ID gerado
    lastID, err := result.LastInsertId()
    if err != nil {
        logger.Error("Error getting last insert id",
            err,
            zap.String("journey", "createUser"),
        )
        return nil, rest_err.NewInternalServerError(err.Error())
    }
    user.ID = int(lastID)

    // 3) Busca created_at e updated_at pelo ID
    row := ur.databaseConnection.QueryRow(
        "SELECT created_at, updated_at FROM users WHERE id = ?",
        lastID,
    )
    if err := row.Scan(&user.CreatedAt, &user.UpdatedAt); err != nil {
        logger.Error("Error scanning timestamps",
            err,
            zap.String("journey", "createUser"),
        )
        return nil, rest_err.NewInternalServerError(err.Error())
    }

    logger.Info("CreateUser repository executed successfully",
        zap.String("journey", "createUser"),
    )

    return ur.GetAllUsers()
}
