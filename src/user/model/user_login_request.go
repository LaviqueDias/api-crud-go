package model


type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=50,containsany=!?*#@&$"`
}

func UserLoginRequestToUser(userLoginRequest UserLoginRequest) *User {
	return &User{
		Email:     userLoginRequest.Email,
		Password: userLoginRequest.Password,
	}
}
