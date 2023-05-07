package controllers

import (
	"github.com/bingemate/media-go-pkg/tmdb"
	"github.com/bingemate/media-service/internal/features"
	"github.com/gin-gonic/gin"
	"strconv"
)

type genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type person struct {
	ID         int    `json:"id"`
	Character  string `json:"character"`
	Name       string `json:"name"`
	ProfileURL string `json:"profile_url"`
}

type studio struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	LogoURL string `json:"logo_url"`
}

type movieResponse struct {
	ID          int      `json:"id"`
	Actors      []person `json:"actors"`
	BackdropURL string   `json:"backdrop_url"`
	Crew        []person `json:"crew"`
	Genres      []genre  `json:"genres"`
	Overview    string   `json:"overview"`
	PosterURL   string   `json:"poster_url"`
	ReleaseDate string   `json:"release_date"`
	Studios     []studio `json:"studios"`
	Title       string   `json:"title"`
	VoteAverage float32  `json:"vote_average"`
	VoteCount   int      `json:"vote_count"`
}

func toMovieResponse(movie tmdb.Movie) *movieResponse {
	return &movieResponse{
		ID: movie.ID,
		Actors: func() []person {
			var actors = make([]person, len(movie.Actors))
			for i, actor := range movie.Actors {
				actors[i] = person{
					ID:         actor.ID,
					Character:  actor.Character,
					Name:       actor.Name,
					ProfileURL: actor.ProfileURL,
				}
			}
			return actors
		}(),
		BackdropURL: movie.BackdropURL,
		Crew: func() []person {
			var crew = make([]person, len(movie.Crew))
			for i, crewP := range movie.Crew {
				crew[i] = person{
					ID:         crewP.ID,
					Character:  crewP.Character,
					Name:       crewP.Name,
					ProfileURL: crewP.ProfileURL,
				}
			}
			return crew
		}(),
		Genres: func() []genre {
			var genres = make([]genre, len(movie.Genres))
			for i, genreP := range movie.Genres {
				genres[i] = genre{
					ID:   genreP.ID,
					Name: genreP.Name,
				}
			}
			return genres
		}(),
		Overview:    movie.Overview,
		PosterURL:   movie.PosterURL,
		ReleaseDate: movie.ReleaseDate,
		Studios: func() []studio {
			var studios = make([]studio, len(movie.Studios))
			for i, studioP := range movie.Studios {
				studios[i] = studio{
					ID:      studioP.ID,
					Name:    studioP.Name,
					LogoURL: studioP.LogoURL,
				}
			}
			return studios
		}(),
		Title:       movie.Title,
		VoteAverage: movie.VoteAverage,
		VoteCount:   movie.VoteCount,
	}
}

func InitMediaDataController(engine *gin.RouterGroup, mediaData *features.MediaData) {
	engine.GET("/movie/:id", func(c *gin.Context) {
		getMovie(c, mediaData)
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
// @Router			/media/movie/{id} [get]
func getMovie(c *gin.Context, mediaData *features.MediaData) {
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
	c.JSON(200, toMovieResponse(*result))
}
