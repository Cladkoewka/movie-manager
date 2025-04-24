package dto

import "github.com/Cladkoewka/movie-manager/internal/model"

type MoviesResponse struct {
	Movies []model.Movie `json:"movies"`
	Total  int64           `json:"total"`
}