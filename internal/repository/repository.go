package repository

import (
	"database/sql"

	"github.com/Morpa/go-movies-backend/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
}
