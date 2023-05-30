package controllers

import (
	"github.com/bingemate/media-service/internal/features"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func InitCommentController(engine *gin.RouterGroup, commentService *features.CommentService) {
	engine.GET("/media/:mediaID", func(c *gin.Context) {
		getComments(c, commentService)
	})
	engine.GET("/user/:userID", func(c *gin.Context) {
		getUserComments(c, commentService)
	})
	engine.POST("/media/:mediaID", func(c *gin.Context) {
		addComment(c, commentService)
	})
	engine.DELETE("/:commentID", func(c *gin.Context) {
		deleteComment(c, commentService)
	})
	engine.PUT("/:commentID", func(c *gin.Context) {
		updateComment(c, commentService)
	})

}

// @Summary Get media's comments
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
}

// @Summary Get user's comments
// @Description Get user's comments
// @Tags Comment
// @Param page query int false "Page number"
// @Param userID path string true "User ID"
// @Produce json
// @Success 200 {object} commentResults
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /comment/user/{userID} [get]
func getUserComments(c *gin.Context, commentService *features.CommentService) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	userID := c.Param("userID")
	if userID == "" {
		c.JSON(400, errorResponse{Error: "userID must be a string"})
		return
	}
	comments, total, err := commentService.GetUserComments(userID, page)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, commentResults{
		Results:     toCommentsResponse(comments),
		TotalResult: total,
	})
}

// @Summary Add comment
// @Description Add comment to a media
// @Tags Comment
// @Param mediaID path int true "Media ID"
// @Param comment body commentRequest true "Comment"
// @Param user-id header string true "User ID"
// @Produce json
// @Success 200 {object} commentResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /comment/media/{mediaID} [post]
func addComment(c *gin.Context, commentService *features.CommentService) {
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
	commentResult, err := commentService.AddComment(userID, mediaID, comment.Content)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, toCommentResponse(commentResult))
}

// @Summary Delete comment
// @Description Delete comment
// @Tags Comment
// @Param commentID path string true "Comment ID"
// @Param user-id header string true "User ID"
// @Param roles header string false "User roles"
// @Produce json
// @Success 204 {string} string "comment deleted"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /comment/{commentID} [delete]
func deleteComment(c *gin.Context, commentService *features.CommentService) {
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

	err := commentService.DeleteComment(userID, commentID, isAdmin)

	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(204, "comment deleted")
}

// @Summary Update comment
// @Description Update comment
// @Tags Comment
// @Param commentID path string true "Comment ID"
// @Param comment body commentRequest true "Comment"
// @Param user-id header string true "User ID"
// @Param roles header string false "User roles"
// @Produce json
// @Success 200 {object} commentResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /comment/{commentID} [put]
func updateComment(c *gin.Context, commentService *features.CommentService) {
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
	commentResult, err := commentService.UpdateComment(userID, commentID, isAdmin, comment.Content)
	if err != nil {
		c.JSON(500, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(200, toCommentResponse(commentResult))
}
