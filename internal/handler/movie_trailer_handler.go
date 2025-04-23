package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"github.com/Cladkoewka/movie-manager/internal/service"
)

type MovieTrailerHandler struct {
	movieTrailerService *service.MovieTrailerService
}

func NewMovieTrailerHandler(movieTrailerService *service.MovieTrailerService) *MovieTrailerHandler {
	return &MovieTrailerHandler{movieTrailerService: movieTrailerService}
}

// UploadTrailer godoc
// @Summary Upload movie trailer
// @Description Uploads a trailer file for a specific movie and stores it in B2
// @Tags Movies
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "Movie ID"
// @Param trailer formData file true "Trailer file"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /movies/{id}/trailer [post]
func (h *MovieTrailerHandler) UploadTrailer(c *gin.Context) {
	movieID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	file, err := c.FormFile("trailer")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload file"})
		return
	}

	err = h.movieTrailerService.UploadTrailer(movieID, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload trailer"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Trailer uploaded successfully"})
}

// SetTrailerURL godoc
// @Summary Set movie trailer URL
// @Description Sets a new trailer URL for a specific movie
// @Tags Movies
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Param url query string true "Trailer URL" 
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /movies/{id}/trailer [put]
func (h *MovieTrailerHandler) SetTrailerUrl(c *gin.Context) {
	movieID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}
	
	url := c.Query("url")
	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing URL query parameter"})
		return
	}

	err = h.movieTrailerService.SetTrailerURL(movieID, url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set trailer URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Trailer URL set successfully"})
}