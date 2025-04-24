package handler

import (
	"net/http"
	"strconv"

	"github.com/Cladkoewka/movie-manager/internal/model"
	"github.com/Cladkoewka/movie-manager/internal/service"
	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	reviewService *service.ReviewService
}

func NewReviewHandler(service *service.ReviewService) *ReviewHandler {
	return &ReviewHandler{reviewService: service}
}

// GetReviewsByMovieID godoc
// @Summary Get all reviews for a movie
// @Description Get all reviews by Movie ID
// @Tags reviews
// @Param movie_id path int true "Movie ID"
// @Produce json
// @Success 200 {array} model.Review
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /reviews/movie/{movie_id} [get]
func (h *ReviewHandler) GetReviewsByMovieID(c *gin.Context) {
	movieID, err := strconv.ParseInt(c.Param("movie_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}
	reviews, err := h.reviewService.GetAllByMovieID(movieID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reviews"})
		return
	}
	c.JSON(http.StatusOK, reviews)
}

// CreateReview godoc
// @Summary Create a review
// @Description Create a new review for a movie
// @Tags reviews
// @Accept json
// @Produce json
// @Param review body model.Review true "Review payload"
// @Success 201 {object} model.Review
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /reviews [post]
func (h *ReviewHandler) CreateReview(c *gin.Context) {
	var review model.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	created, err := h.reviewService.CreateReview(review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create review"})
		return
	}
	c.JSON(http.StatusCreated, created)
}

// DeleteReview godoc
// @Summary Delete a review
// @Description Delete a review by ID
// @Tags reviews
// @Param id path int true "Review ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /reviews/{id} [delete]
func (h *ReviewHandler) DeleteReview(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		return
	}
	if err := h.reviewService.DeleteReview(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete review"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
