package controllers

import (
	"github.com/bingemate/media-service/internal/features"
	"github.com/gin-gonic/gin"
	"strconv"
)

func InitRatingController(engine *gin.RouterGroup, ratingService *features.RatingService) {
	engine.GET("/movie/:mediaID", func(c *gin.Context) {
		getMovieRating(c, ratingService)
	})
	engine.GET("/movie/:mediaID/own", func(c *gin.Context) {
		getUserMovieRating(c, ratingService)
	})
	engine.GET("/tv/:mediaID", func(c *gin.Context) {
		getTVShowRating(c, ratingService)
	})
	engine.GET("/tv/:mediaID/own", func(c *gin.Context) {
		getUserTVShowRating(c, ratingService)
	})
	engine.GET("/movie/user/:userID", func(c *gin.Context) {
		getUserMovieRatings(c, ratingService)
	})
	engine.POST("/movie/:mediaID", func(c *gin.Context) {
		saveMovieRating(c, ratingService)
	})
	engine.GET("/tv/user/:userID", func(c *gin.Context) {
		getUserTVShowRatings(c, ratingService)
	})
	engine.POST("/tv/:mediaID", func(c *gin.Context) {
		saveTVShowRating(c, ratingService)
	})
	engine.GET("/user/count/:userID", func(c *gin.Context) {
		getUserRatingCount(c, ratingService)
	})
	engine.GET("/count", func(c *gin.Context) {
		getRatingCount(c, ratingService)
	})
}

// @Summary Get movie's rating
// @Description Get movie's rating
// @Tags Rating
// @Param mediaID path int true "Movie ID"
// @Param page query int false "Page number"
// @Produce json
// @Success 200 {object} ratingResults
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /rating/movie/{mediaID} [get]
func getMovieRating(c *gin.Context, ratingService *features.RatingService) {
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

	ratings, count, err := ratingService.GetMovieRatings(mediaID, page)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, ratingResults{
		Results:     toMovieRatingsResponse(ratings),
		TotalResult: count,
	})
}

// @Summary Get tv show's rating
// @Description Get tv show's rating
// @Tags Rating
// @Param mediaID path int true "TV Show ID"
// @Param page query int false "Page number"
// @Produce json
// @Success 200 {object} ratingResults
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /rating/tv/{mediaID} [get]
func getTVShowRating(c *gin.Context, ratingService *features.RatingService) {
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

	ratings, count, err := ratingService.GetTvShowRatings(mediaID, page)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, ratingResults{
		Results:     toTVShowRatingsResponse(ratings),
		TotalResult: count,
	})
}

// @Summary Get user's movie rating
// @Description Get user's movie rating
// @Tags Rating
// @Param mediaID path int true "Movie ID"
// @Param user-id header string true "User ID"
// @Produce json
// @Success 200 {object} ratingResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /rating/movie/{mediaID}/own [get]
func getUserMovieRating(c *gin.Context, ratingService *features.RatingService) {
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

	rating, err := ratingService.GetUserMovieRating(userID, mediaID)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, toMovieRatingResponse(rating))
}

// @Summary Get user's tv show rating
// @Description Get user's tv show rating
// @Tags Rating
// @Param mediaID path int true "TV Show ID"
// @Param user-id header string true "User ID"
// @Produce json
// @Success 200 {object} ratingResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /rating/tv/{mediaID}/own [get]
func getUserTVShowRating(c *gin.Context, ratingService *features.RatingService) {
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

	rating, err := ratingService.GetUserTvShowRating(userID, mediaID)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, toTVShowRatingResponse(rating))
}

// @Summary Get user's movie ratings
// @Description Get user's ratings
// @Tags Rating
// @Param userID path string true "User ID"
// @Param page query int false "Page number"
// @Produce json
// @Success 200 {object} ratingResults
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /rating/movie/user/{userID} [get]
func getUserMovieRatings(c *gin.Context, ratingService *features.RatingService) {
	userID := c.Param("userID")
	if userID == "" {
		c.JSON(400, errorResponse{Error: "userID is required"})
		return
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	ratings, count, err := ratingService.GetUsersMovieRatings(userID, page)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, ratingResults{
		Results:     toMovieRatingsResponse(ratings),
		TotalResult: count,
	})
}

// @Summary Get user's tv show ratings
// @Description Get user's ratings
// @Tags Rating
// @Param userID path string true "User ID"
// @Param page query int false "Page number"
// @Produce json
// @Success 200 {object} ratingResults
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /rating/tv/user/{userID} [get]
func getUserTVShowRatings(c *gin.Context, ratingService *features.RatingService) {
	userID := c.Param("userID")
	if userID == "" {
		c.JSON(400, errorResponse{Error: "userID is required"})
		return
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	ratings, count, err := ratingService.GetUsersTvShowRatings(userID, page)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, ratingResults{
		Results:     toTVShowRatingsResponse(ratings),
		TotalResult: count,
	})
}

// @Summary Save movie's rating
// @Description Save movie's rating
// @Tags Rating
// @Param mediaID path int true "Movie ID"
// @Param user-id header string true "User ID"
// @Param rating body ratingRequest true "Rating"
// @Produce json
// @Success 200 {object} ratingResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /rating/movie/{mediaID} [post]
func saveMovieRating(c *gin.Context, ratingService *features.RatingService) {
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

	rating, err := ratingService.RateMovie(userID, mediaID, ratingRequest.Rating)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, toMovieRatingResponse(rating))
}

// @Summary Save tv show's rating
// @Description Save tv show's rating
// @Tags Rating
// @Param mediaID path int true "TV Show ID"
// @Param user-id header string true "User ID"
// @Param rating body ratingRequest true "Rating"
// @Produce json
// @Success 200 {object} ratingResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /rating/tv/{mediaID} [post]
func saveTVShowRating(c *gin.Context, ratingService *features.RatingService) {
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

	rating, err := ratingService.RateTvShow(userID, mediaID, ratingRequest.Rating)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, toTVShowRatingResponse(rating))
}

// @Summary Get User's rating count
// @Description Get User's rating count
// @Tags Rating
// @Param userID path string true "User ID"
// @Produce json
// @Success 200 {object} int
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /rating/user/count/{userID} [get]
func getUserRatingCount(c *gin.Context, ratingService *features.RatingService) {
	userID := c.Param("userID")
	if userID == "" {
		c.JSON(400, errorResponse{Error: "userID is required"})
		return
	}

	count, err := ratingService.CountUserRatings(userID)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, count)
}

// @Summary Get rating count
// @Description Get rating count
// @Tags Rating
// @Produce json
// @Success 200 {object} int
// @Failure 500 {object} errorResponse
// @Router /rating/count [get]
func getRatingCount(c *gin.Context, ratingService *features.RatingService) {
	count, err := ratingService.CountRatings()
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, count)
}
