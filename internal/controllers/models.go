package controllers

import (
	"github.com/bingemate/media-go-pkg/repository"
	"github.com/bingemate/media-go-pkg/tmdb"
	"time"
)

type errorResponse struct {
	Error string `json:"error" example:"error message"`
}

type genre struct {
	ID   int    `json:"id" example:"35"`
	Name string `json:"name" example:"Comédie"`
}

type person struct {
	ID         int    `json:"id" example:"3292"`
	Character  string `json:"character" example:"R.M. Renfield"`
	Name       string `json:"name" example:"Nicholas Hoult"`
	ProfileURL string `json:"profileUrl" example:"https://image.tmdb.org/t/p/original/rbyi6sOw0dGV3wJzKXDopm2h0NO.jpg"`
}

type crew struct {
	ID         int    `json:"id" example:"24310"`
	Role       string `json:"role" example:"Director"`
	Name       string `json:"name" example:"Mitchell Amundsen"`
	ProfileURL string `json:"profileUrl" example:"https://image.tmdb.org/t/p/original/zK4o3lfe6chC2qDHhxpuO1RJl2X.jpg"`
}

type studio struct {
	ID      int    `json:"id" example:"33"`
	Name    string `json:"name" example:"Universal Pictures"`
	LogoURL string `json:"logoUrl" example:"https://image.tmdb.org/t/p/original/8lvHyhjr8oUKOOy2dKXoALWKdp0.png"`
}

type movieResponse struct {
	ID          int      `json:"id" example:"649609"`
	Actors      []person `json:"actors"`
	BackdropURL string   `json:"backdropUrl" example:"https://image.tmdb.org/t/p/original/e7FzphKs5gzoghDotAEp2FeP46u.jpg"`
	Crew        []crew   `json:"crew"`
	Genres      []genre  `json:"genres"`
	Overview    string   `json:"overview" example:"Le mal ne saurait survivre une éternité sans un petit coup de pouce.\r Dans cette version moderne du mythe de Dracula..."`
	PosterURL   string   `json:"posterUrl" example:"https://image.tmdb.org/t/p/original/lm3y4RNPu4aRDePsX5CkB9ndEdQ.jpg"`
	ReleaseDate string   `json:"releaseDate" example:"2023-04-07"`
	Studios     []studio `json:"studios"`
	Title       string   `json:"title" example:"Renfield"`
	VoteAverage float32  `json:"voteAverage" example:"7.252"`
	VoteCount   int      `json:"voteCount" example:"278"`
}

type tvShowResponse struct {
	ID           int                `json:"id" example:"200777"`
	Actors       []person           `json:"actors"`
	BackdropURL  string             `json:"backdropUrl" example:"https://image.tmdb.org/t/p/original/oL459mgvcnc3jL90K7zkfvXQu0.jpg"`
	Crew         []crew             `json:"crew"`
	Genres       []genre            `json:"genres"`
	Overview     string             `json:"overview" example:"Ray White est un jeune homme venant d'entrer dans la populaire académie de magie Arnold. En tant que..."`
	PosterURL    string             `json:"posterUrl" example:"https://image.tmdb.org/t/p/original/aiJd0oGkBhf98uEH3F3yC7O48vr.jpg"`
	ReleaseDate  string             `json:"releaseDate" example:"2023-01-06"`
	Networks     []studio           `json:"networks"`
	Status       string             `json:"status" example:"Ended"`
	NextEpisode  *tvEpisodeResponse `json:"nextEpisode"`
	Title        string             `json:"title" example:"The Iceblade Sorcerer Shall Rule the World"`
	SeasonsCount int                `json:"seasonsCount" example:"1"`
	VoteAverage  float32            `json:"voteAverage" example:"6.7"`
	VoteCount    int                `json:"voteCount" example:"11"`
}

type tvEpisodeResponse struct {
	ID            int    `json:"id" example:"4137463"`
	TVShowID      int    `json:"tvShowId" example:"200777"`
	PosterURL     string `json:"posterUrl" example:"https://image.tmdb.org/t/p/original/uVqsuh8qrNX8tkQDpDF7nDZdg0w.jpg"`
	EpisodeNumber int    `json:"episodeNumber" example:"12"`
	SeasonNumber  int    `json:"seasonNumber" example:"1"`
	Name          string `json:"name" example:"Le plus puissant sorcier du monde révèle Akasha"`
	Overview      string `json:"overview" example:"Ray et ses amis volent au secours de Rebecca, qui montre des signes..."`
	AirDate       string `json:"airDate" example:"2023-03-24"`
}

type mediaResponse struct {
	ID          int       `json:"id" example:"134564"`
	CreatedAt   time.Time `json:"createdAt" example:"2023-05-07T20:31:28.327382+02:00"`
	UpdatedAt   time.Time `json:"updatedAt" example:"2023-05-07T20:31:28.327382+02:00"`
	MediaType   string    `json:"mediaType" example:"TvShow"`
	ReleaseDate string    `json:"releaseDate" example:"2023-01-06"`
}
type mediaFileResponse struct {
	ID        string             `json:"id" example:"eec1d6b7-97c9-47e9-846b-6817d0e3d4ed"`
	CreatedAt time.Time          `json:"createdAt" example:"2023-05-07T20:31:28.327382+02:00"`
	UpdatedAt time.Time          `json:"updatedAt" example:"2023-05-07T20:31:28.327382+02:00"`
	Filename  string             `json:"filename" example:"The Iceblade Sorcerer Shall Rule the World - S1E09.mkv"`
	Duration  float64            `json:"duration" example:"1450.76"`
	Audios    []audioResponse    `json:"audios"`
	Subtitles []subtitleResponse `json:"subtitles"`
}

type audioResponse struct {
	Language string `json:"language" example:"jpn"`
	Filename string `json:"filename" example:"audio_1.m3u8"`
}

type subtitleResponse struct {
	Language string `json:"language" example:"fre"`
	Filename string `json:"filename" example:"subtitle_1.vtt"`
}

type tvShowResults struct {
	Results     []*tvShowResponse `json:"results"`
	TotalPage   int               `json:"totalPage" example:"71"`
	TotalResult int               `json:"totalResult" example:"1412"`
}

type movieResults struct {
	Results     []*movieResponse `json:"results"`
	TotalPage   int              `json:"totalPage" example:"71"`
	TotalResult int              `json:"totalResult" example:"1412"`
}

func toMovieResponse(movie *tmdb.Movie) *movieResponse {
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
		Crew: func() []crew {
			var crews = make([]crew, len(movie.Crew))
			for i, crewP := range movie.Crew {
				crews[i] = crew{
					ID:         crewP.ID,
					Role:       crewP.Character,
					Name:       crewP.Name,
					ProfileURL: crewP.ProfileURL,
				}
			}
			return crews
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

func toMoviesResponse(movies []*tmdb.Movie) []*movieResponse {
	var moviesResponse = make([]*movieResponse, len(movies))
	for i, movie := range movies {
		moviesResponse[i] = toMovieResponse(movie)
	}
	return moviesResponse
}

func toTVShowResponse(tvShow *tmdb.TVShow) *tvShowResponse {
	return &tvShowResponse{
		ID: tvShow.ID,
		Actors: func() []person {
			var actors = make([]person, len(tvShow.Actors))
			for i, actor := range tvShow.Actors {
				actors[i] = person{
					ID:         actor.ID,
					Character:  actor.Character,
					Name:       actor.Name,
					ProfileURL: actor.ProfileURL,
				}
			}
			return actors
		}(),
		BackdropURL: tvShow.BackdropURL,
		Crew: func() []crew {
			var crews = make([]crew, len(tvShow.Crew))
			for i, crewP := range tvShow.Crew {
				crews[i] = crew{
					ID:         crewP.ID,
					Role:       crewP.Character,
					Name:       crewP.Name,
					ProfileURL: crewP.ProfileURL,
				}
			}
			return crews
		}(),
		Genres: func() []genre {
			var genres = make([]genre, len(tvShow.Genres))
			for i, genreP := range tvShow.Genres {
				genres[i] = genre{
					ID:   genreP.ID,
					Name: genreP.Name,
				}
			}
			return genres
		}(),
		Overview:    tvShow.Overview,
		PosterURL:   tvShow.PosterURL,
		ReleaseDate: tvShow.ReleaseDate,
		Networks: func() []studio {
			var studios = make([]studio, len(tvShow.Networks))
			for i, studioP := range tvShow.Networks {
				studios[i] = studio{
					ID:      studioP.ID,
					Name:    studioP.Name,
					LogoURL: studioP.LogoURL,
				}
			}
			return studios
		}(),
		Title: tvShow.Title,
		NextEpisode: func() *tvEpisodeResponse {
			if tvShow.NextEpisode == nil {
				return nil
			}
			return toTVEpisodeResponse(tvShow.NextEpisode)
		}(),
		SeasonsCount: tvShow.SeasonsCount,
		Status:       tvShow.Status,
		VoteAverage:  tvShow.VoteAverage,
		VoteCount:    tvShow.VoteCount,
	}
}

func toTVShowsResponse(tvShows []*tmdb.TVShow) []*tvShowResponse {
	var tvShowsResponse = make([]*tvShowResponse, len(tvShows))
	for i, tvShow := range tvShows {
		tvShowsResponse[i] = toTVShowResponse(tvShow)
	}
	return tvShowsResponse
}

func toTVEpisodeResponse(tvEpisode *tmdb.TVEpisode) *tvEpisodeResponse {
	return &tvEpisodeResponse{
		ID:            tvEpisode.ID,
		TVShowID:      tvEpisode.TVShowID,
		PosterURL:     tvEpisode.PosterURL,
		EpisodeNumber: tvEpisode.EpisodeNumber,
		SeasonNumber:  tvEpisode.SeasonNumber,
		Name:          tvEpisode.Name,
		Overview:      tvEpisode.Overview,
		AirDate:       tvEpisode.AirDate,
	}
}

func toTVEpisodesResponse(tvEpisodes []*tmdb.TVEpisode) []*tvEpisodeResponse {
	var tvEpisodesResponse = make([]*tvEpisodeResponse, len(tvEpisodes))
	for i, tvEpisode := range tvEpisodes {
		tvEpisodesResponse[i] = toTVEpisodeResponse(tvEpisode)
	}
	return tvEpisodesResponse
}

func toMediaResponse(media *repository.Media) *mediaResponse {
	return &mediaResponse{
		ID:          media.ID,
		CreatedAt:   media.CreatedAt,
		UpdatedAt:   media.UpdatedAt,
		MediaType:   string(media.MediaType),
		ReleaseDate: media.ReleaseDate.Format("2006-01-02"),
	}
}

func toMediaFileResponse(mediaFile *repository.MediaFile) *mediaFileResponse {
	return &mediaFileResponse{
		ID:        mediaFile.ID,
		CreatedAt: mediaFile.CreatedAt,
		UpdatedAt: mediaFile.UpdatedAt,
		Filename:  mediaFile.Filename,
		Duration:  mediaFile.Duration,
		Audios: func() []audioResponse {
			var audios = make([]audioResponse, len(mediaFile.Audio))
			for i, audio := range mediaFile.Audio {
				audios[i] = audioResponse{
					Filename: audio.Filename,
					Language: audio.Language,
				}
			}
			return audios
		}(),
		Subtitles: func() []subtitleResponse {
			var subtitles = make([]subtitleResponse, len(mediaFile.Subtitles))
			for i, subtitle := range mediaFile.Subtitles {
				subtitles[i] = subtitleResponse{
					Filename: subtitle.Filename,
					Language: subtitle.Language,
				}
			}
			return subtitles
		}(),
	}
}
