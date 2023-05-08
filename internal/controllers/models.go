package controllers

import (
	"github.com/bingemate/media-go-pkg/repository"
	"github.com/bingemate/media-go-pkg/tmdb"
	"time"
)

type errorResponse struct {
	Error string `json:"error" example:"media not found"`
}

type genre struct {
	ID   int    `json:"id" example:"35"`
	Name string `json:"name" example:"Comédie"`
}

type person struct {
	ID         int    `json:"id" example:"3292"`
	Character  string `json:"character" example:"R.M. Renfield"`
	Name       string `json:"name" example:"Nicholas Hoult"`
	ProfileURL string `json:"profile_url" example:"https://image.tmdb.org/t/p/original/rbyi6sOw0dGV3wJzKXDopm2h0NO.jpg"`
}

type crew struct {
	ID         int    `json:"id" example:"24310"`
	Role       string `json:"role" example:"Director"`
	Name       string `json:"name" example:"Mitchell Amundsen"`
	ProfileURL string `json:"profile_url" example:"https://image.tmdb.org/t/p/original/zK4o3lfe6chC2qDHhxpuO1RJl2X.jpg"`
}

type studio struct {
	ID      int    `json:"id" example:"33"`
	Name    string `json:"name" example:"Universal Pictures"`
	LogoURL string `json:"logo_url" example:"https://image.tmdb.org/t/p/original/8lvHyhjr8oUKOOy2dKXoALWKdp0.png"`
}

type movieResponse struct {
	ID          int      `json:"id" example:"649609"`
	Actors      []person `json:"actors"`
	BackdropURL string   `json:"backdrop_url" example:"https://image.tmdb.org/t/p/original/e7FzphKs5gzoghDotAEp2FeP46u.jpg"`
	Crew        []crew   `json:"crew"`
	Genres      []genre  `json:"genres"`
	Overview    string   `json:"overview" example:"Le mal ne saurait survivre une éternité sans un petit coup de pouce.\r Dans cette version moderne du mythe de Dracula..."`
	PosterURL   string   `json:"poster_url" example:"https://image.tmdb.org/t/p/original/lm3y4RNPu4aRDePsX5CkB9ndEdQ.jpg"`
	ReleaseDate string   `json:"release_date" example:"2023-04-07"`
	Studios     []studio `json:"studios"`
	Title       string   `json:"title" example:"Renfield"`
	VoteAverage float32  `json:"vote_average" example:"7.252"`
	VoteCount   int      `json:"vote_count" example:"278"`
}

type tvShowResponse struct {
	ID           int                `json:"id" example:"200777"`
	Actors       []person           `json:"actors"`
	BackdropURL  string             `json:"backdrop_url" example:"https://image.tmdb.org/t/p/original/oL459mgvcnc3jL90K7zkfvXQu0.jpg"`
	Crew         []crew             `json:"crew"`
	Genres       []genre            `json:"genres"`
	Overview     string             `json:"overview" example:"Ray White est un jeune homme venant d'entrer dans la populaire académie de magie Arnold. En tant que..."`
	PosterURL    string             `json:"poster_url" example:"https://image.tmdb.org/t/p/original/aiJd0oGkBhf98uEH3F3yC7O48vr.jpg"`
	ReleaseDate  string             `json:"release_date" example:"2023-01-06"`
	Studios      []studio           `json:"studios"`
	Status       string             `json:"status" example:"Ended"`
	NextEpisode  *tvEpisodeResponse `json:"next_episode"`
	Title        string             `json:"title" example:"The Iceblade Sorcerer Shall Rule the World"`
	SeasonsCount int                `json:"seasons_count" example:"1"`
	VoteAverage  float32            `json:"vote_average" example:"6.7"`
	VoteCount    int                `json:"vote_count" example:"11"`
}

type tvEpisodeResponse struct {
	ID            int    `json:"id" example:"4137463"`
	TVShowID      int    `json:"tv_show_id" example:"200777"`
	PosterURL     string `json:"poster_url" example:"https://image.tmdb.org/t/p/original/uVqsuh8qrNX8tkQDpDF7nDZdg0w.jpg"`
	EpisodeNumber int    `json:"episode_number" example:"12"`
	SeasonNumber  int    `json:"season_number" example:"1"`
	Name          string `json:"name" example:"Le plus puissant sorcier du monde révèle Akasha"`
	Overview      string `json:"overview" example:"Ray et ses amis volent au secours de Rebecca, qui montre des signes..."`
	AirDate       string `json:"air_date" example:"2023-03-24"`
}

type mediaResponse struct {
	ID          string    `json:"id" example:"eec1d6b7-97c9-47e9-846b-6817d0e3d4ed"`
	CreatedAt   time.Time `json:"created_at" example:"2023-05-07T20:31:28.327382+02:00"`
	UpdatedAt   time.Time `json:"updated_at" example:"2023-05-07T20:31:28.327382+02:00"`
	MediaType   string    `json:"media_type" example:"TvShow"`
	TmdbID      int       `json:"tmdb_id" example:"200777"`
	ReleaseDate string    `json:"release_date" example:"2023-01-06"`
}
type mediaFileResponse struct {
	ID        string             `json:"id" example:"eec1d6b7-97c9-47e9-846b-6817d0e3d4ed"`
	CreatedAt time.Time          `json:"created_at" example:"2023-05-07T20:31:28.327382+02:00"`
	UpdatedAt time.Time          `json:"updated_at" example:"2023-05-07T20:31:28.327382+02:00"`
	Filename  string             `json:"filename" example:"The Iceblade Sorcerer Shall Rule the World - S1E09.mkv"`
	Size      float64            `json:"size" example:"771367779"`
	Duration  float64            `json:"duration" example:"1450.76"`
	MimeType  string             `json:"mime_type" example:"video/x-matroska"`
	Codec     string             `json:"codec" example:"H264"`
	Audios    []audioResponse    `json:"audios"`
	Subtitles []subtitleResponse `json:"subtitles"`
}

type audioResponse struct {
	Codec    string  `json:"codec" example:"AAC"`
	Language string  `json:"language" example:"jpn"`
	Bitrate  float64 `json:"bitrate" example:"160"`
}

type subtitleResponse struct {
	Codec    string `json:"code" example:"ASS"`
	Language string `json:"language" example:"fre"`
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
		ReleaseDate: media.ReleaseDate.Format("2006-01-02"),
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
