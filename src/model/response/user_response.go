package response

import "time"

type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"         `
	Email     string    `json:"email"`
	Role      string    `json:"role" `
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}