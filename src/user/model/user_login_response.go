package model

// UserLoginResponse represents the response returned after a successful login
type UserLoginResponse struct {
	// JWT token issued upon successful authentication
	// example: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
	Token string `json:"token"`
}

