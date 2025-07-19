package repository

import (
	"github.com/LaviqueDias/api-crud-go/src/configuration/logger"
	"github.com/LaviqueDias/api-crud-go/src/configuration/rest_err"
	"github.com/LaviqueDias/api-crud-go/src/user/model"
	"go.uber.org/zap"
)

func (ur *userRepositoryInterface) GetAllUsers() ([]*model.User, *rest_err.RestErr) {
	logger.Info("Init GetAllUsers service",
		zap.String("journey", "getAllUsers"),
	)

	query := `
      SELECT id, name, email, role, active, created_at, updated_at
      FROM users
    `

	rows, err := ur.databaseConnection.Query(query)
	if err != nil {
		logger.Error("Error executing query for GetAllUsers",
			err,
			zap.String("journey", "getAllUsers"),
		)
		return nil, rest_err.NewInternalServerError("error executing query")
	}
	defer rows.Close()

	var users []*model.User
	
	for rows.Next() {

		user := &model.User{}

		if err := rows.Scan(
            &user.ID,
            &user.Name,
            &user.Email,
            &user.Role,
            &user.Active,
            &user.CreatedAt,
            &user.UpdatedAt,
        ); err != nil {
            logger.Error("Error scanning row in GetAllUsers",
                err,
                zap.String("journey", "getAllUsers"),
            )
            return nil, rest_err.NewInternalServerError("error scanning row")
        }
		
		users = append(users, user)
	}

	logger.Info("GetAllUsers repository executed successfully",
        zap.String("journey", "getAllUsers"),
    )

	return users, nil
}