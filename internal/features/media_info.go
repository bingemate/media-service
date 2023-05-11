package features

import (
	"errors"
	repository2 "github.com/bingemate/media-go-pkg/repository"
	"github.com/bingemate/media-go-pkg/tmdb"
	"github.com/bingemate/media-service/internal/repository"
	"gorm.io/gorm"
)

type MediaData struct {
	mediaClient     tmdb.MediaClient
	mediaRepository *repository.MediaRepository
}

func NewMediaData(mediaClient tmdb.MediaClient, mediaRepository *repository.MediaRepository) *MediaData {
	return &MediaData{
		mediaClient:     mediaClient,
		mediaRepository: mediaRepository,
	}
}

func (m *MediaData) GetMediaByTmdbID(tmdbID int) (*repository2.Media, error) {
	media, err := m.mediaRepository.GetMediaByTmdbID(tmdbID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMediaNotFound
		}
		return nil, err
	}
	return media, nil
}

func (m *MediaData) GetMediaByID(id string) (*repository2.Media, error) {
	media, err := m.mediaRepository.GetMedia(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMediaNotFound
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
			return nil, ErrMediaNotFound
		}
		return nil, err
	}
	if media.MediaType != repository2.MediaTypeMovie {
		return nil, ErrInvalidMediaType
	}
	return m.GetMovieInfoByTMDB(media.TmdbID)
}

func (m *MediaData) GetEpisodeInfoByTMDB(tvTmdbID, season, episodeNumber int) (*tmdb.TVEpisode, error) {
	episode, err := m.mediaClient.GetTVEpisode(tvTmdbID, season, episodeNumber)
	if err != nil {
		return nil, err
	}
	return episode, nil
}

func (m *MediaData) GetEpisodeInfo(mediaID string) (*tmdb.TVEpisode, error) {
	media, err := m.mediaRepository.GetMedia(mediaID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMediaNotFound
		}
		return nil, err
	}
	if media.MediaType != repository2.MediaTypeEpisode {
		return nil, ErrInvalidMediaType
	}
	episode, err := m.mediaRepository.GetEpisode(mediaID)
	if err != nil {
		return nil, err
	}
	return m.GetEpisodeInfoByTMDB(episode.TvShow.Media.TmdbID, episode.NbSeason, episode.NbEpisode)
}

func (m *MediaData) GetTvShowInfoByTMDB(tmdbID int) (*tmdb.TVShow, error) {
	tvShow, err := m.mediaClient.GetTVShow(tmdbID)
	if err != nil {
		return nil, err
	}
	voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(tmdbID)
	if err == nil {
		tvShow.VoteAverage = voteAverage
		tvShow.VoteCount = voteCount
	}
	return tvShow, nil
}

func (m *MediaData) GetTvShowInfo(mediaID string) (*tmdb.TVShow, error) {
	media, err := m.mediaRepository.GetMedia(mediaID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMediaNotFound
		}
		return nil, err
	}
	if media.MediaType != repository2.MediaTypeTvShow {
		return nil, ErrInvalidMediaType
	}
	return m.GetTvShowInfoByTMDB(media.TmdbID)
}

func (m *MediaData) GetSeasonEpisodesByTMDB(tvTmdbID, season int) ([]*tmdb.TVEpisode, error) {
	episodes, err := m.mediaClient.GetTVSeasonEpisodes(tvTmdbID, season)
	if err != nil {
		return nil, err
	}
	return episodes, nil
}

func (m *MediaData) GetSeasonEpisodes(tvMediaID string, season int) ([]*tmdb.TVEpisode, error) {
	media, err := m.mediaRepository.GetMedia(tvMediaID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMediaNotFound
		}
		return nil, err
	}
	if media.MediaType != repository2.MediaTypeTvShow {
		return nil, ErrInvalidMediaType
	}

	return m.GetSeasonEpisodesByTMDB(media.TmdbID, season)
}
