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
	engine.GET("/movie-tmdb/:id/short", func(c *gin.Context) {
		getMovieShortByTMDB(c, mediaData)
	})
	engine.GET("/movies-tmdb", func(c *gin.Context) {
		getMoviesShortByTMDB(c, mediaData)
	})
	engine.GET("/tvshow-tmdb/:id", func(c *gin.Context) {
		getTvShowByTMDB(c, mediaData)
	})
	engine.GET("/tvshow-tmdb/:id/short", func(c *gin.Context) {
		getTvShowShortByTMDB(c, mediaData)
	})
	engine.GET("/tvshows-tmdb", func(c *gin.Context) {
		getTvShowsShortByTMDB(c, mediaData)
	})
	engine.GET("/tvshow-episode-tmdb/:id/:season/:episode", func(c *gin.Context) {
		getTvShowEpisodeByTMDB(c, mediaData)
	})
	engine.GET("/tvshow-season-episodes-tmdb/:id/:season", func(c *gin.Context) {
		getTvShowSeasonEpisodesByTMDB(c, mediaData)
	})
	engine.GET("/episode-tmdb/:id", func(c *gin.Context) {
		getEpisodeByTMDB(c, mediaData)
	})
	engine.GET("/episodes-tmdb", func(c *gin.Context) {
		getEpisodesByTMDB(c, mediaData)
	})
	engine.GET("/base/movie/:id", func(c *gin.Context) {
		getMovieBaseByTMDB(c, mediaData)
	})
	engine.GET("/base/tv/:id", func(c *gin.Context) {
		getTvShowBaseByTMDB(c, mediaData)
	})
	engine.GET("/base/episode/:id", func(c *gin.Context) {
		getEpisodeBaseByTMDB(c, mediaData)
	})
	engine.GET("/base/episodes", func(c *gin.Context) {
		getEpisodesBaseByTMDB(c, mediaData)
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
	result, presence, err := mediaData.GetMovieInfo(id)
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, toMovieResponse(result, presence))
}

// @Summary		Get Movie Short Metadata
// @Description	Get Movie Short Metadata by TMDB ID
// @Description	The rating is from BingeMate, not from TMDB (only if available, else from TMDB)
// @Tags			Media Data
// @Tags			Movie
// @Param			id path int true "TMDB ID"
// @Produce		json
// @Success		200	{object} movieResponse
// @Failure		400	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/media/movie-tmdb/{id}/short [get]
func getMovieShortByTMDB(c *gin.Context, mediaData *features.MediaData) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: err.Error(),
		})
		return
	}
	result, presence, err := mediaData.GetMovieShortInfo(id)
	if err != nil {
		c.JSON(404, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, toMovieResponse(result, presence))
}

// @Summary		Get Movies Short Metadata
// @Description	Get Movies Short Metadata by TMDB ID
// @Description	The rating is from BingeMate, not from TMDB (only if available, else from TMDB)
// @Tags			Media Data
// @Tags			Movie
// @Param			ids query []int true "TMDB IDs"
// @Produce		json
// @Success		200	{array} movieResponse
// @Failure		400	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/media/movies-tmdb [get]
func getMoviesShortByTMDB(c *gin.Context, mediaData *features.MediaData) {
	idsStr := c.QueryArray("ids")
	ids := make([]int, len(idsStr))
	for i, idStr := range idsStr {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(400, errorResponse{
				Error: err.Error(),
			})
			return
		}
		ids[i] = id
	}

	result, presence, err := mediaData.GetMoviesShortInfo(ids)
	if err != nil {
		c.JSON(404, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, toMoviesResponse(result, presence))
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
	result, presence, err := mediaData.GetTvShowInfo(id)
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, toTVShowResponse(result, presence))
}

// @Summary		Get TvShow Short Metadata
// @Description	Get TvShow Short Metadata by TMDB ID
// @Description	The rating is from BingeMate, not from TMDB (only if available, else from TMDB)
// @Tags			Media Data
// @Tags			TvShow
// @Param			id path int true "TMDB ID"
// @Produce		json
// @Success		200	{object} tvShowResponse
// @Failure		400	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/media/tvshow-tmdb/{id}/short [get]
func getTvShowShortByTMDB(c *gin.Context, mediaData *features.MediaData) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: err.Error(),
		})
		return
	}
	result, presence, err := mediaData.GetTvShowShortInfo(id)
	if err != nil {
		c.JSON(404, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, toTVShowResponse(result, presence))
}

// @Summary		Get TvShows Short Metadata
// @Description	Get TvShows Short Metadata by TMDB ID
// @Description	The rating is from BingeMate, not from TMDB (only if available, else from TMDB)
// @Tags			Media Data
// @Tags			TvShow
// @Param			ids query []int true "TMDB IDs"
// @Produce		json
// @Success		200	{array} tvShowResponse
// @Failure		400	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/media/tvshows-tmdb [get]
func getTvShowsShortByTMDB(c *gin.Context, mediaData *features.MediaData) {
	idsStr := c.QueryArray("ids")
	ids := make([]int, len(idsStr))
	for i, idStr := range idsStr {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(400, errorResponse{
				Error: err.Error(),
			})
			return
		}
		ids[i] = id
	}

	result, presence, err := mediaData.GetTvShowsShortInfo(ids)
	if err != nil {
		c.JSON(404, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, toTVShowsResponse(result, presence))
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
	result, presence, err := mediaData.GetEpisodeInfo(id, season, episode)
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
	c.JSON(200, toTVEpisodeResponse(result, presence))
}

// @Summary Get TvShow Episode Metadata by TMDB ID
// @Description Get TvShow Episode Metadata by TMDB ID
// @Tags Media Data
// @Tags TvEpisode
// @Param id path int true "TMDB ID"
// @Produce json
// @Success 200 {object} tvEpisodeResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /media/episode-tmdb/{id} [get]
func getEpisodeByTMDB(c *gin.Context, mediaData *features.MediaData) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: err.Error(),
		})
		return
	}
	result, presence, err := mediaData.GetEpisodeInfoByID(id)
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
	c.JSON(200, toTVEpisodeResponse(result, presence))
}

// @Summary		Get TvShow Episodes metadata by TMDB IDs
// @Description	Get TvShow Episodes metadata by TMDB IDs
// @Tags			Media Data
// @Tags			TvEpisode
// @Param			ids query []int true "TMDB IDs"
// @Produce		json
// @Success		200	{array} tvEpisodeResponse
// @Failure		400	{object} errorResponse
// @Failure		404	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/media/episodes-tmdb [get]
func getEpisodesByTMDB(c *gin.Context, mediaData *features.MediaData) {
	idsStr := c.QueryArray("ids")
	ids := make([]int, len(idsStr))
	for i, idStr := range idsStr {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(400, errorResponse{
				Error: err.Error(),
			})
			return
		}
		ids[i] = id
	}

	results, presences, err := mediaData.GetEpisodesInfoByIDs(ids)
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
	c.JSON(200, toTVEpisodesResponse(results, presences))
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
	result, presence, err := mediaData.GetSeasonEpisodes(id, season)
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
	c.JSON(200, toTVEpisodesResponse(result, presence))
}

// @Summary		Get movie base info
// @Description	Get movie base info by TMDB ID
// @Tags			Movie Data
// @Tags			Base
// @Param			id path int true "TMDB ID"
// @Produce		json
// @Success		200	{object} mediaResponse
// @Failure		400	{object} errorResponse
// @Failure		404	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/media/base/movie/{id} [get]
func getMovieBaseByTMDB(c *gin.Context, mediaData *features.MediaData) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: err.Error(),
		})
		return
	}
	result, err := mediaData.GetMovieByID(id)
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
	c.JSON(200, toMovieMediaResponse(result))
}

// @Summary		Get tvshow base info
// @Description	Get tvshow base info by TMDB ID
// @Tags			TvShow Data
// @Tags			Base
// @Param			id path int true "TMDB ID"
// @Produce		json
// @Success		200	{object} mediaResponse
// @Failure		400	{object} errorResponse
// @Failure		404	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/media/base/tv/{id} [get]
func getTvShowBaseByTMDB(c *gin.Context, mediaData *features.MediaData) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: err.Error(),
		})
		return
	}
	result, err := mediaData.GetTvShowByID(id)
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
	c.JSON(200, toTVShowMediaResponse(result))
}

// @Summary		Get episode base info
// @Description	Get episode base info by TMDB ID
// @Tags			TvEpisode Data
// @Tags			Base
// @Param			id path int true "TMDB ID"
// @Produce		json
// @Success		200	{object} mediaResponse
// @Failure		400	{object} errorResponse
// @Failure		404	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/media/base/episode/{id} [get]
func getEpisodeBaseByTMDB(c *gin.Context, mediaData *features.MediaData) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: err.Error(),
		})
		return
	}
	result, err := mediaData.GetEpisodeByID(id)
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
	c.JSON(200, toEpisodeMediaResponse(result))
}

// @Summary		Get episodes base info
// @Description	Get episodes base info by TMDB ID
// @Tags			TvEpisode Data
// @Tags			Base
// @Param			ids query []int true "TMDB IDs"
// @Produce		json
// @Success		200	{array} mediaResponse
// @Failure		400	{object} errorResponse
// @Failure		404	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/media/base/episodes [get]
func getEpisodesBaseByTMDB(c *gin.Context, mediaData *features.MediaData) {
	idsStr := c.QueryArray("ids")
	ids := make([]int, len(idsStr))
	for i, idStr := range idsStr {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(400, errorResponse{
				Error: err.Error(),
			})
			return
		}
		ids[i] = id
	}
	result, err := mediaData.GetEpisodesByIDs(ids)
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
	c.JSON(200, toEpisodesMediaResponse(result))
}
