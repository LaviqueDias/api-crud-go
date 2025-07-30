package model


// UserLoginRequest represents the payload to authenticate a user
type UserLoginRequest struct {
	// Email address of the user
	// required: true
	// example: user@example.com
	Email string `json:"email" binding:"required,email"`

	// Password for the user account
	// required: true
	// example: MyS3cr3t!#
	Password string `json:"password" binding:"required,min=6,max=50,containsany=!?*#@&$"`
}


func UserLoginRequestToUser(userLoginRequest UserLoginRequest) *User {
	return &User{
		Email:     userLoginRequest.Email,
		Password: userLoginRequest.Password,
	}
}
