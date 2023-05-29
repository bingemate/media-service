package controllers

import (
	"github.com/bingemate/media-service/internal/features"
	"github.com/gin-gonic/gin"
	"strconv"
)

func InitMediaAssetsController(engine *gin.RouterGroup, mediaAssets *features.MediaAssetsData) {
	engine.GET("/movie-genre/:id", func(c *gin.Context) {
		getMovieGenre(c, mediaAssets)
	})
	engine.GET("/movie-genres", func(c *gin.Context) {
		getMovieGenres(c, mediaAssets)
	})
	engine.GET("/tv-genre/:id", func(c *gin.Context) {
		getTVGenre(c, mediaAssets)
	})
	engine.GET("/tv-genres", func(c *gin.Context) {
		getTVGenres(c, mediaAssets)
	})
	engine.GET("/studio/:id", func(c *gin.Context) {
		getStudio(c, mediaAssets)
	})
	engine.GET("/network/:id", func(c *gin.Context) {
		getNetwork(c, mediaAssets)
	})
	engine.GET("/actor/:id", func(c *gin.Context) {
		getActor(c, mediaAssets)
	})
}

// @Summary Get movie genre
// @Description Get movie genre by id
// @Tags Media Assets
// @Tags Movie
// @Param id path int true "Genre ID"
// @Produce json
// @Success 200 {object} genre
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /assets/movie-genre/{id} [get]
func getMovieGenre(c *gin.Context, mediaAssets *features.MediaAssetsData) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: err.Error(),
		})
		return
	}
	result, err := mediaAssets.GetMovieGenre(id)
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, toGenreResponse(result))
}

// @Summary Get movie genres
// @Description Get all movie genres
// @Tags Media Assets
// @Tags Movie
// @Produce json
// @Success 200 {array} genre
// @Failure 500 {object} errorResponse
// @Router /assets/movie-genres [get]
func getMovieGenres(c *gin.Context, mediaAssets *features.MediaAssetsData) {
	result, err := mediaAssets.GetMovieGenres()
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, toGenresResponse(result))
}

// @Summary Get tv genre
// @Description Get tv genre by id
// @Tags Media Assets
// @Tags TvShow
// @Param id path int true "Genre ID"
// @Produce json
// @Success 200 {object} genre
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /assets/tv-genre/{id} [get]
func getTVGenre(c *gin.Context, mediaAssets *features.MediaAssetsData) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: err.Error(),
		})
		return
	}
	result, err := mediaAssets.GetTVGenre(id)
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, toGenreResponse(result))
}

// @Summary Get tv genres
// @Description Get all tv genres
// @Tags Media Assets
// @Tags TvShow
// @Produce json
// @Success 200 {array} genre
// @Failure 500 {object} errorResponse
// @Router /assets/tv-genres [get]
func getTVGenres(c *gin.Context, mediaAssets *features.MediaAssetsData) {
	result, err := mediaAssets.GetTVGenres()
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, toGenresResponse(result))
}

// @Summary Get studio
// @Description Get studio by id
// @Tags Media Assets
// @Tags Movie
// @Param id path int true "Studio ID"
// @Produce json
// @Success 200 {object} studio
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /assets/studio/{id} [get]
func getStudio(c *gin.Context, mediaAssets *features.MediaAssetsData) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: err.Error(),
		})
		return
	}
	result, err := mediaAssets.GetStudio(id)
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, toStudioResponse(result))
}

// @Summary Get network
// @Description Get network by id
// @Tags Media Assets
// @Tags TvShow
// @Param id path int true "Network ID"
// @Produce json
// @Success 200 {object} studio
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /assets/network/{id} [get]
func getNetwork(c *gin.Context, mediaAssets *features.MediaAssetsData) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: err.Error(),
		})
		return
	}
	result, err := mediaAssets.GetNetwork(id)
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, toStudioResponse(result))
}

// @Summary Get actor
// @Description Get actor by id
// @Tags Media Assets
// @Param id path int true "Actor ID"
// @Produce json
// @Success 200 {object} actor
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /assets/actor/{id} [get]
func getActor(c *gin.Context, mediaAssets *features.MediaAssetsData) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: err.Error(),
		})
		return
	}
	result, err := mediaAssets.GetActor(id)
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, toActorResponse(result))
}
