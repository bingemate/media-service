package features

import (
	"errors"
	"github.com/bingemate/media-go-pkg/tmdb"
	"github.com/bingemate/media-service/internal/repository"
)

var MediaNotFoundError = errors.New("media not found")

type Rating struct {
	Rating float32 `json:"rating"`
	Count  int     `json:"count"`
}

type MediaData struct {
	moviePath       string
	tvPath          string
	mediaClient     tmdb.MediaClient
	movieRepository *repository.MediaRepository
}

func NewMediaData(moviePath, tvPath string, mediaClient tmdb.MediaClient, movieRepository *repository.MediaRepository) *MediaData {
	return &MediaData{
		moviePath:       moviePath,
		tvPath:          tvPath,
		mediaClient:     mediaClient,
		movieRepository: movieRepository,
	}
}

// GetMovieInfo returns the movie info from TMDB with given TMDB ID
func (m *MediaData) GetMovieInfo(tmdbID int) (*tmdb.Movie, error) {
	movie, err := m.mediaClient.GetMovie(tmdbID)
	if err != nil {
		return nil, err
	}
	voteAverage, voteCount, err := m.movieRepository.GetMediaRating(tmdbID)
	if err == nil {
		movie.VoteAverage = voteAverage
		movie.VoteCount = voteCount
	}
	return movie, nil
}
