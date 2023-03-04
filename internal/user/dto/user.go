package dto

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct{}

type LoginUserRequest struct{}

type LoginUserResponse struct{}
