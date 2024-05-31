package web

// membuat struct request untuk user register
type UserRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
	Name     string `validate:"required" json:"name"`
}

// membuat struct request untuk user login
type UserLoginRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

type TokenRequest struct {
	Token string `validate:"required" json:"token"`
}
