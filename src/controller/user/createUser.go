package user

import (
	"log"
	"time"

	"github.com/LaviqueDias/api-crud-go/src/configuration/validation"
	"github.com/LaviqueDias/api-crud-go/src/model/request"
	"github.com/LaviqueDias/api-crud-go/src/model/response"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context){
	log.Println("Init CreateUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil{
		log.Printf("Error trying to marshal object, error=%s", err.Error())
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	userResponse := response.UserResponse{
		ID:        "te",
		Name:      userRequest.Name,
		Email:     userRequest.Email,
		Role:      userRequest.Role,
		Active:    userRequest.Active,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	c.JSON(200, userResponse)

}