package controllers

import (
	"github.com/bingemate/media-go-pkg/repository"
	"github.com/bingemate/media-go-pkg/tmdb"
	"time"
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

type tvShowResponse struct {
	ID           int                `json:"id"`
	Actors       []person           `json:"actors"`
	BackdropURL  string             `json:"backdrop_url"`
	Crew         []person           `json:"crew"`
	Genres       []genre            `json:"genres"`
	Overview     string             `json:"overview"`
	PosterURL    string             `json:"poster_url"`
	ReleaseDate  string             `json:"release_date"`
	Studios      []studio           `json:"studios"`
	Status       string             `json:"status"`
	NextEpisode  *tvEpisodeResponse `json:"next_episode"`
	Title        string             `json:"title"`
	SeasonsCount int                `json:"seasons_count"`
	VoteAverage  float32            `json:"vote_average"`
	VoteCount    int                `json:"vote_count"`
}

type tvEpisodeResponse struct {
	ID            int    `json:"id"`
	TVShowID      int    `json:"tv_show_id"`
	PosterURL     string `json:"poster_url"`
	EpisodeNumber int    `json:"episode_number"`
	SeasonNumber  int    `json:"season_number"`
	Name          string `json:"name"`
	Overview      string `json:"overview"`
	AirDate       string `json:"air_date"`
}

type mediaResponse struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	MediaType   string    `json:"media_type"`
	TmdbID      int       `json:"tmdb_id"`
	ReleaseDate time.Time `json:"release_date"`
}

type mediaFileResponse struct {
	ID        string             `json:"id"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	Filename  string             `json:"filename"`
	Size      float64            `json:"size"`
	Duration  float64            `json:"duration"`
	MimeType  string             `json:"mime_type"`
	Codec     string             `json:"codec"`
	Audios    []audioResponse    `json:"audios"`
	Subtitles []subtitleResponse `json:"subtitles"`
}

type audioResponse struct {
	Codec    string  `json:"codec"`
	Language string  `json:"language"`
	Bitrate  float64 `json:"bitrate"`
}

type subtitleResponse struct {
	Codec    string `json:"code"`
	Language string `json:"language"`
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
		Crew: func() []person {
			var crew = make([]person, len(tvShow.Crew))
			for i, crewP := range tvShow.Crew {
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
		Studios: func() []studio {
			var studios = make([]studio, len(tvShow.Studios))
			for i, studioP := range tvShow.Studios {
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

func toTVEpisodesResponse(tvEpisodes *[]tmdb.TVEpisode) *[]tvEpisodeResponse {
	var tvEpisodesResponse = make([]tvEpisodeResponse, len(*tvEpisodes))
	for i, tvEpisode := range *tvEpisodes {
		var episode = tvEpisode
		tvEpisodesResponse[i] = *toTVEpisodeResponse(&episode)
	}
	return &tvEpisodesResponse
}

func toMediaResponse(media *repository.Media) *mediaResponse {
	return &mediaResponse{
		ID:          media.ID,
		CreatedAt:   media.CreatedAt,
		UpdatedAt:   media.UpdatedAt,
		MediaType:   string(media.MediaType),
		TmdbID:      media.TmdbID,
		ReleaseDate: media.ReleaseDate,
	}
}

func toMediaFileResponse(mediaFile *repository.MediaFile) *mediaFileResponse {
	return &mediaFileResponse{
		ID:        mediaFile.ID,
		CreatedAt: mediaFile.CreatedAt,
		UpdatedAt: mediaFile.UpdatedAt,
		Filename:  mediaFile.Filename,
		Size:      mediaFile.Size,
		Duration:  mediaFile.Duration,
		MimeType:  mediaFile.Mimetype,
		Codec:     string(mediaFile.Codec),
		Audios: func() []audioResponse {
			var audios = make([]audioResponse, len(mediaFile.Audio))
			for i, audio := range mediaFile.Audio {
				audios[i] = audioResponse{
					Codec:    string(audio.Codec),
					Language: audio.Language,
					Bitrate:  audio.Bitrate,
				}
			}
			return audios
		}(),
		Subtitles: func() []subtitleResponse {
			var subtitles = make([]subtitleResponse, len(mediaFile.Subtitles))
			for i, subtitle := range mediaFile.Subtitles {
				subtitles[i] = subtitleResponse{
					Codec:    string(subtitle.Codec),
					Language: subtitle.Language,
				}
			}
			return subtitles
		}(),
	}
}
