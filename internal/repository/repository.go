package repository

import "github.com/EvgeniyBudaev/working-with-golang/internal/models"

type DatabaseRepo interface {
	AllMovies() ([]*models.Movie, error)
}
