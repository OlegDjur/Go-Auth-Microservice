package service

import (
	"context"

	"github.com/OlegDjur/Go-Auth-Microservice/internal/models"
	"github.com/OlegDjur/Go-Auth-Microservice/internal/user"
	"github.com/OlegDjur/Go-Auth-Microservice/internal/user/dto"
	"github.com/OlegDjur/Go-Auth-Microservice/pkg/utils"
)

type service struct {
	repo *user.Repository
}

func NewService(repo *user.Repository) *service {
	return &service{repo}
}

func (s *service) Create(cxt context.Context, req *dto.CreateUserRequest) (*models.User, error) {
	err := utils.ValidateUser(req)
	if err != nil {
		return nil, err
	}

	req.Password, err = utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return user, nil
}
