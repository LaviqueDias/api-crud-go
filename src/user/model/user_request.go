package model


type UserRequest struct {
	Name     string `json:"name" binding:"required,min=4,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=50,containsany=!?*#@&$"`
	Role     string `json:"role" binding:"required,oneof=admin user"`
	Active   bool   `json:"active"`
}

func UserRequestToUser(userRequest UserRequest) *User {
	return &User{
		Name:      userRequest.Name,
		Email:     userRequest.Email,
		Password: userRequest.Password,
		Role:      userRequest.Role,
		Active:    userRequest.Active,
	}
}
