package dto

import (
	"time"

	"github.com/OlegDjur/Go-Auth-Microservice/internal/models"
)

type CreateUserRequest struct {
	Username string `josn:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginUserRequest struct {
	Username string `josn:"username"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        UserResponse `josn:"user"`
}

func NewUserResponse(user models.User) *UserResponse {
	return &UserResponse{
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}
