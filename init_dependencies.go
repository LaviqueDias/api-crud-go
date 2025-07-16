package main

import (
	"database/sql"

	"github.com/LaviqueDias/api-crud-go/src/user/controller"
	"github.com/LaviqueDias/api-crud-go/src/user/repository"
	"github.com/LaviqueDias/api-crud-go/src/user/service"
)

func initDependencies(database *sql.DB) controller.UserControllerInterface {
	repo := repository.NewUserReposirotyInterface(database)
	service := service.NewUserServiceInterface(repo)
	return controller.NewUserControllerInterface(service)

}