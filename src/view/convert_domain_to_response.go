package view

import (
	"time"

	"github.com/LaviqueDias/api-crud-go/src/controller/model/response"
	"github.com/LaviqueDias/api-crud-go/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
    now := time.Now()
    return response.UserResponse{
        ID:        "",            // preencha aqui com o ID se houver um método GetID
        Name:      userDomain.GetName(),
        Email:     userDomain.GetEmail(),
        Role:      userDomain.GetRole(),
        Active:    userDomain.IsActive(),
        CreatedAt: now,
        UpdatedAt: now,
    }
}