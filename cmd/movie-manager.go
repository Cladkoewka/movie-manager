package main

import (
	"github.com/Cladkoewka/movie-manager/internal/repository"
	"github.com/Cladkoewka/movie-manager/internal/service"
	"github.com/Cladkoewka/movie-manager/internal/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/Cladkoewka/movie-manager/internal/model"
	"log"
)

func runMigrations(db *gorm.DB) {
	if err := db.AutoMigrate(&model.Movie{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}

func main() {
	db, err := repository.NewDBConnection()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	repository := repository.NewMovieRepository(db)
	service := service.NewMovieService(repository)
	handler := handler.NewMovieHandler(service)

	r := gin.Default()

	r.GET("/movies", handler.GetAllMovies)
	r.GET("/movies/:id", handler.GetMovieByID)
	r.POST("/movies", handler.CreateMovie)
	r.PUT("/movies/:id", handler.UpdateMovie)
	r.DELETE("/movies/:id", handler.DeleteMovie)

	if err := r.Run("localhost:8080"); err != nil {
		log.Fatal("Failed to start server", err)
	}
}