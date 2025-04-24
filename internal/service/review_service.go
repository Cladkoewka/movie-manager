package service

import (
	"github.com/Cladkoewka/movie-manager/internal/model"
	"github.com/Cladkoewka/movie-manager/internal/repository"
)

type ReviewService struct {
	repo repository.ReviewRepository
}

func NewReviewService(repo repository.ReviewRepository) *ReviewService {
	return &ReviewService{repo: repo}
}

func (s *ReviewService) GetAllByMovieID(movieID int64) ([]model.Review, error) {
	return s.repo.GetAllByMovieID(movieID)
}

func (s *ReviewService) CreateReview(review model.Review) (*model.Review, error) {
	return s.repo.Create(review)
}

func (s *ReviewService) DeleteReview(id int64) error {
	return s.repo.Delete(id)
}
