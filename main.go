// @title Movie Manager API
// @version 1.0
// @description This is a simple Movie Manager API
// @host localhost:8080
// @BasePath /

// @schemes http
package main

import (
	"github.com/Cladkoewka/movie-manager/internal/repository"
	"github.com/Cladkoewka/movie-manager/internal/service"
	"github.com/Cladkoewka/movie-manager/internal/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/Cladkoewka/movie-manager/internal/model"
	"log"
	_ "github.com/Cladkoewka/movie-manager/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/movies", handler.GetAllMovies)
	r.GET("/movies/:id", handler.GetMovieByID)
	r.POST("/movies", handler.CreateMovie)
	r.PUT("/movies/:id", handler.UpdateMovie)
	r.DELETE("/movies/:id", handler.DeleteMovie)

	if err := r.Run("localhost:8080"); err != nil {
		log.Fatal("Failed to start server", err)
	}
}