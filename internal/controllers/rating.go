package controllers

import (
	"github.com/bingemate/media-service/internal/features"
	"github.com/gin-gonic/gin"
	"strconv"
)

func InitRatingController(engine *gin.RouterGroup, ratingService *features.RatingService) {
	engine.GET("/media/:mediaID", func(c *gin.Context) {
		getMediaRating(c, ratingService)
	})
	engine.GET("/media/:mediaID/own", func(c *gin.Context) {
		getUserMediaRating(c, ratingService)
	})
	engine.GET("/user/:userID", func(c *gin.Context) {
		getUserRating(c, ratingService)
	})
	engine.POST("/media/:mediaID", func(c *gin.Context) {
		saveMediaRating(c, ratingService)
	})
}

/*// @Summary Get media's comments
// @Description Get media's comments
// @Tags Comment
// @Param page query int false "Page number"
// @Param mediaID path int true "Media ID"
// @Produce json
// @Success 200 {object} commentResults
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /comment/media/{mediaID} [get]
func getComments(c *gin.Context, commentService *features.CommentService) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	mediaID, err := strconv.Atoi(c.Param("mediaID"))
	if err != nil {
		c.JSON(400, errorResponse{Error: "mediaID must be a number"})
		return
	}
	if mediaID <= 0 {
		c.JSON(400, errorResponse{Error: "mediaID must be a positive number"})
		return
	}
	comments, total, err := commentService.GetComments(mediaID, page)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, commentResults{
		Results:     toCommentsResponse(comments),
		TotalResult: total,
	})
}*/

// @Summary Get media's rating
// @Description Get media's rating
// @Tags Rating
// @Param mediaID path int true "Media ID"
// @Param page query int false "Page number"
// @Produce json
// @Success 200 {object} ratingResults
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /rating/media/{mediaID} [get]
func getMediaRating(c *gin.Context, ratingService *features.RatingService) {
	mediaID, err := strconv.Atoi(c.Param("mediaID"))
	if err != nil {
		c.JSON(400, errorResponse{Error: "mediaID must be a number"})
		return
	}
	if mediaID <= 0 {
		c.JSON(400, errorResponse{Error: "mediaID must be a positive number"})
		return
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	ratings, count, err := ratingService.GetMediaRating(mediaID, page)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, ratingResults{
		Results:     toRatingsResponse(ratings),
		TotalResult: count,
	})
}

// @Summary Get user's media rating
// @Description Get user's media rating
// @Tags Rating
// @Param mediaID path int true "Media ID"
// @Param user-id header string true "User ID"
// @Produce json
// @Success 200 {object} ratingResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /rating/media/{mediaID}/own [get]
func getUserMediaRating(c *gin.Context, ratingService *features.RatingService) {
	mediaID, err := strconv.Atoi(c.Param("mediaID"))
	if err != nil {
		c.JSON(400, errorResponse{Error: "mediaID must be a number"})
		return
	}
	if mediaID <= 0 {
		c.JSON(400, errorResponse{Error: "mediaID must be a positive number"})
		return
	}
	userID := c.GetHeader("user-id")
	if userID == "" {
		c.JSON(400, errorResponse{Error: "user-id header is required"})
		return
	}

	rating, err := ratingService.GetUserMediaRating(userID, mediaID)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, toRatingResponse(rating))
}

// @Summary Get user's rating
// @Description Get user's rating
// @Tags Rating
// @Param userID path string true "User ID"
// @Param page query int false "Page number"
// @Produce json
// @Success 200 {object} ratingResults
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /rating/user/{userID} [get]
func getUserRating(c *gin.Context, ratingService *features.RatingService) {
	userID := c.Param("userID")
	if userID == "" {
		c.JSON(400, errorResponse{Error: "userID is required"})
		return
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	ratings, count, err := ratingService.GetUsersRating(userID, page)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, ratingResults{
		Results:     toRatingsResponse(ratings),
		TotalResult: count,
	})
}

// @Summary Save media's rating
// @Description Save media's rating
// @Tags Rating
// @Param mediaID path int true "Media ID"
// @Param user-id header string true "User ID"
// @Param rating body ratingRequest true "Rating"
// @Produce json
// @Success 200 {object} ratingResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /rating/media/{mediaID} [post]
func saveMediaRating(c *gin.Context, ratingService *features.RatingService) {
	mediaID, err := strconv.Atoi(c.Param("mediaID"))
	if err != nil {
		c.JSON(400, errorResponse{Error: "mediaID must be a number"})
		return
	}
	if mediaID <= 0 {
		c.JSON(400, errorResponse{Error: "mediaID must be a positive number"})
		return
	}
	userID := c.GetHeader("user-id")
	if userID == "" {
		c.JSON(400, errorResponse{Error: "user-id header is required"})
		return
	}

	var ratingRequest ratingRequest
	if err := c.ShouldBindJSON(&ratingRequest); err != nil {
		c.JSON(400, errorResponse{Error: err.Error()})
		return
	}

	rating, err := ratingService.RateMedia(userID, mediaID, ratingRequest.Rating)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, toRatingResponse(rating))
}
