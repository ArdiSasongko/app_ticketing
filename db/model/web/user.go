package web

type UserRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
	Name     string `validate:"required" json:"name"`
}
