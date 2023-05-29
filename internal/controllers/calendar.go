package controllers

import (
	"github.com/bingemate/media-service/internal/features"
	"github.com/gin-gonic/gin"
)

func InitCalendarController(engine *gin.RouterGroup, calendarService *features.CalendarService) {
	engine.GET("/movies", func(c *gin.Context) {
		getMoviesCalendar(c, calendarService)
	})
	engine.GET("/tvshows", func(c *gin.Context) {
		getTvShowsCalendar(c, calendarService)
	})
}

// @Summary Get movies calendar
// @Description Get movies calendar
// @Tags  Calendar
// @Tags Movie
// @Param user-id header string true "User ID"
// @Produce  json
// @Success 200 {array} movieResults
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /calendar/movies [get]
func getMoviesCalendar(c *gin.Context, calendarService *features.CalendarService) {
	userID := c.GetHeader("user-id")
	if userID == "" {
		c.JSON(400, gin.H{"error": "user-id header is required"})
		return
	}
	movies, presence, err := calendarService.GetMoviesCalendar(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, toMoviesResponse(movies, presence))
}

// @Summary Get tv shows calendar
// @Description Get tv shows calendar
// @Tags  Calendar
// @Tags TvShow
// @Param user-id header string true "User ID"
// @Produce  json
// @Success 200 {array} tvShowResults
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /calendar/tvshows [get]
func getTvShowsCalendar(c *gin.Context, calendarService *features.CalendarService) {
	userID := c.GetHeader("user-id")
	if userID == "" {
		c.JSON(400, gin.H{"error": "user-id header is required"})
		return
	}
	tvShows, presence, err := calendarService.GetTvShowCalendar(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, toTVEpisodesResponse(tvShows, presence))
}
