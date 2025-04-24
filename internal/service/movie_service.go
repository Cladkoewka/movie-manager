package service

import (
	"github.com/Cladkoewka/movie-manager/internal/constants"
	"github.com/Cladkoewka/movie-manager/internal/model"
	"github.com/Cladkoewka/movie-manager/internal/model/dto"
	"github.com/Cladkoewka/movie-manager/internal/repository"
)

type MovieService struct {
	repo repository.MovieRepository
}

func NewMovieService(repo repository.MovieRepository) *MovieService {
	return &MovieService{repo: repo}
}

func (s *MovieService) GetAllMovies(params dto.MovieQueryParams) (dto.MoviesResponse, error) {
	if !constants.AllowedSortFields[params.SortBy] {
		params.SortBy = constants.DefaultSortBy
	}

	if params.OrderBy != "asc" && params.OrderBy != "desc" {
		params.OrderBy = constants.DefaultOrderBy
	}

	if params.Page <= 0 {
		params.Page = constants.DefaultPage
	}
	if params.PageSize <= 0 {
		params.PageSize = constants.DefaultPageSize
	}

	if params.Rating != nil {
		if *params.Rating < 0 || *params.Rating > 10 {
			params.Rating = nil
		}
	}

	moviesResponse, err := s.repo.GetAllMovies(params)
	if err != nil {
		return dto.MoviesResponse{}, err
	}
	return moviesResponse, nil
}

func (s *MovieService) GetMovieByID(id int64) (*model.Movie, error) {
	movie, err := s.repo.GetMovieByID(id)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (s *MovieService) CreateMovie(movie model.Movie) (*model.Movie, error) {
	newMovie, err := s.repo.CreateMovie(movie)
	if err != nil {
		return nil, err
	}
	return newMovie, nil
}

func (s *MovieService) UpdateMovie(movie model.Movie) (*model.Movie, error) {
	updateMovie, err := s.repo.UpdateMovie(movie)
	if err != nil {
		return nil, err
	}
	return updateMovie, nil
}

func (s *MovieService) DeleteMovie(id int64) error {
	err := s.repo.DeleteMovie(id)
	if err != nil {
		return err
	}
	return nil
}
