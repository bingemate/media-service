package features

import (
	"errors"
	repository2 "github.com/bingemate/media-go-pkg/repository"
	"github.com/bingemate/media-go-pkg/tmdb"
	"github.com/bingemate/media-service/internal/repository"
	"gorm.io/gorm"
)

var MediaNotFoundError = errors.New("media not found")
var InvalidMediaTypeError = errors.New("invalid media type")

type Rating struct {
	Rating float32 `json:"rating"`
	Count  int     `json:"count"`
}

type MediaData struct {
	moviePath       string
	tvPath          string
	mediaClient     tmdb.MediaClient
	mediaRepository *repository.MediaRepository
}

func NewMediaData(moviePath, tvPath string, mediaClient tmdb.MediaClient, movieRepository *repository.MediaRepository) *MediaData {
	return &MediaData{
		moviePath:       moviePath,
		tvPath:          tvPath,
		mediaClient:     mediaClient,
		mediaRepository: movieRepository,
	}
}

func (m *MediaData) GetMediaByTmdbID(tmdbID int) (*repository2.Media, error) {
	media, err := m.mediaRepository.GetMediaByTmdbID(tmdbID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, MediaNotFoundError
		}
		return nil, err
	}
	return media, nil
}

func (m *MediaData) GetMediaByID(id string) (*repository2.Media, error) {
	media, err := m.mediaRepository.GetMedia(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, MediaNotFoundError
		}
		return nil, err
	}
	return media, nil
}

// GetMovieInfoByTMDB returns the movie info from TMDB with given TMDB ID
func (m *MediaData) GetMovieInfoByTMDB(tmdbID int) (*tmdb.Movie, error) {
	movie, err := m.mediaClient.GetMovie(tmdbID)
	if err != nil {
		return nil, err
	}
	voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(tmdbID)
	if err == nil {
		movie.VoteAverage = voteAverage
		movie.VoteCount = voteCount
	}
	return movie, nil
}

func (m *MediaData) GetMovieInfo(mediaID string) (*tmdb.Movie, error) {
	media, err := m.mediaRepository.GetMedia(mediaID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, MediaNotFoundError
		}
		return nil, err
	}
	return m.GetMovieInfoByTMDB(media.TmdbID)
}

func (m *MediaData) GetMediaFileInfo(mediaID string) (*repository2.MediaFile, error) {
	media, err := m.mediaRepository.GetMedia(mediaID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, MediaNotFoundError
		}
		return nil, err
	}
	if media.MediaType == repository2.MediaTypeTvShow {
		return nil, InvalidMediaTypeError
	}
	if media.MediaType == repository2.MediaTypeMovie {
		return m.mediaRepository.GetMovieFileInfo(mediaID)
	}
	if media.MediaType == repository2.MediaTypeEpisode {
		return m.mediaRepository.GetEpisodeFileInfo(mediaID)
	}
	return nil, InvalidMediaTypeError
}
