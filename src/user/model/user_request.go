package model


// UserRequest represents the payload to create or update a user
type UserRequest struct {
	// Full name of the user
	// required: true
	// min length: 4
	// max length: 100
	// example: John Doe
	Name string `json:"name" binding:"required,min=4,max=100"`

	// Email address of the user
	// required: true
	// example: john.doe@example.com
	Email string `json:"email" binding:"required,email"`

	// Password for the user account
	// required: true
	// example: MyP@ssw0rd!
	Password string `json:"password" binding:"required,min=6,max=50,containsany=!?*#@&$"`

	// Role assigned to the user (admin or user)
	// required: true
	// enum: admin,user
	// example: user
	Role string `json:"role" binding:"required,oneof=admin user"`

	// Whether the user is active
	// example: true
	Active bool `json:"active"`
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
