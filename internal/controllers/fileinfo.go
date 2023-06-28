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
	engine.GET("movie/search", func(c *gin.Context) {
		searchMovies(c, fileInfo)
	})
	engine.GET("movie/count", func(c *gin.Context) {
		countAvailableMovies(c, fileInfo)
	})
	engine.GET("/movie/duration", func(c *gin.Context) {
		countMoviesTotalDuration(c, fileInfo)
	})
	engine.GET("episode/:id", func(c *gin.Context) {
		getEpisodeFileInfo(c, fileInfo)
	})
	engine.GET("episode/search", func(c *gin.Context) {
		searchEpisodes(c, fileInfo)
	})
	engine.GET("episode/count", func(c *gin.Context) {
		countAvailableEpisodes(c, fileInfo)
	})
	engine.GET("episode/duration", func(c *gin.Context) {
		countEpisodesTotalDuration(c, fileInfo)
	})
	engine.GET("tv/:id/available", func(c *gin.Context) {
		getAvailableEpisodes(c, fileInfo)
	})
	engine.GET("tv/count", func(c *gin.Context) {
		countAvailableTvShows(c, fileInfo)
	})
	engine.DELETE(":id", func(c *gin.Context) {
		deleteFile(c, fileInfo)
	})
	engine.GET("size", func(c *gin.Context) {
		getTotalSize(c, fileInfo)
	})
	engine.GET("count", func(c *gin.Context) {
		countFiles(c, fileInfo)
	})
	engine.GET("available", func(c *gin.Context) {
		getAvailableSpace(c, fileInfo)
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

// @Summary Search tv show episodes files
// @Description Search tv show episodes files
// @Tags File
// @Param page query int false "Page number"
// @Param limit query int false "Page limit"
// @Param query query string true "Search query"
// @Produce json
// @Success 200 {object} episodeFilesResult
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /file/episode/search [get]
func searchEpisodes(c *gin.Context, mediaData *features.MediaFile) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: err.Error(),
		})
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: err.Error(),
		})
		return
	}
	query := c.Query("query")
	result, total, err := mediaData.SearchEpisodeFiles(query, page, limit)

	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, episodeFilesResult{
		Results: toEpisodeFilesResponse(result),
		Total:   total,
	})
}

// @Summary Search movie files
// @Description Search movie files
// @Tags File
// @Param page query int false "Page number"
// @Param limit query int false "Page limit"
// @Param query query string true "Search query"
// @Produce json
// @Success 200 {object} movieFilesResult
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /file/movie/search [get]
func searchMovies(c *gin.Context, mediaData *features.MediaFile) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: err.Error(),
		})
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: err.Error(),
		})
		return
	}
	query := c.Query("query")
	result, total, err := mediaData.SearchMovieFiles(query, page, limit)

	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, movieFilesResult{
		Results: toMovieFilesResponse(result),
		Total:   total,
	})
}

// @Summary Delete a file
// @Description Delete a file
// @Tags File
// @Param id path int true "File ID"
// @Produce json
// @Success 200 {string} string "OK"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /file/{id} [delete]
func deleteFile(c *gin.Context, mediaData *features.MediaFile) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, errorResponse{
			Error: "id is required",
		})
		return
	}

	err := mediaData.DeleteMediaFile(id)
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, "OK")
}

// @Summary Get total size
// @Description Get total size taken by all files
// @Tags File
// @Produce json
// @Success 200 {int} int "Total size in bytes"
// @Failure 500 {object} errorResponse
// @Router /file/size [get]
func getTotalSize(c *gin.Context, mediaData *features.MediaFile) {
	size, err := mediaData.MediaFilesTotalSize()
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, size)
}

// @Summary Count files
// @Description Count files
// @Tags File
// @Produce json
// @Success 200 {int} int "Total number of files"
// @Failure 500 {object} errorResponse
// @Router /file/count [get]
func countFiles(c *gin.Context, mediaData *features.MediaFile) {
	count, err := mediaData.MediaFilesCount()
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, count)
}

// @Summary Get available space
// @Description Get available space
// @Tags File
// @Produce json
// @Success 200 {int} int "Available space in bytes"
// @Failure 500 {object} errorResponse
// @Router /file/available [get]
func getAvailableSpace(c *gin.Context, mediaData *features.MediaFile) {
	size, err := mediaData.AvailableSpace()
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, size)
}

// @Summary Count available movies
// @Description Count available movies
// @Tags File
// @Produce json
// @Success 200 {int} int "Total number of available movies"
// @Failure 500 {object} errorResponse
// @Router /file/movie/count [get]
func countAvailableMovies(c *gin.Context, mediaData *features.MediaFile) {
	count, err := mediaData.CountAvailableMovies()
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, count)
}

// @Summary Count available episodes
// @Description Count available episodes
// @Tags File
// @Produce json
// @Success 200 {int} int "Total number of available episodes"
// @Failure 500 {object} errorResponse
// @Router /file/episode/count [get]
func countAvailableEpisodes(c *gin.Context, mediaData *features.MediaFile) {
	count, err := mediaData.CountAvailableEpisodes()
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, count)
}

// @Summary Count available tv shows
// @Description Count available tv shows
// @Tags File
// @Produce json
// @Success 200 {int} int "Total number of available tv shows"
// @Failure 500 {object} errorResponse
// @Router /file/tv/count [get]
func countAvailableTvShows(c *gin.Context, mediaData *features.MediaFile) {
	count, err := mediaData.CountAvailableTvShows()
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, count)
}

// @Summary Count movies total duration
// @Description Count movies total duration
// @Tags File
// @Produce json
// @Success 200 {int} int "Total duration in seconds"
// @Failure 500 {object} errorResponse
// @Router /file/movie/duration [get]
func countMoviesTotalDuration(c *gin.Context, mediaData *features.MediaFile) {
	duration, err := mediaData.CountMoviesTotalDuration()
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, duration)
}

// @Summary Count episodes total duration
// @Description Count episodes total duration
// @Tags File
// @Produce json
// @Success 200 {int} int "Total duration in seconds"
// @Failure 500 {object} errorResponse
// @Router /file/episode/duration [get]
func countEpisodesTotalDuration(c *gin.Context, mediaData *features.MediaFile) {
	duration, err := mediaData.CountEpisodesTotalDuration()
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, duration)
}
