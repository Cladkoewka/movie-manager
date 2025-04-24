package repository

import (
	"gorm.io/gorm"
	"github.com/Cladkoewka/movie-manager/internal/model"
)

type ReviewRepository interface {
	GetAllByMovieID(movieID int64) ([]model.Review, error)
	Create(review model.Review) (*model.Review, error)
	Delete(reviewID int64) error
}

type ReviewRepositoryImpl struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &ReviewRepositoryImpl{db: db}
}

func (r *ReviewRepositoryImpl) GetAllByMovieID(movieID int64) ([]model.Review, error) {
	var reviews []model.Review
	err := r.db.Where("movie_id = ?", movieID).Find(&reviews).Error
	return reviews, err
}

func (r *ReviewRepositoryImpl) Create(review model.Review) (*model.Review, error) {
	if err := r.db.Create(&review).Error; err != nil {
		return nil, err
	}
	return &review, nil
}

func (r *ReviewRepositoryImpl) Delete(reviewID int64) error {
	return r.db.Delete(&model.Review{}, reviewID).Error
}
