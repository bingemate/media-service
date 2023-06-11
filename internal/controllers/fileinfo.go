package controllers

import (
	"errors"
	"github.com/bingemate/media-service/internal/features"
	"github.com/gin-gonic/gin"
	"strconv"
)

func InitFileInfoController(engine *gin.RouterGroup, fileInfo *features.MediaFile) {
	engine.GET("movie/:id", func(c *gin.Context) {
		getMovieFileInfo(c, fileInfo)
	})
	engine.GET("episode/:id", func(c *gin.Context) {
		getEpisodeFileInfo(c, fileInfo)
	})
	engine.GET("tv/:id/available", func(c *gin.Context) {
		getAvailableEpisodes(c, fileInfo)
	})
}

// @Summary		Get movie file info by its Movie TMDB ID
// @Description	Get movie file info by its Movie TMDB ID
// @Tags			File
// @Param			id path int true "TMDB ID"
// @Produce		json
// @Success		200	{object} mediaFileResponse
// @Failure		400	{object} errorResponse
// @Failure		404	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/file/movie/{id} [get]
func getMovieFileInfo(c *gin.Context, mediaData *features.MediaFile) {
	mediaID, err := strconv.Atoi(c.Param("id"))
	result, err := mediaData.GetMovieFileInfo(mediaID)
	if err != nil {
		if errors.Is(err, features.ErrMediaNotFound) {
			c.JSON(404, errorResponse{
				Error: err.Error(),
			})
			return
		}
		if errors.Is(err, features.ErrInvalidMediaType) {
			c.JSON(400, errorResponse{
				Error: err.Error(),
			})
			return
		}
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, toMediaFileResponse(result))
}

// @Summary		Get episode file info by its Episode TMDB ID
// @Description	Get episode file info by its Episode TMDB ID
// @Tags			File
// @Param			id path int true "TMDB ID"
// @Produce		json
// @Success		200	{object} mediaFileResponse
// @Failure		400	{object} errorResponse
// @Failure		404	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/file/episode/{id} [get]
func getEpisodeFileInfo(c *gin.Context, mediaData *features.MediaFile) {
	mediaID, err := strconv.Atoi(c.Param("id"))
	result, err := mediaData.GetEpisodeFileInfo(mediaID)
	if err != nil {
		if errors.Is(err, features.ErrMediaNotFound) {
			c.JSON(404, errorResponse{
				Error: err.Error(),
			})
			return
		}
		if errors.Is(err, features.ErrInvalidMediaType) {
			c.JSON(400, errorResponse{
				Error: err.Error(),
			})
			return
		}
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, toMediaFileResponse(result))
}

// @Summary Get tv show available episodes
// @Description Get all available episodes id for a tv show
// @Tags File
// @Param id path int true "TMDB ID"
// @Produce json
// @Success 200 {array} int
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /file/tv/{id}/available [get]
func getAvailableEpisodes(c *gin.Context, mediaData *features.MediaFile) {
	mediaID, err := strconv.Atoi(c.Param("id"))
	result, err := mediaData.GetAvailableEpisode(mediaID)
	if err != nil {
		if errors.Is(err, features.ErrMediaNotFound) {
			c.JSON(404, errorResponse{
				Error: err.Error(),
			})
			return
		}
		if errors.Is(err, features.ErrInvalidMediaType) {
			c.JSON(400, errorResponse{
				Error: err.Error(),
			})
			return
		}
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, result)
}
