package repository

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Cladkoewka/movie-manager/internal/cache"
	"github.com/Cladkoewka/movie-manager/internal/config"
	"github.com/Cladkoewka/movie-manager/internal/model"
	"github.com/Cladkoewka/movie-manager/internal/model/dto"
)

type MovieRepository interface {
	GetAllMovies(params dto.MovieQueryParams) (dto.MoviesResponse, error)
	GetMovieByID(id int64) (*model.Movie, error)
	CreateMovie(movie model.Movie) (*model.Movie, error)
	UpdateMovie(movie model.Movie) (*model.Movie, error)
	DeleteMovie(id int64) error
	UpdateMovieTrailer(movieID int64, trailerURL string) error
}

type MovieRepositoryImpl struct {
	db *gorm.DB 
	redisService *cache.RedisService 
}

func NewMovieRepository(db *gorm.DB, redisService *cache.RedisService) MovieRepository {
	return &MovieRepositoryImpl{db: db, redisService: redisService}
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

func (r *MovieRepositoryImpl) GetAllMovies(params dto.MovieQueryParams) (dto.MoviesResponse, error) {
	var movies []model.Movie
	var total int64

	// Caching
	// key, err := r.redisService.GenerateCacheKey("movies", params)
	// if err == nil {
	// 	if err := r.redisService.GetCache(context.Background(), key, &movies); err == nil {
	// 		return dto.MoviesResponse{
	// 			Movies: movies,
	// 			Total:  int64(len(movies)), 
	// 		}, nil
	// 	}
	// }

	query := r.db.Model(&model.Movie{})

	if params.Search != "" {
		query = query.Where("title ILIKE ?", "%" + params.Search + "%")
	}

	if params.Genre != "" {
		query = query.Where("genre ILIKE ?", "%" + params.Genre + "%")
	}

	if params.Language != "" {
		query = query.Where("language ILIKE ?", "%" + params.Language + "%")
	}

	if params.Rating != nil {
		query = query.Where("rating >= ?", *params.Rating)
	}

	if err := query.Count(&total).Error; err != nil {
		return dto.MoviesResponse{}, err
	}

	query = query.Order(params.SortBy + " " + params.OrderBy)

	offset := (params.Page - 1) * params.PageSize
	query = query.Limit(params.PageSize).Offset(offset)

	if err := query.Find(&movies).Error; err != nil {
		return dto.MoviesResponse{}, err
	}

	// if cacheKey, err := r.redisService.GenerateCacheKey("movies", params); err == nil {
  //   _ = r.redisService.SetCache(context.Background(), cacheKey, movies, 10*time.Minute)
	// }	

	return dto.MoviesResponse{
		Movies: movies,
		Total:  total,
	}, nil
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

func (r *MovieRepositoryImpl) UpdateMovieTrailer(movieID int64, trailerURL string) error {
	var movie model.Movie
	if err := r.db.First(&movie, "id = ?", movieID).Error; err != nil {
		return err
	}

	movie.TrailerURL = trailerURL
	if err := r.db.Save(&movie).Error; err != nil {
		return err
	}

	return nil
}
