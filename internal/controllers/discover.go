package controllers

import (
	"github.com/bingemate/media-service/internal/features"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func InitDiscoverController(engine *gin.RouterGroup, mediaDiscover *features.MediaDiscovery) {
	engine.GET("movie/search", func(c *gin.Context) {
		searchMovie(c, mediaDiscover)
	})
	engine.GET("tv/search", func(c *gin.Context) {
		searchTv(c, mediaDiscover)
	})
	engine.GET("movie/popular", func(c *gin.Context) {
		getPopularMovies(c, mediaDiscover)
	})
	engine.GET("tv/popular", func(c *gin.Context) {
		getPopularTvShows(c, mediaDiscover)
	})
	engine.GET("movie/recent", func(c *gin.Context) {
		getRecentMovies(c, mediaDiscover)
	})
	engine.GET("tv/recent", func(c *gin.Context) {
		getRecentTvShows(c, mediaDiscover)
	})
	engine.GET("movie/genre", func(c *gin.Context) {
		getMoviesByGenre(c, mediaDiscover)
	})
	engine.GET("tv/genre", func(c *gin.Context) {
		getTvShowsByGenre(c, mediaDiscover)
	})
	engine.GET("movie/actor", func(c *gin.Context) {
		getMoviesByActor(c, mediaDiscover)
	})
	engine.GET("tv/actor", func(c *gin.Context) {
		getTvShowsByActor(c, mediaDiscover)
	})
	engine.GET("movie/director", func(c *gin.Context) {
		getMoviesByDirector(c, mediaDiscover)
	})
	engine.GET("movie/studio", func(c *gin.Context) {
		getMoviesByStudio(c, mediaDiscover)
	})
	engine.GET("tv/network", func(c *gin.Context) {
		getTvShowsByNetwork(c, mediaDiscover)
	})
	engine.GET("movie/recommendations/:movie", func(c *gin.Context) {
		getMovieRecommendations(c, mediaDiscover)
	})
	engine.GET("tv/recommendations/:tv", func(c *gin.Context) {
		getTvShowRecommendations(c, mediaDiscover)
	})
	engine.GET("media/comments", func(c *gin.Context) {
		getMediasByComments(c, mediaDiscover)
	})
}

// @Summary		Search movies
// @Description	Search movies by query
// @Tags			Discover
// @Tags			Movie
// @Param			query query string true "Search query"
// @Param			page query int false "Page number"
// @Param           available query bool false "Only available movies"
// @Produce		json
// @Success		200	{object} movieResults
// @Failure		400	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/discover/movie/search [get]
func searchMovie(c *gin.Context, mediaDiscover *features.MediaDiscovery) {
	query := strings.TrimSpace(c.Query("query"))
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	available, err := strconv.ParseBool(c.Query("available"))
	if err != nil {
		available = false
	}
	if query == "" {
		c.JSON(400, errorResponse{
			Error: "query is required",
		})
		return
	}
	result, presence, err := mediaDiscover.SearchMovie(query, page, available)
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, movieResults{
		TotalPage:   result.TotalPage,
		TotalResult: result.TotalResult,
		Results:     toMoviesResponse(result.Results, presence),
	})
}

// @Summary		Search tv shows
// @Description	Search tv shows by query
// @Tags			Discover
// @Tags			TvShow
// @Param			query query string true "Search query"
// @Param			page query int false "Page number"
// @Param           available query bool false "Only available tv shows"
// @Produce		json
// @Success		200	{object} tvShowResults
// @Failure		400	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/discover/tv/search [get]
func searchTv(c *gin.Context, mediaDiscover *features.MediaDiscovery) {
	query := strings.TrimSpace(c.Query("query"))
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	available, err := strconv.ParseBool(c.Query("available"))
	if err != nil {
		available = false
	}
	if query == "" {
		c.JSON(400, errorResponse{
			Error: "query is required",
		})
		return
	}
	result, presence, err := mediaDiscover.SearchShow(query, page, available)
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, tvShowResults{
		TotalPage:   result.TotalPage,
		TotalResult: result.TotalResult,
		Results:     toTVShowsResponse(result.Results, presence),
	})
}

// @Summary		Get popular movies
// @Description	Get popular movies
// @Tags			Discover
// @Tags			Movie
// @Param			page query int false "Page number"
// @Param           available query bool false "Only available movies"
// @Produce		json
// @Success		200	{object} movieResults
// @Failure		500	{object} errorResponse
// @Router			/discover/movie/popular [get]
func getPopularMovies(c *gin.Context, mediaDiscover *features.MediaDiscovery) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	available, err := strconv.ParseBool(c.Query("available"))
	if err != nil {
		available = false
	}
	result, presence, err := mediaDiscover.GetPopularMovies(page, available)
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, movieResults{
		TotalPage:   result.TotalPage,
		TotalResult: result.TotalResult,
		Results:     toMoviesResponse(result.Results, presence),
	})
}

// @Summary		Get popular tv shows
// @Description	Get popular tv shows
// @Tags			Discover
// @Tags			TvShow
// @Param			page query int false "Page number"
// @Param           available query bool false "Only available tv shows"
// @Produce		json
// @Success		200	{object} tvShowResults
// @Failure		500	{object} errorResponse
// @Router			/discover/tv/popular [get]
func getPopularTvShows(c *gin.Context, mediaDiscover *features.MediaDiscovery) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	available, err := strconv.ParseBool(c.Query("available"))
	if err != nil {
		available = false
	}
	result, presence, err := mediaDiscover.GetPopularShows(page, available)
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, tvShowResults{
		TotalPage:   result.TotalPage,
		TotalResult: result.TotalResult,
		Results:     toTVShowsResponse(result.Results, presence),
	})
}

// @Summary		Get recent movies
// @Description	Get recent movies
// @Tags			Discover
// @Tags			Movie
// @Param           available query bool false "Only available movies"
// @Produce		json
// @Success		200	{array} movieResponse
// @Failure		500	{object} errorResponse
// @Router			/discover/movie/recent [get]
func getRecentMovies(c *gin.Context, mediaDiscover *features.MediaDiscovery) {
	available, err := strconv.ParseBool(c.Query("available"))
	if err != nil {
		available = false
	}
	result, presence, err := mediaDiscover.GetRecentMovies(available)
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, toMoviesResponse(result, presence))
}

// @Summary		Get recent tv shows
// @Description	Get recent tv shows
// @Tags			Discover
// @Tags			TvShow
// @Param           available query bool false "Only available tv shows"
// @Produce		json
// @Success		200	{array} tvShowResponse
// @Failure		500	{object} errorResponse
// @Router			/discover/tv/recent [get]
func getRecentTvShows(c *gin.Context, mediaDiscover *features.MediaDiscovery) {
	available, err := strconv.ParseBool(c.Query("available"))
	if err != nil {
		available = false
	}
	result, presence, err := mediaDiscover.GetRecentShows(available)
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, toTVShowsResponse(result, presence))
}

// @Summary		Get movies by genre
// @Description	Get movies by genre
// @Tags			Discover
// @Tags			Movie
// @Param			genre query int true "Genre id"
// @Param			page query int false "Page number"
// @Produce		json
// @Success		200	{object} movieResults
// @Failure		400	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/discover/movie/genre [get]
func getMoviesByGenre(c *gin.Context, mediaDiscover *features.MediaDiscovery) {
	genre, err := strconv.Atoi(c.Query("genre"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: "genre is required",
		})
		return
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	result, presence, err := mediaDiscover.GetMoviesByGenre(genre, page)
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, movieResults{
		TotalPage:   result.TotalPage,
		TotalResult: result.TotalResult,
		Results:     toMoviesResponse(result.Results, presence),
	})
}

// @Summary		Get tv shows by genre
// @Description	Get tv shows by genre
// @Tags			Discover
// @Tags			TvShow
// @Param			genre query int true "Genre id"
// @Param			page query int false "Page number"
// @Produce		json
// @Success		200	{object} tvShowResults
// @Failure		400	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/discover/tv/genre [get]
func getTvShowsByGenre(c *gin.Context, mediaDiscover *features.MediaDiscovery) {
	genre, err := strconv.Atoi(c.Query("genre"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: "genre is required",
		})
		return
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	result, presence, err := mediaDiscover.GetShowsByGenre(genre, page)
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, tvShowResults{
		TotalPage:   result.TotalPage,
		TotalResult: result.TotalResult,
		Results:     toTVShowsResponse(result.Results, presence),
	})
}

// @Summary		Get movies by actor
// @Description	Get movies by actor
// @Tags			Discover
// @Tags			Movie
// @Param			actor query int true "Actor id"
// @Param			page query int false "Page number"
// @Produce		json
// @Success		200	{object} movieResults
// @Failure		400	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/discover/movie/actor [get]
func getMoviesByActor(c *gin.Context, mediaDiscover *features.MediaDiscovery) {
	actor, err := strconv.Atoi(c.Query("actor"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: "actor is required",
		})
		return
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	result, presence, err := mediaDiscover.GetMoviesByActor(actor, page)
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, movieResults{
		TotalPage:   result.TotalPage,
		TotalResult: result.TotalResult,
		Results:     toMoviesResponse(result.Results, presence),
	})
}

// @Summary		Get tv shows by actor
// @Description	Get tv shows by actor
// @Tags			Discover
// @Tags			TvShow
// @Param			actor query int true "Actor id"
// @Param			page query int false "Page number"
// @Produce		json
// @Success		200	{object} tvShowResults
// @Failure		400	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/discover/tv/actor [get]
func getTvShowsByActor(c *gin.Context, mediaDiscover *features.MediaDiscovery) {
	actor, err := strconv.Atoi(c.Query("actor"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: "actor is required",
		})
		return
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	result, presence, err := mediaDiscover.GetShowsByActor(actor, page)
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, tvShowResults{
		TotalPage:   result.TotalPage,
		TotalResult: result.TotalResult,
		Results:     toTVShowsResponse(result.Results, presence),
	})
}

// @Summary		Get movies by director
// @Description	Get movies by director
// @Tags			Discover
// @Tags			Movie
// @Param			director query int true "Director id"
// @Param			page query int false "Page number"
// @Produce		json
// @Success		200	{object} movieResults
// @Failure		400	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/discover/movie/director [get]
func getMoviesByDirector(c *gin.Context, mediaDiscover *features.MediaDiscovery) {
	director, err := strconv.Atoi(c.Query("director"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: "director is required",
		})
		return
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	result, presence, err := mediaDiscover.GetMoviesByDirector(director, page)
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, movieResults{
		TotalPage:   result.TotalPage,
		TotalResult: result.TotalResult,
		Results:     toMoviesResponse(result.Results, presence),
	})
}

// @Summary		Get movies by studio
// @Description	Get movies by studio
// @Tags			Discover
// @Tags			Movie
// @Param			studio query int true "Studio id"
// @Param			page query int false "Page number"
// @Produce		json
// @Success		200	{object} movieResults
// @Failure		400	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/discover/movie/studio [get]
func getMoviesByStudio(c *gin.Context, mediaDiscover *features.MediaDiscovery) {
	studio, err := strconv.Atoi(c.Query("studio"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: "studio is required",
		})
		return
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	result, presence, err := mediaDiscover.GetMoviesByStudio(studio, page)
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, movieResults{
		TotalPage:   result.TotalPage,
		TotalResult: result.TotalResult,
		Results:     toMoviesResponse(result.Results, presence),
	})
}

// @Summary		Get tv shows by network
// @Description	Get tv shows by network
// @Tags			Discover
// @Tags			TvShow
// @Param			network query int true "Network id"
// @Param			page query int false "Page number"
// @Produce		json
// @Success		200	{object} tvShowResults
// @Failure		400	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/discover/tv/network [get]
func getTvShowsByNetwork(c *gin.Context, mediaDiscover *features.MediaDiscovery) {
	network, err := strconv.Atoi(c.Query("network"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: "network is required",
		})
		return
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	result, presence, err := mediaDiscover.GetShowsByNetwork(network, page)
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, tvShowResults{
		TotalPage:   result.TotalPage,
		TotalResult: result.TotalResult,
		Results:     toTVShowsResponse(result.Results, presence),
	})
}

// @Summary		Get movie's recommendations
// @Description	Get movie's recommendations
// @Tags			Discover
// @Tags			Movie
// @Param			movie path int true "Movie id"
// @Produce		json
// @Success		200	{array} movieResponse
// @Failure		400	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/discover/movie/recommendations/{movie} [get]
func getMovieRecommendations(c *gin.Context, mediaDiscover *features.MediaDiscovery) {
	movie, err := strconv.Atoi(c.Param("movie"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: "movie is required",
		})
		return
	}
	result, presence, err := mediaDiscover.GetMovieRecommendations(movie)
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, toMoviesResponse(result, presence))
}

// @Summary		Get tv show's recommendations
// @Description	Get tv show's recommendations
// @Tags			Discover
// @Tags			TvShow
// @Param			show path int true "Show id"
// @Produce		json
// @Success		200	{array} tvShowResponse
// @Failure		400	{object} errorResponse
// @Failure		500	{object} errorResponse
// @Router			/discover/tv/recommendations/{tv} [get]
func getTvShowRecommendations(c *gin.Context, mediaDiscover *features.MediaDiscovery) {
	show, err := strconv.Atoi(c.Param("tv"))
	if err != nil {
		c.JSON(400, errorResponse{
			Error: "tv is required",
		})
		return
	}
	result, presence, err := mediaDiscover.GetShowRecommendations(show)
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, toTVShowsResponse(result, presence))
}

// @Summary		Get medias by comment
// @Description	Get medias ordered by number of comments
// @Tags			Discover
// @Tags			Media
// @Produce		json
// @Success		200	{array} int
// @Failure		500	{object} errorResponse
// @Router			/discover/media/comments [get]
func getMediasByComments(c *gin.Context, mediaDiscover *features.MediaDiscovery) {
	result, err := mediaDiscover.GetMediasByComments()
	if err != nil {
		c.JSON(500, errorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(200, result)
}
