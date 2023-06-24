package controllers

import (
	"fmt"
	ics "github.com/arran4/golang-ical"
	"github.com/bingemate/media-go-pkg/tmdb"
	"github.com/bingemate/media-service/internal/features"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func InitCalendarController(engine *gin.RouterGroup, calendarService *features.CalendarService) {
	engine.GET("/movies", func(c *gin.Context) {
		getMoviesCalendar(c, calendarService)
	})
	engine.GET("/tvshows", func(c *gin.Context) {
		getTvShowsCalendar(c, calendarService)
	})
	engine.GET("/movies/ical/:user-id", func(c *gin.Context) {
		getMoviesCalendarIcal(c, calendarService)
	})
	engine.GET("/tvshows/ical/:user-id", func(c *gin.Context) {
		getTvShowsCalendarIcal(c, calendarService)
	})
}

// @Summary Get movies calendar
// @Description Get movies calendar
// @Tags  Calendar
// @Tags Movie
// @Param month query int true "Month"
// @Param year query int false "Year"
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
	month, err := strconv.Atoi(c.Query("month"))
	if err != nil || month < 1 || month > 12 {
		c.JSON(400, gin.H{"error": "month query param is required"})
		return
	}
	year, err := strconv.Atoi(c.Query("year"))
	if err != nil {
		year = time.Now().Year()
	}

	movies, presence, err := calendarService.GetMoviesCalendar(userID, month, year)
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
// @Param month query int true "Month"
// @Param year query int false "Year"
// @Produce  json
// @Success 200 {object} tvReleasesResults
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /calendar/tvshows [get]
func getTvShowsCalendar(c *gin.Context, calendarService *features.CalendarService) {
	userID := c.GetHeader("user-id")
	if userID == "" {
		c.JSON(400, gin.H{"error": "user-id header is required"})
		return
	}
	month, err := strconv.Atoi(c.Query("month"))
	if err != nil || month < 1 || month > 12 {
		c.JSON(400, gin.H{"error": "month query param is required"})
		return
	}
	year, err := strconv.Atoi(c.Query("year"))
	if err != nil {
		year = time.Now().Year()
	}
	episodes, tvShows, presence, err := calendarService.GetTvShowCalendar(userID, month, year)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, toTVReleasesResult(episodes, tvShows, presence))
}

// @Summary Get movies calendar in iCal format
// @Description Get movies calendar in iCal format
// @Tags  Calendar
// @Tags Movie
// @Param user-id path string true "User ID"
// @Produce text/calendar
// @Success 200 {string} string "iCal data"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /calendar/movies/ical/{user-id} [get]
func getMoviesCalendarIcal(c *gin.Context, calendarService *features.CalendarService) {
	userID := c.Param("user-id")
	if userID == "" {
		c.JSON(400, gin.H{"error": "user-id header is required"})
		return
	}
	/*month := int(time.Now().Month())
	nextMonth := (month)%12 + 1
	var movies []*tmdb.Movie
	var lock sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		results, _, err := calendarService.GetMoviesCalendar(userID, month)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		lock.Lock()
		defer lock.Unlock()
		movies = append(movies, results...)
	}()
	go func() {
		defer wg.Done()
		results, _, err := calendarService.GetMoviesCalendar(userID, nextMonth)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		lock.Lock()
		defer lock.Unlock()
		movies = append(movies, results...)
	}()

	wg.Wait()*/
	now := time.Now()
	sixMonthAgo := now.AddDate(0, -6, 0)
	sixMonthLater := now.AddDate(0, 6, 0)
	movies, _, err := calendarService.GetMoviesCalendarInRange(userID, sixMonthAgo, sixMonthLater)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)
	for _, movie := range movies {
		releaseDate, _ := time.Parse("2006-01-02", movie.ReleaseDate)
		event := cal.AddEvent(strconv.Itoa(movie.ID))
		event.SetSummary(movie.Title)
		event.SetDescription(movie.Overview)
		event.SetCreatedTime(time.Now())
		event.SetDtStampTime(time.Now())
		event.SetModifiedAt(time.Now())
		event.SetStartAt(releaseDate)
		event.SetEndAt(releaseDate.Add(time.Minute * 30))
	}

	c.Header("Content-Type", "text/calendar")
	c.String(200, cal.Serialize())
}

// @Summary Get tv shows calendar in iCal format
// @Description Get tv shows calendar in iCal format
// @Tags  Calendar
// @Tags TvShow
// @Param user-id path string true "User ID"
// @Produce text/calendar
// @Success 200 {string} string "iCal data"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /calendar/tvshows/ical/{user-id} [get]
func getTvShowsCalendarIcal(c *gin.Context, calendarService *features.CalendarService) {
	userID := c.Param("user-id")
	if userID == "" {
		c.JSON(400, gin.H{"error": "user-id header is required"})
		return
	}
	/*month := int(time.Now().Month())
	nextMonth := (month)%12 + 1
	var episodes []*tmdb.TVEpisode
	var tvShowsMap = make(map[int]*tmdb.TVShow)
	var lock sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		results, tvShows, _, err := calendarService.GetTvShowCalendar(userID, month)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		lock.Lock()
		defer lock.Unlock()
		episodes = append(episodes, results...)
		for _, tvShow := range tvShows {
			tvShowsMap[tvShow.ID] = tvShow
		}
	}()
	go func() {
		defer wg.Done()
		results, tvShows, _, err := calendarService.GetTvShowCalendar(userID, nextMonth)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		lock.Lock()
		defer lock.Unlock()
		episodes = append(episodes, results...)
		for _, tvShow := range tvShows {
			tvShowsMap[tvShow.ID] = tvShow
		}
	}()

	wg.Wait()*/
	var tvShowsMap = make(map[int]*tmdb.TVShow)
	now := time.Now()
	sixMonthAgo := now.AddDate(0, -6, 0)
	sixMonthLater := now.AddDate(0, 6, 0)
	episodes, tvShows, _, err := calendarService.GetTvShowCalendarInRange(userID, sixMonthAgo, sixMonthLater)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	for _, tvShow := range tvShows {
		tvShowsMap[tvShow.ID] = tvShow
	}

	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)
	for _, episode := range episodes {
		releaseDate, _ := time.Parse("2006-01-02", episode.AirDate)
		event := cal.AddEvent(strconv.Itoa(episode.ID))
		event.SetSummary(fmt.Sprintf("%s - %dx%02d - %s",
			tvShowsMap[episode.TVShowID].Title,
			episode.SeasonNumber,
			episode.EpisodeNumber,
			episode.Name,
		))
		event.SetDescription(episode.Overview)
		event.SetCreatedTime(time.Now())
		event.SetDtStampTime(time.Now())
		event.SetModifiedAt(time.Now())
		event.SetStartAt(releaseDate)
		event.SetEndAt(releaseDate.Add(time.Minute * 30))
	}

	c.Header("Content-Type", "text/calendar")
	c.String(200, cal.Serialize())
}
