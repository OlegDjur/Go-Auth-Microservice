package user

import (
	"github.com/OlegDjur/Go-Auth-Microservice/internal/models"
	"github.com/OlegDjur/Go-Auth-Microservice/internal/user/dto"
)

type Service interface {
	Create(*dto.CreateUserRequest) (models.User, error)
	Login(*dto.LoginUserRequest) (models.User, string, error)
}
