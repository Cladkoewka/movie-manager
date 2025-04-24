// @title Movie Manager API
// @version 1.0
// @description This is a simple Movie Manager API
// @host localhost:8080
// @BasePath /
// @schemes http
package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/gin-contrib/cors"

	_ "github.com/Cladkoewka/movie-manager/docs"
	"github.com/Cladkoewka/movie-manager/internal/cache"
	"github.com/Cladkoewka/movie-manager/internal/handler"
	"github.com/Cladkoewka/movie-manager/internal/loader"
	"github.com/Cladkoewka/movie-manager/internal/model"
	"github.com/Cladkoewka/movie-manager/internal/repository"
	"github.com/Cladkoewka/movie-manager/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/kurin/blazer/b2"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

var (
	shouldMigrate bool
	shouldLoadInitialData bool
)

func init() {
	flag.BoolVar(&shouldMigrate, "migrate", false, "Run database migrations")
	flag.BoolVar(&shouldLoadInitialData, "load", false, "Load initial data from JSON")
	flag.Parse()
}


func main() {
	db := initDB()

	//if shouldMigrate {
		runMigrations(db)
	//}

	//bucket, bucketURL := initB2()
	
	cacheService := cache.NewRedisService()

	reviewRepository := repository.NewReviewRepository(db)
	reviewService := service.NewReviewService(reviewRepository)
	reviewHandler := handler.NewReviewHandler(reviewService)
	movieRepository := repository.NewMovieRepository(db, cacheService)
	movieService := service.NewMovieService(movieRepository)
	moviePosterRepository := repository.NewMoviePosterRepository(db)
	moviePosterService := service.NewMoviePosterService(moviePosterRepository)
	movieHandler := handler.NewMovieHandler(movieService, moviePosterService)
	//movieTrailerService := service.NewMovieTrailerService(movieRepository, bucket, bucketURL)
	//movieTrailerHandler := handler.NewMovieTrailerHandler(movieTrailerService)

	//if shouldLoadInitialData {
		loadInitialData(movieService, reviewService)
	//}

	r := gin.Default()
	
	r.Use(cors.Default())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/movies", movieHandler.GetAllMovies)
	r.GET("/movies/:id", movieHandler.GetMovieByID)
	r.POST("/movies", movieHandler.CreateMovie)
	r.PUT("/movies/:id", movieHandler.UpdateMovie)
	r.DELETE("/movies/:id", movieHandler.DeleteMovie)
	r.POST("/movies/:id/poster", movieHandler.UploadPoster)
	r.GET("/movies/:id/poster", movieHandler.GetPoster)
	//r.POST("/movies/:id/trailer", movieTrailerHandler.UploadTrailer)
	//r.PUT("/movies/:id/trailer", movieTrailerHandler.SetTrailerUrl)
	r.GET("/reviews/movie/:movie_id", reviewHandler.GetReviewsByMovieID)
	r.POST("/reviews", reviewHandler.CreateReview)
	r.DELETE("/reviews/:id", reviewHandler.DeleteReview)


	startServer(r)
}

func initDB() *gorm.DB {
	db, err := repository.NewDBConnection()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	return db
}

func runMigrations(db *gorm.DB) {
	if err := db.AutoMigrate(&model.Movie{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	if err := db.AutoMigrate(&model.MoviePoster{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	if err := db.AutoMigrate(&model.Review{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}

func initB2() (*b2.Bucket, string) {
	keyID := os.Getenv("B2_KEY_ID")
	appKey := os.Getenv("B2_APP_KEY")
	bucketName := os.Getenv("B2_BUCKET")
	bucketURL := os.Getenv("B2_BUCKET_URL")

	client, err := b2.NewClient(context.Background(), keyID, appKey)
	if err != nil {
		log.Fatalf("Failed to create B2 client: %v", err)
	}

	bucket, err := client.Bucket(context.Background(), bucketName)
	if err != nil {
		log.Fatal("Failed to get B2 bucket:", err)
	}

	return bucket, bucketURL
}

func loadInitialData(movieService *service.MovieService, reviewService *service.ReviewService) {
	err := loader.LoadMoviesFromJSON(movieService, "movies_dump.json")
	if err != nil {
	log.Fatal("Failed to load movies from JSON:", err)
	}
	if err := loader.LoadReviewsFromJSON(reviewService, "reviews_dump.json"); err != nil {
		log.Fatal("Failed to load reviews from JSON:", err)
	}
}

func startServer(r *gin.Engine) {
	port := os.Getenv("PORT")
	if port == "" {
			port = "8080" // fallback для локального запуска
	}
	
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
