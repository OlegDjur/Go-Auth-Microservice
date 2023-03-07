package service

import (
	"github.com/OlegDjur/Go-Auth-Microservice/internal/models"
	"github.com/OlegDjur/Go-Auth-Microservice/internal/user"
	"github.com/OlegDjur/Go-Auth-Microservice/internal/user/dto"
)

type service struct {
	repo *user.Repository
}

func NewService(repo *user.Repository) *service {
	return &service{repo}
}

func (s *service) Create(user *dto.CreateUserRequest) (*models.User, error) {

	user, err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}
}
