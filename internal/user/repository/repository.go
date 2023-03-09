package repository

import (
	"database/sql"

	"github.com/OlegDjur/Go-Auth-Microservice/internal/models"
	"github.com/OlegDjur/Go-Auth-Microservice/internal/user/dto"
)

type Repository struct {
	db *sql.DB
}

func NewRopository(db *sql.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Create(req dto.CreateUserRequest) (*models.User, error) {
	var user *models.User

	return user, nil
}
