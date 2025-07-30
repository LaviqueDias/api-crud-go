package model

import "time"

// UserResponse represents the data returned for a user
type UserResponse struct {
	// Unique identifier of the user
	// example: 1
	ID int `json:"id"`

	// Full name of the user
	// example: John Doe
	Name string `json:"name"`

	// Email address of the user
	// example: john.doe@example.com
	Email string `json:"email"`

	// Role assigned to the user
	// example: user
	Role string `json:"role"`

	// Whether the user is active
	// example: true
	Active bool `json:"active"`

	// Timestamp when the user was created
	// example: 2025-07-30T10:15:30Z
	CreatedAt time.Time `json:"created_at"`

	// Timestamp when the user was last updated
	// example: 2025-07-30T12:45:00Z
	UpdatedAt time.Time `json:"updated_at"`
}

func UserToUserResponse(user *User) UserResponse{
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		Active:    user.Active,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
