package repository

import (
	"time"

	"github.com/Cladkoewka/movie-manager/internal/model"
	"gorm.io/gorm"
)

type MoviePosterRepository interface {
	SavePoster(movieID int64, poster []byte, mimeType string) error
	GetPosterByMovieID(movieID int64) (*model.MoviePoster, error)
	DeletePoster(movieID int64) error
}

type MoviePosterRepositoryImpl struct {
	db *gorm.DB
}

func NewMoviePosterRepository(db *gorm.DB) *MoviePosterRepositoryImpl {
	return &MoviePosterRepositoryImpl{db: db}
}

func (r *MoviePosterRepositoryImpl) SavePoster(movieID int64, poster []byte, mimeType string) error {
	posterRecord := &model.MoviePoster{
		Poster:   poster,
		MovieID: movieID,
		MimeType: mimeType,
		CreatedAt: time.Now(),
	}
	return r.db.Create(posterRecord).Error
}

func (r *MoviePosterRepositoryImpl) GetPosterByMovieID(movieID int64) (*model.MoviePoster, error) {
	var poster model.MoviePoster
	err := r.db.Where("movie_id = ?", movieID).First(&poster).Error
	if err != nil {
		return nil, err
	}
	return &poster, nil
}

func (r *MoviePosterRepositoryImpl) DeletePoster(movieID int64) error {
	return r.db.Where("movie_id = ?", movieID).Delete(&model.MoviePoster{}).Error
}
