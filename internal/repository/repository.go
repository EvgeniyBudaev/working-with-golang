package repository

import (
	"database/sql"
	"github.com/EvgeniyBudaev/working-with-golang/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
}
