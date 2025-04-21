package service

import (
	"github.com/Cladkoewka/movie-manager/internal/model"
	"github.com/Cladkoewka/movie-manager/internal/repository"
)

type MovieService struct {
	repo repository.MovieRepository
}

func NewMovieService(repo repository.MovieRepository) *MovieService {
	return &MovieService{repo: repo}
}

func (s *MovieService) GetAllMovies() ([]model.Movie, error) {
	movies, err := s.repo.GetAllMovies()
	if err != nil {
		return nil, err
	}
	return movies, nil
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