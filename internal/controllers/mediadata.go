package controllers

import (
	"errors"
	"github.com/bingemate/media-service/internal/features"
	"github.com/gin-gonic/gin"
	"strconv"
)

func InitMediaDataController(engine *gin.RouterGroup, mediaData *features.MediaData) {
	engine.GET("/movie-tmdb/:id", func(c *gin.Context) {
		getMovieByTMDB(c, mediaData)
	})
	engine.GET("/movie/:id", func(c *gin.Context) {
		getMovieByID(c, mediaData)
	})
	engine.GET("/base-tmdb/:id", func(c *gin.Context) {
		getMediaByTMDB(c, mediaData)
	})
	engine.GET("/base/:id", func(c *gin.Context) {
		getMediaByID(c, mediaData)
	})
}

// @Summary		Get Movie Metadata
// @Description	Get Movie Metadata by TMDB ID
// @Description	The rating is from BingeMate, not from TMDB (only if available, else from TMDB)
// @Tags			Media Data
// @Param			id path int true "TMDB ID"
// @Produce		json
// @Success		200	{object} movieResponse
// @Failure		400	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/media/movie-tmdb/{id} [get]
func getMovieByTMDB(c *gin.Context, mediaData *features.MediaData) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: err.Error(),
		})
		return
	}
	result, err := mediaData.GetMovieInfoByTMDB(id)
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, toMovieResponse(result))
}

// @Summary		Get Movie Metadata
// @Description	Get Movie Metadata by media ID
// @Description	The rating is from BingeMate, not from TMDB (only if available, else from TMDB)
// @Tags			Media Data
// @Param			id path int true "Media ID"
// @Produce		json
// @Success		200	{object} movieResponse
// @Failure		404	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/media/movie/{id} [get]
func getMovieByID(c *gin.Context, mediaData *features.MediaData) {
	id := c.Param("id")

	result, err := mediaData.GetMovieInfo(id)
	if err != nil {
		if errors.Is(err, features.ErrMediaNotFound) {
			c.JSON(404, errorResponse{
				Error: err.Error(),
			})
			return
		}
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, toMovieResponse(result))
}

// @Summary		Get media info
// @Description	Get media info by TMDB ID
// @Tags			Media Data
// @Param			id path int true "TMDB ID"
// @Produce		json
// @Success		200	{object} mediaResponse
// @Failure		400	{object} errorResponse
// @Failure		404	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/media/base-tmdb/{id} [get]
func getMediaByTMDB(c *gin.Context, mediaData *features.MediaData) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: err.Error(),
		})
		return
	}
	result, err := mediaData.GetMediaByTmdbID(id)
	if err != nil {
		if errors.Is(err, features.ErrMediaNotFound) {
			c.JSON(404, errorResponse{
				Error: err.Error(),
			})
			return
		}
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, toMediaResponse(result))
}

// @Summary		Get media info
// @Description	Get media info by Media ID
// @Tags			Media Data
// @Param			id path int true "Media ID"
// @Produce		json
// @Success		200	{object} mediaResponse
// @Failure		404	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/media/base/{id} [get]
func getMediaByID(c *gin.Context, mediaData *features.MediaData) {
	id := c.Param("id")
	result, err := mediaData.GetMediaByID(id)
	if err != nil {
		if errors.Is(err, features.ErrMediaNotFound) {
			c.JSON(404, errorResponse{
				Error: err.Error(),
			})
			return
		}
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, toMediaResponse(result))
}
