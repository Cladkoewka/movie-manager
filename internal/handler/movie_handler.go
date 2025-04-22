package handler

import (
	"net/http"
	"strconv"

	"github.com/Cladkoewka/movie-manager/internal/constants"
	"github.com/Cladkoewka/movie-manager/internal/model"
	"github.com/Cladkoewka/movie-manager/internal/model/dto"
	"github.com/Cladkoewka/movie-manager/internal/service"
	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	service *service.MovieService
}

func NewMovieHandler(service *service.MovieService) *MovieHandler {
	return &MovieHandler{service: service}
}

// GetAllMovies godoc
// @Summary Get all movies
// @Description Get list of all movies
// @Tags movies
// @Accept json
// @Produce json
// @Param search query string false "Search term for movie title"
// @Param genre query string false "Genre of the movie"
// @Param language query string false "Language of the movie"
// @Param rating query float64 false "Rating of the movie (0-10)"
// @Param sort_by query string false "Field to sort by (e.g. 'title', 'rating')"
// @Param order query string false "Order of sorting ('asc' or 'desc')"
// @Param page query int false "Page number for pagination"
// @Param pageSize query int false "Number of items per page"
// @Success 200 {array} model.Movie
// @Failure 500 {object} map[string]interface{}
// @Router /movies [get]
func (h *MovieHandler) GetAllMovies(c *gin.Context) {
	var params dto.MovieQueryParams
	params.Search = c.DefaultQuery("search", "")
	params.Genre = c.DefaultQuery("genre", "")
	params.Language = c.DefaultQuery("language", "")

	if ratingStr := c.Query("rating"); ratingStr != "" {
		rating, err := strconv.ParseFloat(ratingStr, 64)
		if err == nil {
			params.Rating = &rating
		}
	}

	params.SortBy = c.DefaultQuery("sort_by", constants.DefaultSortBy)
	params.OrderBy = c.DefaultQuery("order", constants.DefaultOrderBy)

	if page, err := strconv.Atoi(c.DefaultQuery("page", strconv.Itoa(constants.DefaultPage))); err == nil && page > 0 {
		params.Page = page
	} else {
		params.Page = constants.DefaultPage
	}

	if size, err := strconv.Atoi(c.DefaultQuery("pageSize", strconv.Itoa(constants.DefaultPageSize))); err == nil && size > 0 {
		params.PageSize = size
	} else {
		params.PageSize = constants.DefaultPageSize
	}

	movies, err := h.service.GetAllMovies(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies"})
		return 
	}
	c.JSON(http.StatusOK, movies)
}

// GetMovieByID godoc
// @Summary Get a movie by ID
// @Description Get a movie details by movie ID
// @Tags movies
// @Accept json
// @Produce json
// @Param id path int64 true "Movie ID"
// @Success 200 {object} model.Movie
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /movies/{id} [get]
func (h *MovieHandler) GetMovieByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}
	movie, err := h.service.GetMovieByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}
	c.JSON(http.StatusOK, movie)
}

// CreateMovie godoc
// @Summary Create a new movie
// @Description Add a new movie to the database
// @Tags movies
// @Accept json
// @Produce json
// @Param movie body model.Movie true "Movie details"
// @Success 201 {object} model.Movie
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /movies [post]
func (h *MovieHandler) CreateMovie(c *gin.Context) {
	var movie model.Movie
	if err := c.ShouldBindJSON(&movie); err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	newMovie, err := h.service.CreateMovie(movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create movie"})
		return
	}
	c.JSON(http.StatusCreated, newMovie)
}

// UpdateMovie godoc
// @Summary Update an existing movie
// @Description Update the details of an existing movie
// @Tags movies
// @Accept json
// @Produce json
// @Param id path int64 true "Movie ID"
// @Param movie body model.Movie true "Movie details"
// @Success 200 {object} model.Movie
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /movies/{id} [put]
func (h *MovieHandler) UpdateMovie(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}
	var movie model.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	movie.ID = id
	updatedMovie, err := h.service.UpdateMovie(movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update movie"})
		return
	}
	c.JSON(http.StatusOK, updatedMovie)
}

// DeleteMovie godoc
// @Summary Delete a movie by ID
// @Description Remove a movie from the database by its ID
// @Tags movies
// @Accept json
// @Produce json
// @Param id path int64 true "Movie ID"
// @Success 204 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /movies/{id} [delete]
func (h *MovieHandler) DeleteMovie(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}
	err = h.service.DeleteMovie(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete movie"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
