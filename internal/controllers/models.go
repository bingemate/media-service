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
