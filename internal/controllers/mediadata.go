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
	engine.GET("/tvshow-tmdb/:id", func(c *gin.Context) {
		getTvShowByTMDB(c, mediaData)
	})
	engine.GET("/tvshow-episode-tmdb/:id/:season/:episode", func(c *gin.Context) {
		getTvShowEpisodeByTMDB(c, mediaData)
	})
	engine.GET("/tvshow-season-episodes-tmdb/:id/:season", func(c *gin.Context) {
		getTvShowSeasonEpisodesByTMDB(c, mediaData)
	})
	engine.GET("/base-tmdb/:id", func(c *gin.Context) {
		getMediaByTMDB(c, mediaData)
	})
}

// @Summary		Get Movie Metadata
// @Description	Get Movie Metadata by TMDB ID
// @Description	The rating is from BingeMate, not from TMDB (only if available, else from TMDB)
// @Tags			Media Data
// @Tags			Movie
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
	result, err := mediaData.GetMovieInfo(id)
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, toMovieResponse(result))
}

// @Summary		Get TvShow Metadata
// @Description	Get TvShow Metadata by TMDB ID
// @Description	The rating is from BingeMate, not from TMDB (only if available, else from TMDB)
// @Tags			Media Data
// @Tags			TvShow
// @Param			id path int true "TMDB ID"
// @Produce		json
// @Success		200	{object} tvShowResponse
// @Failure		400	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/media/tvshow-tmdb/{id} [get]
func getTvShowByTMDB(c *gin.Context, mediaData *features.MediaData) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: err.Error(),
		})
		return
	}
	result, err := mediaData.GetTvShowInfo(id)
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, toTVShowResponse(result))
}

// @Summary		Get TvShow Episode Metadata
// @Description	Get TvShow Episode Metadata by TvShow TMDB ID, Season and Episode Number
// @Tags			Media Data
// @Tags			TvEpisode
// @Param			id path int true "TvShow TMDB ID"
// @Param			season path int true "Season Number"
// @Param			episode path int true "Episode Number"
// @Produce		json
// @Success		200	{object} tvEpisodeResponse
// @Failure		400	{object} errorResponse
// @Failure		404	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/media/tvshow-episode-tmdb/{id}/{season}/{episode} [get]
func getTvShowEpisodeByTMDB(c *gin.Context, mediaData *features.MediaData) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: err.Error(),
		})
		return
	}
	season, err := strconv.Atoi(c.Param("season"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: err.Error(),
		})
		return
	}
	episode, err := strconv.Atoi(c.Param("episode"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: err.Error(),
		})
		return
	}
	result, err := mediaData.GetEpisodeInfo(id, season, episode)
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
	c.JSON(200, toTVEpisodeResponse(result))
}

// @Summary		Get TvShow Season Episodes Metadata
// @Description	Get TvShow Season Episodes Metadata by TvShow TMDB ID and Season Number
// @Tags			Media Data
// @Tags			TvEpisode
// @Param			id path int true "TvShow TMDB ID"
// @Param			season path int true "Season Number"
// @Produce		json
// @Success		200	{array} tvEpisodeResponse
// @Failure		400	{object} errorResponse
// @Failure		404	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/media/tvshow-season-episodes-tmdb/{id}/{season} [get]
func getTvShowSeasonEpisodesByTMDB(c *gin.Context, mediaData *features.MediaData) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: err.Error(),
		})
		return
	}
	season, err := strconv.Atoi(c.Param("season"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: err.Error(),
		})
		return
	}
	result, err := mediaData.GetSeasonEpisodes(id, season)
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
	c.JSON(200, toTVEpisodesResponse(result))
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
