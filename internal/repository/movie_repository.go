package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/Cladkoewka/movie-manager/internal/model"
	"github.com/Cladkoewka/movie-manager/internal/model/dto"
	"github.com/Cladkoewka/movie-manager/internal/config"
)

type MovieRepository interface {
	GetAllMovies(params dto.MovieQueryParams) ([]model.Movie, error)
	GetMovieByID(id int64) (*model.Movie, error)
	CreateMovie(movie model.Movie) (*model.Movie, error)
	UpdateMovie(movie model.Movie) (*model.Movie,error)
	DeleteMovie(id int64) error
}

type MovieRepositoryImpl struct {
	db *gorm.DB 
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return &MovieRepositoryImpl{db: db}
}

func NewDBConnection() (*gorm.DB, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (r *MovieRepositoryImpl) GetAllMovies(params dto.MovieQueryParams) ([]model.Movie, error) {
	var movies []model.Movie
	query := r.db.Model(&model.Movie{})

	if params.Search != "" {
		query = query.Where("title LIKE ?", "%"+params.Search+"%")
	}

	if params.Genre != "" {
		query = query.Where("genre = ?", params.Genre)
	}

	if params.Language != "" {
		query = query.Where("language = ?", params.Language)
	}

	if params.Rating != nil {
		query = query.Where("rating >= ?", *params.Rating)
	}

	query = query.Order(params.SortBy + " " + params.OrderBy)

	offset := (params.Page - 1) * params.PageSize
	query = query.Limit(params.PageSize).Offset(offset)

	if err := query.Find(&movies).Error; err != nil {
		return nil, err
	}

	return movies, nil
}

func (r *MovieRepositoryImpl) GetMovieByID(id int64) (*model.Movie, error) {
	var movie model.Movie
	if err := r.db.First(&movie, id).Error; err != nil {
		return nil, err
	}
	return &movie, nil
}

func (r *MovieRepositoryImpl) CreateMovie(movie model.Movie) (*model.Movie, error) {
	if err := r.db.Create(&movie).Error; err != nil {
		return nil, err
	}
	return &movie, nil
}

func (r *MovieRepositoryImpl) UpdateMovie(movie model.Movie) (*model.Movie, error) {
	if err := r.db.Save(&movie).Error; err != nil {
		return nil, err
	}
	return &movie, nil	
}

func (r *MovieRepositoryImpl) DeleteMovie(id int64) error {
	if err := r.db.Delete(&model.Movie{}, id).Error; err != nil {
		return err
	}
	return nil
}
