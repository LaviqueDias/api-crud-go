package request


type UserRequest struct {
    Name      string    `json:"name" binding:"required,min=4,max=100"`
    Email     string    `json:"email" binding:"required,email"`
    Password  string    `json:"password" binding:"required,min=6,max=50,containsAny=!?*#@&$"`
    Role      string    `json:"role" binding:"required,oneof=admin user"`
    Active    bool      `json:"active" binding:"required"`
}
