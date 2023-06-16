package controllers

import (
	"github.com/bingemate/media-service/internal/features"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

func InitCommentController(engine *gin.RouterGroup, commentService *features.CommentService) {
	engine.GET("movie/:mediaID", func(c *gin.Context) {
		getMovieComments(c, commentService)
	})
	engine.GET("movie/user/:userID", func(c *gin.Context) {
		getUserMovieComments(c, commentService)
	})
	engine.POST("movie/:mediaID", func(c *gin.Context) {
		addMovieComment(c, commentService)
	})
	engine.DELETE("movie/:commentID", func(c *gin.Context) {
		deleteMovieComment(c, commentService)
	})
	engine.PUT("movie/:commentID", func(c *gin.Context) {
		updateMovieComment(c, commentService)
	})
	engine.GET("tv/:mediaID", func(c *gin.Context) {
		getTVShowComments(c, commentService)
	})
	engine.GET("tv/user/:userID", func(c *gin.Context) {
		getUserTVShowComments(c, commentService)
	})
	engine.POST("tv/:mediaID", func(c *gin.Context) {
		addTVShowComment(c, commentService)
	})
	engine.DELETE("tv/:commentID", func(c *gin.Context) {
		deleteTVShowComment(c, commentService)
	})
	engine.PUT("tv/:commentID", func(c *gin.Context) {
		updateTVShowComment(c, commentService)
	})
	engine.GET("user/history/:userID", func(c *gin.Context) {
		getUserCommentHistory(c, commentService)
	})
	engine.GET("user/count/:userID", func(c *gin.Context) {
		getUserCommentCount(c, commentService)
	})
	engine.GET("/history", func(c *gin.Context) {
		getCommentHistory(c, commentService)
	})
	engine.GET("/count", func(c *gin.Context) {
		getCommentCount(c, commentService)
	})
}

// @Summary Get movie's comments
// @Description Get movie's comments
// @Tags Comment
// @Param page query int false "Page number"
// @Param mediaID path int true "Movie ID"
// @Produce json
// @Success 200 {object} commentResults
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /comment/movie/{mediaID} [get]
func getMovieComments(c *gin.Context, commentService *features.CommentService) {
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
	comments, total, err := commentService.GetMovieComments(mediaID, page)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, commentResults{
		Results:     toMovieCommentsResponse(comments),
		TotalResult: total,
	})
}

// @Summary Get tv show's comments
// @Description Get tv show's comments
// @Tags Comment
// @Param page query int false "Page number"
// @Param mediaID path int true "TV Show ID"
// @Produce json
// @Success 200 {object} commentResults
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /comment/tv/{mediaID} [get]
func getTVShowComments(c *gin.Context, commentService *features.CommentService) {
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
	comments, total, err := commentService.GetTvShowComments(mediaID, page)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, commentResults{
		Results:     toTVShowCommentsResponse(comments),
		TotalResult: total,
	})
}

// @Summary Get user's movie comments
// @Description Get user's movie comments
// @Tags Comment
// @Param page query int false "Page number"
// @Param userID path string true "User ID"
// @Produce json
// @Success 200 {object} commentResults
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /comment/movie/user/{userID} [get]
func getUserMovieComments(c *gin.Context, commentService *features.CommentService) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	userID := c.Param("userID")
	if userID == "" {
		c.JSON(400, errorResponse{Error: "userID must be a string"})
		return
	}
	comments, total, err := commentService.GetMovieUserComments(userID, page)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, commentResults{
		Results:     toMovieCommentsResponse(comments),
		TotalResult: total,
	})
}

// @Summary Get user's tv show comments
// @Description Get user's tv show comments
// @Tags Comment
// @Param page query int false "Page number"
// @Param userID path string true "User ID"
// @Produce json
// @Success 200 {object} commentResults
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /comment/tv/user/{userID} [get]
func getUserTVShowComments(c *gin.Context, commentService *features.CommentService) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	userID := c.Param("userID")
	if userID == "" {
		c.JSON(400, errorResponse{Error: "userID must be a string"})
		return
	}
	comments, total, err := commentService.GetTvShowUserComments(userID, page)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, commentResults{
		Results:     toTVShowCommentsResponse(comments),
		TotalResult: total,
	})
}

// @Summary Add movie comment
// @Description Add comment to a movie
// @Tags Comment
// @Param mediaID path int true "Movie ID"
// @Param comment body commentRequest true "Comment"
// @Param user-id header string true "User ID"
// @Produce json
// @Success 200 {object} commentResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /comment/movie/{mediaID} [post]
func addMovieComment(c *gin.Context, commentService *features.CommentService) {
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
		c.JSON(400, errorResponse{Error: "user-id must not be empty"})
		return
	}
	var comment commentRequest
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(400, errorResponse{Error: err.Error()})
		return
	}
	if comment.Content == "" || strings.TrimSpace(comment.Content) == "" {
		c.JSON(400, errorResponse{Error: "comment must not be empty"})
		return
	}
	if len(comment.Content) > 1000 {
		c.JSON(400, errorResponse{Error: "comment must not be longer than 1000 characters"})
		return
	}
	commentResult, err := commentService.AddMovieComment(userID, mediaID, comment.Content)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, toMovieCommentResponse(commentResult))
}

// @Summary Add tv show comment
// @Description Add comment to a tv show
// @Tags Comment
// @Param mediaID path int true "TV Show ID"
// @Param comment body commentRequest true "Comment"
// @Param user-id header string true "User ID"
// @Produce json
// @Success 200 {object} commentResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /comment/tv/{mediaID} [post]
func addTVShowComment(c *gin.Context, commentService *features.CommentService) {
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
		c.JSON(400, errorResponse{Error: "user-id must not be empty"})
		return
	}
	var comment commentRequest
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(400, errorResponse{Error: err.Error()})
		return
	}
	if comment.Content == "" || strings.TrimSpace(comment.Content) == "" {
		c.JSON(400, errorResponse{Error: "comment must not be empty"})
		return
	}
	if len(comment.Content) > 1000 {
		c.JSON(400, errorResponse{Error: "comment must not be longer than 1000 characters"})
		return
	}
	commentResult, err := commentService.AddTvShowComment(userID, mediaID, comment.Content)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, toTVShowCommentResponse(commentResult))
}

// @Summary Delete movie comment
// @Description Delete movie comment
// @Tags Comment
// @Param commentID path string true "Movie Comment ID"
// @Param user-id header string true "User ID"
// @Param roles header string false "User roles"
// @Produce json
// @Success 204 {string} string "comment deleted"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /comment/movie/{commentID} [delete]
func deleteMovieComment(c *gin.Context, commentService *features.CommentService) {
	commentID := c.Param("commentID")
	if commentID == "" {
		c.JSON(400, errorResponse{Error: "commentID is required"})
		return
	}
	userID := c.GetHeader("user-id")
	if userID == "" {
		c.JSON(400, errorResponse{Error: "user-id must not be empty"})
		return
	}
	roles := c.GetHeader("roles")
	isAdmin := strings.Contains(roles, "bingemate-admin")

	err := commentService.DeleteMovieComment(commentID, userID, isAdmin)

	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(204, "comment deleted")
}

// @Summary Delete tv show comment
// @Description Delete tv show comment
// @Tags Comment
// @Param commentID path string true "TV Show Comment ID"
// @Param user-id header string true "User ID"
// @Param roles header string false "User roles"
// @Produce json
// @Success 204 {string} string "comment deleted"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /comment/tv/{commentID} [delete]
func deleteTVShowComment(c *gin.Context, commentService *features.CommentService) {
	commentID := c.Param("commentID")
	if commentID == "" {
		c.JSON(400, errorResponse{Error: "commentID is required"})
		return
	}
	userID := c.GetHeader("user-id")
	if userID == "" {
		c.JSON(400, errorResponse{Error: "user-id must not be empty"})
		return
	}
	roles := c.GetHeader("roles")
	isAdmin := strings.Contains(roles, "bingemate-admin")

	err := commentService.DeleteTvShowComment(commentID, userID, isAdmin)

	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(204, "comment deleted")
}

// @Summary Update movie comment
// @Description Update movie comment
// @Tags Comment
// @Param commentID path string true "Movie Comment ID"
// @Param comment body commentRequest true "Comment"
// @Param user-id header string true "User ID"
// @Param roles header string false "User roles"
// @Produce json
// @Success 200 {object} commentResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /comment/movie/{commentID} [put]
func updateMovieComment(c *gin.Context, commentService *features.CommentService) {
	commentID := c.Param("commentID")
	if commentID == "" {
		c.JSON(400, errorResponse{Error: "commentID is required"})
		return
	}
	userID := c.GetHeader("user-id")
	if userID == "" {
		c.JSON(400, errorResponse{Error: "user-id must not be empty"})
		return
	}
	roles := c.GetHeader("roles")
	isAdmin := strings.Contains(roles, "bingemate-admin")

	var comment commentRequest
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(400, errorResponse{Error: err.Error()})
		return
	}
	if comment.Content == "" || strings.TrimSpace(comment.Content) == "" {
		c.JSON(400, errorResponse{Error: "comment must not be empty"})
		return
	}
	if len(comment.Content) > 1000 {
		c.JSON(400, errorResponse{Error: "comment must not be longer than 1000 characters"})
		return
	}
	commentResult, err := commentService.UpdateMovieComment(commentID, userID, isAdmin, comment.Content)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, toMovieCommentResponse(commentResult))
}

// @Summary Update tv show comment
// @Description Update tv show comment
// @Tags Comment
// @Param commentID path string true "TV Show Comment ID"
// @Param comment body commentRequest true "Comment"
// @Param user-id header string true "User ID"
// @Param roles header string false "User roles"
// @Produce json
// @Success 200 {object} commentResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /comment/tv/{commentID} [put]
func updateTVShowComment(c *gin.Context, commentService *features.CommentService) {
	commentID := c.Param("commentID")
	if commentID == "" {
		c.JSON(400, errorResponse{Error: "commentID is required"})
		return
	}
	userID := c.GetHeader("user-id")
	if userID == "" {
		c.JSON(400, errorResponse{Error: "user-id must not be empty"})
		return
	}
	roles := c.GetHeader("roles")
	isAdmin := strings.Contains(roles, "bingemate-admin")

	var comment commentRequest
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(400, errorResponse{Error: err.Error()})
		return
	}
	if comment.Content == "" || strings.TrimSpace(comment.Content) == "" {
		c.JSON(400, errorResponse{Error: "comment must not be empty"})
		return
	}
	if len(comment.Content) > 1000 {
		c.JSON(400, errorResponse{Error: "comment must not be longer than 1000 characters"})
		return
	}
	commentResult, err := commentService.UpdateTvShowComment(commentID, userID, isAdmin, comment.Content)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, toTVShowCommentResponse(commentResult))
}

// @Summary Get User's comments history
// @Description Get User's comments history
// @Tags Comment
// @Param userID path string true "User ID"
// @Param start query string false "Start date (YYYY-MM-DD)"
// @Param end query string false "End date (YYYY-MM-DD)"
// @Produce json
// @Success 200 {array} commentHistoryReponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /comment/user/history/{userID} [get]
func getUserCommentHistory(c *gin.Context, commentService *features.CommentService) {
	userID := c.Param("userID")
	if userID == "" {
		c.JSON(400, errorResponse{Error: "userID is required"})
		return
	}

	start := c.Query("start")
	end := c.Query("end")

	if start == "" {
		now := time.Now()
		currentYear := now.Year()
		currentMonth := now.Month()
		currentLocation := now.Location()

		firstDayOfCurrentMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
		start = firstDayOfCurrentMonth.Format("2006-01-02")
	}
	if end == "" {
		now := time.Now()
		currentYear := now.Year()
		currentMonth := now.Month()
		currentLocation := now.Location()

		firstDayOfCurrentMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
		lastDayOfCurrentMonth := firstDayOfCurrentMonth.AddDate(0, 1, -1)
		end = lastDayOfCurrentMonth.Format("2006-01-02")
	}

	movieCommentsHistory, err := commentService.GetUserMovieCommentsByRange(userID, start, end)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	tvShowCommentsHistory, err := commentService.GetUserTvShowCommentsByRange(userID, start, end)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, toCommentHistories(movieCommentsHistory, tvShowCommentsHistory))
}

// @Summary Get User's comments count
// @Description Get User's comments count
// @Tags Comment
// @Param userID path string true "User ID"
// @Produce json
// @Success 200 {object} int
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /comment/user/count/{userID} [get]
func getUserCommentCount(c *gin.Context, commentService *features.CommentService) {
	userID := c.Param("userID")
	if userID == "" {
		c.JSON(400, errorResponse{Error: "userID is required"})
		return
	}

	count, err := commentService.CountUserComments(userID)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, count)
}

// @Summary Get comments history
// @Description Get comments history
// @Tags Comment
// @Param start query string false "Start date (YYYY-MM-DD)"
// @Param end query string false "End date (YYYY-MM-DD)"
// @Produce json
// @Success 200 {array} commentHistoryReponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /comment/history [get]
func getCommentHistory(c *gin.Context, commentService *features.CommentService) {
	start := c.Query("start")
	end := c.Query("end")

	if start == "" {
		now := time.Now()
		currentYear := now.Year()
		currentMonth := now.Month()
		currentLocation := now.Location()

		firstDayOfCurrentMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
		start = firstDayOfCurrentMonth.Format("2006-01-02")
	}
	if end == "" {
		now := time.Now()
		currentYear := now.Year()
		currentMonth := now.Month()
		currentLocation := now.Location()

		firstDayOfCurrentMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
		lastDayOfCurrentMonth := firstDayOfCurrentMonth.AddDate(0, 1, -1)
		end = lastDayOfCurrentMonth.Format("2006-01-02")
	}

	movieCommentsHistory, err := commentService.GetMovieCommentsByRange(start, end)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	tvShowCommentsHistory, err := commentService.GetTvShowCommentsByRange(start, end)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, toCommentHistories(movieCommentsHistory, tvShowCommentsHistory))
}

// @Summary Get comments count
// @Description Get comments count
// @Tags Comment
// @Produce json
// @Success 200 {object} int
// @Failure 500 {object} errorResponse
// @Router /comment/count [get]
func getCommentCount(c *gin.Context, commentService *features.CommentService) {
	count, err := commentService.CountComments()
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, count)
}
