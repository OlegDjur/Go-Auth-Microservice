package repository

import (
	"context"
	"database/sql"

	"github.com/OlegDjur/Go-Auth-Microservice/internal/user/dto"
)

type Repository struct {
	db *sql.DB
}

func NewRopository(db *sql.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Create(ctx context.Context, req dto.CreateUserRequest) (*dto.UserResponse, error) {
	var user *dto.UserResponse

	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING username, email, created_at`

	row := r.db.QueryRowContext(ctx, query, req.Username, req.Email, req.Password)

	err := row.Scan(&user.Username, &user.Email, user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}
