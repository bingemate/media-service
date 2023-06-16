package controllers

import (
	"github.com/bingemate/media-go-pkg/repository"
	"github.com/bingemate/media-go-pkg/tmdb"
	"sort"
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

type actor struct {
	ID         int    `json:"id" example:"3292"`
	Name       string `json:"name" example:"Nicholas Hoult"`
	ProfileURL string `json:"profileUrl" example:"https://image.tmdb.org/t/p/original/rbyi6sOw0dGV3wJzKXDopm2h0NO.jpg"`
	Overview   string `json:"overview" example:"Nicholas Caradoc Hoult (born 7 December 1989) is an English actor. His body of work includes supporting work in big-budget mainstream productions and starring roles in independent projects in both the American and the British film industries. He has been nominated for awards such as a British Academy Film Award and a Critics Choice Award for his work."`
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
	Present     bool     `json:"present" example:"true"`
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
	ID            int                `json:"id" example:"200777"`
	Present       bool               `json:"present" example:"true"`
	Actors        []person           `json:"actors"`
	BackdropURL   string             `json:"backdropUrl" example:"https://image.tmdb.org/t/p/original/oL459mgvcnc3jL90K7zkfvXQu0.jpg"`
	Crew          []crew             `json:"crew"`
	Genres        []genre            `json:"genres"`
	Overview      string             `json:"overview" example:"Ray White est un jeune homme venant d'entrer dans la populaire académie de magie Arnold. En tant que..."`
	PosterURL     string             `json:"posterUrl" example:"https://image.tmdb.org/t/p/original/aiJd0oGkBhf98uEH3F3yC7O48vr.jpg"`
	ReleaseDate   string             `json:"releaseDate" example:"2023-01-06"`
	Networks      []studio           `json:"networks"`
	Status        string             `json:"status" example:"Ended"`
	NextEpisode   *tvEpisodeResponse `json:"nextEpisode"`
	Title         string             `json:"title" example:"The Iceblade Sorcerer Shall Rule the World"`
	SeasonsCount  int                `json:"seasonsCount" example:"1"`
	EpisodesCount int                `json:"episodesCount" example:"12"`
	VoteAverage   float32            `json:"voteAverage" example:"6.7"`
	VoteCount     int                `json:"voteCount" example:"11"`
}

type tvEpisodeResponse struct {
	ID            int    `json:"id" example:"4137463"`
	Present       bool   `json:"present" example:"true"`
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
	Name        string    `json:"name" example:"The Iceblade Sorcerer Shall Rule the World"`
	ReleaseDate string    `json:"releaseDate" example:"2023-01-06"`
}
type episodeMediaResponse struct {
	ID            int       `json:"id" example:"4137463"`
	CreatedAt     time.Time `json:"createdAt" example:"2023-05-07T20:31:28.327382+02:00"`
	UpdatedAt     time.Time `json:"updatedAt" example:"2023-05-07T20:31:28.327382+02:00"`
	Name          string    `json:"name" example:"Le plus puissant sorcier du monde révèle Akasha"`
	ReleaseDate   string    `json:"releaseDate" example:"2023-03-24"`
	EpisodeNumber int       `json:"episodeNumber" example:"12"`
	SeasonNumber  int       `json:"seasonNumber" example:"1"`
}

type episodeFileResponse struct {
	ID            int                `json:"id" example:"4137463"`
	Name          string             `json:"name" example:"Le plus puissant sorcier du monde révèle Akasha"`
	ReleaseDate   string             `json:"releaseDate" example:"2023-03-24"`
	EpisodeNumber int                `json:"episodeNumber" example:"12"`
	SeasonNumber  int                `json:"seasonNumber" example:"1"`
	TvShowId      int                `json:"tvShowId" example:"200777"`
	TvShowName    string             `json:"tvShowName" example:"The Iceblade Sorcerer Shall Rule the World"`
	File          *mediaFileResponse `json:"file"`
}

type episodeFilesResult struct {
	Results []*episodeFileResponse `json:"results"`
	Total   int                    `json:"total"`
}

type movieFileResponse struct {
	ID          int                `json:"id" example:"134564"`
	Name        string             `json:"name" example:"Apex"`
	ReleaseDate string             `json:"releaseDate" example:"2023-01-06"`
	File        *mediaFileResponse `json:"file"`
}

type movieFilesResult struct {
	Results []*movieFileResponse `json:"results"`
	Total   int                  `json:"total"`
}

type mediaFileResponse struct {
	ID        string             `json:"id" example:"eec1d6b7-97c9-47e9-846b-6817d0e3d4ed"`
	CreatedAt time.Time          `json:"createdAt" example:"2023-05-07T20:31:28.327382+02:00"`
	UpdatedAt time.Time          `json:"updatedAt" example:"2023-05-07T20:31:28.327382+02:00"`
	Size      int64              `json:"size" example:"123456789"`
	Filename  string             `json:"filename" example:"index.m3u8"`
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

type commentRequest struct {
	Content string `json:"content" example:"This is a comment"`
}

type commentResponse struct {
	ID        string    `json:"id" example:"eec1d6b7-97c9-47e9-846b-6817d0e3d4ed"`
	CreatedAt time.Time `json:"createdAt" example:"2023-05-07T20:31:28.327382+02:00"`
	UpdatedAt time.Time `json:"updatedAt" example:"2023-05-07T20:31:28.327382+02:00"`
	Content   string    `json:"content" example:"This is a comment"`
	MediaID   int       `json:"mediaId" example:"134564"`
	UserID    string    `json:"userId" example:"eec1d6b7-97c9-47e9-846b-6817d0e3d4ed"`
}

type commentResults struct {
	Results     []*commentResponse `json:"results"`
	TotalResult int                `json:"totalResult" example:"1412"`
}

type commentHistoryReponse struct {
	Date  string `json:"date" example:"2023-05-07"`
	Count int    `json:"count" example:"12"`
}

type ratingRequest struct {
	Rating int `json:"rating" example:"5"`
}

type ratingResponse struct {
	UserID    string    `json:"userId" example:"eec1d6b7-97c9-47e9-846b-6817d0e3d4ed"`
	MediaID   int       `json:"mediaId" example:"134564"`
	Rating    int       `json:"rating" example:"5"`
	CreatedAt time.Time `json:"createdAt" example:"2023-05-07T20:31:28.327382+02:00"`
	UpdatedAt time.Time `json:"updatedAt" example:"2023-05-07T20:31:28.327382+02:00"`
}

type ratingResults struct {
	Results     []*ratingResponse `json:"results"`
	TotalResult int               `json:"totalResult" example:"14"`
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
type tvReleasesResults struct {
	Episodes []*tvEpisodeResponse `json:"episodes"`
	TvShows  []*tvShowResponse    `json:"tvShows"`
}

func toMovieResponse(movie *tmdb.Movie, present bool) *movieResponse {
	return &movieResponse{
		ID:      movie.ID,
		Present: present,
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

func toMoviesResponse(movies []*tmdb.Movie, presence *[]bool) []*movieResponse {
	var moviesResponse = make([]*movieResponse, len(movies))
	for i, movie := range movies {
		moviesResponse[i] = toMovieResponse(movie, (*presence)[i])
	}
	return moviesResponse
}

func toTVShowResponse(tvShow *tmdb.TVShow, present bool) *tvShowResponse {
	return &tvShowResponse{
		ID:      tvShow.ID,
		Present: present,
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
			return toTVEpisodeResponse(tvShow.NextEpisode, false)
		}(),
		SeasonsCount:  tvShow.SeasonsCount,
		EpisodesCount: tvShow.EpisodesCount,
		Status:        tvShow.Status,
		VoteAverage:   tvShow.VoteAverage,
		VoteCount:     tvShow.VoteCount,
	}
}

func toTVShowsResponse(tvShows []*tmdb.TVShow, presence *[]bool) []*tvShowResponse {
	var tvShowsResponse = make([]*tvShowResponse, len(tvShows))
	for i, tvShow := range tvShows {
		tvShowsResponse[i] = toTVShowResponse(tvShow, (*presence)[i])
	}
	return tvShowsResponse
}

func toTVEpisodeResponse(tvEpisode *tmdb.TVEpisode, present bool) *tvEpisodeResponse {
	return &tvEpisodeResponse{
		ID:            tvEpisode.ID,
		Present:       present,
		TVShowID:      tvEpisode.TVShowID,
		PosterURL:     tvEpisode.PosterURL,
		EpisodeNumber: tvEpisode.EpisodeNumber,
		SeasonNumber:  tvEpisode.SeasonNumber,
		Name:          tvEpisode.Name,
		Overview:      tvEpisode.Overview,
		AirDate:       tvEpisode.AirDate,
	}
}

func toTVEpisodesResponse(tvEpisodes []*tmdb.TVEpisode, presence *[]bool) []*tvEpisodeResponse {
	var tvEpisodesResponse = make([]*tvEpisodeResponse, len(tvEpisodes))
	for i, tvEpisode := range tvEpisodes {
		tvEpisodesResponse[i] = toTVEpisodeResponse(tvEpisode, (*presence)[i])
	}
	return tvEpisodesResponse
}

func toTVReleasesResult(tvEpisodes []*tmdb.TVEpisode, tvShows []*tmdb.TVShow, presence *[]bool) *tvReleasesResults {
	return &tvReleasesResults{
		Episodes: toTVEpisodesResponse(tvEpisodes, presence),
		TvShows:  toTVShowsResponse(tvShows, presence),
	}
}

func toMovieMediaResponse(media *repository.Movie) *mediaResponse {
	return &mediaResponse{
		ID:          media.ID,
		CreatedAt:   media.CreatedAt,
		UpdatedAt:   media.UpdatedAt,
		Name:        media.Name,
		ReleaseDate: media.ReleaseDate.Format("2006-01-02"),
	}
}

func toTVShowMediaResponse(media *repository.TvShow) *mediaResponse {
	return &mediaResponse{
		ID:          media.ID,
		CreatedAt:   media.CreatedAt,
		UpdatedAt:   media.UpdatedAt,
		Name:        media.Name,
		ReleaseDate: media.ReleaseDate.Format("2006-01-02"),
	}
}

func toEpisodeMediaResponse(media *repository.Episode) *episodeMediaResponse {
	return &episodeMediaResponse{
		ID:            media.ID,
		CreatedAt:     media.CreatedAt,
		UpdatedAt:     media.UpdatedAt,
		Name:          media.Name,
		ReleaseDate:   media.ReleaseDate.Format("2006-01-02"),
		EpisodeNumber: media.NbEpisode,
		SeasonNumber:  media.NbSeason,
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
		Audios: func() []audioResponse {
			var audios = make([]audioResponse, len(mediaFile.Audios))
			for i, audio := range mediaFile.Audios {
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

func toGenreResponse(tmdbGenre *tmdb.Genre) *genre {
	return &genre{
		ID:   tmdbGenre.ID,
		Name: tmdbGenre.Name,
	}
}

func toGenresResponse(tmdbGenres []*tmdb.Genre) []*genre {
	var genres = make([]*genre, len(tmdbGenres))
	for i, tmdbGenre := range tmdbGenres {
		genres[i] = toGenreResponse(tmdbGenre)
	}
	return genres
}

func toStudioResponse(tmdbStudio *tmdb.Studio) *studio {
	return &studio{
		ID:      tmdbStudio.ID,
		Name:    tmdbStudio.Name,
		LogoURL: tmdbStudio.LogoURL,
	}
}

func toActorResponse(tmdbActor *tmdb.Actor) *actor {
	return &actor{
		ID:         tmdbActor.ID,
		Name:       tmdbActor.Name,
		ProfileURL: tmdbActor.ProfileURL,
		Overview:   tmdbActor.Overview,
	}
}

func toMovieCommentResponse(comment *repository.MovieComment) *commentResponse {
	return &commentResponse{
		ID:        comment.ID,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		Content:   comment.Content,
		UserID:    comment.UserID,
		MediaID:   comment.MovieID,
	}
}

func toMovieCommentsResponse(comments []*repository.MovieComment) []*commentResponse {
	var commentsResponse = make([]*commentResponse, len(comments))
	for i, comment := range comments {
		commentsResponse[i] = toMovieCommentResponse(comment)
	}
	return commentsResponse
}

func toTVShowCommentResponse(comment *repository.TvShowComment) *commentResponse {
	return &commentResponse{
		ID:        comment.ID,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		Content:   comment.Content,
		UserID:    comment.UserID,
		MediaID:   comment.TvShowID,
	}
}

func toTVShowCommentsResponse(comments []*repository.TvShowComment) []*commentResponse {
	var commentsResponse = make([]*commentResponse, len(comments))
	for i, comment := range comments {
		commentsResponse[i] = toTVShowCommentResponse(comment)
	}
	return commentsResponse
}

func toMovieRatingResponse(rating *repository.MovieRating) *ratingResponse {
	return &ratingResponse{
		CreatedAt: rating.CreatedAt,
		UpdatedAt: rating.UpdatedAt,
		UserID:    rating.UserID,
		MediaID:   rating.MovieID,
		Rating:    rating.Rating,
	}
}

func toMovieRatingsResponse(ratings []*repository.MovieRating) []*ratingResponse {
	var ratingsResponse = make([]*ratingResponse, len(ratings))
	for i, rating := range ratings {
		ratingsResponse[i] = toMovieRatingResponse(rating)
	}
	return ratingsResponse
}

func toTVShowRatingResponse(rating *repository.TvShowRating) *ratingResponse {
	return &ratingResponse{
		CreatedAt: rating.CreatedAt,
		UpdatedAt: rating.UpdatedAt,
		UserID:    rating.UserID,
		MediaID:   rating.TvShowID,
		Rating:    rating.Rating,
	}
}

func toTVShowRatingsResponse(ratings []*repository.TvShowRating) []*ratingResponse {
	var ratingsResponse = make([]*ratingResponse, len(ratings))
	for i, rating := range ratings {
		ratingsResponse[i] = toTVShowRatingResponse(rating)
	}
	return ratingsResponse
}

func toCommentHistories(movieComments []*repository.MovieComment, tvShowComments []*repository.TvShowComment) []*commentHistoryReponse {
	var commentMap = make(map[string]*commentHistoryReponse)
	for _, comment := range movieComments {
		date := comment.CreatedAt.Format("2006-01-02")
		if _, ok := commentMap[date]; !ok {
			commentMap[date] = &commentHistoryReponse{
				Date:  date,
				Count: 1,
			}
		} else {
			commentMap[date].Count++
		}
	}
	for _, comment := range tvShowComments {
		date := comment.CreatedAt.Format("2006-01-02")
		if _, ok := commentMap[date]; !ok {
			commentMap[date] = &commentHistoryReponse{
				Date:  date,
				Count: 1,
			}
		} else {
			commentMap[date].Count++
		}
	}
	var commentHistories = make([]*commentHistoryReponse, len(commentMap))
	i := 0
	for _, commentHistory := range commentMap {
		commentHistories[i] = commentHistory
		i++
	}
	sort.Slice(commentHistories, func(i, j int) bool {
		return commentHistories[i].Date < commentHistories[j].Date
	})
	return commentHistories
}

func toEpisodeFileResponse(episode *repository.Episode) *episodeFileResponse {
	return &episodeFileResponse{
		ID:            episode.ID,
		Name:          episode.Name,
		ReleaseDate:   episode.ReleaseDate.Format("2006-01-02"),
		SeasonNumber:  episode.NbSeason,
		EpisodeNumber: episode.NbEpisode,
		TvShowId:      episode.TvShowID,
		TvShowName:    episode.TvShow.Name,
		File:          toMediaFileResponse(episode.MediaFile),
	}
}

func toEpisodeFilesResponse(episodes []*repository.Episode) []*episodeFileResponse {
	var episodesResponse = make([]*episodeFileResponse, len(episodes))
	for i, episode := range episodes {
		episodesResponse[i] = toEpisodeFileResponse(episode)
	}
	return episodesResponse
}

func toMovieFileResponse(movie *repository.Movie) *movieFileResponse {
	return &movieFileResponse{
		ID:          movie.ID,
		Name:        movie.Name,
		ReleaseDate: movie.ReleaseDate.Format("2006-01-02"),
		File:        toMediaFileResponse(movie.MediaFile),
	}
}

func toMovieFilesResponse(movies []*repository.Movie) []*movieFileResponse {
	var moviesResponse = make([]*movieFileResponse, len(movies))
	for i, movie := range movies {
		moviesResponse[i] = toMovieFileResponse(movie)
	}
	return moviesResponse
}
