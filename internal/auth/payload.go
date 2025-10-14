package auth

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=6"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name"`
	Password string `json:"password" validate:"required,gte=6"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}
