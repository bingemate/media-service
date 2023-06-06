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

// GetMediaByID returns a media given the mediaID (TMDB ID)
func (m *MediaData) GetMediaByID(id int) (*repository2.Media, error) {
	media, err := m.mediaRepository.GetMedia(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMediaNotFound
		}
		return nil, err
	}
	return media, nil
}

// GetMovieInfo returns a movie given the mediaID (TMDB ID)
func (m *MediaData) GetMovieInfo(id int) (*tmdb.Movie, bool, error) {
	movie, err := m.mediaClient.GetMovie(id)
	if err != nil {
		return nil, false, err
	}
	voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(id)
	if err == nil {
		movie.VoteAverage = voteAverage
		movie.VoteCount = voteCount
	}
	return movie, m.mediaRepository.IsMediaPresent(id), nil
}

// GetMovieShortInfo returns a movie given the mediaID (TMDB ID)
func (m *MediaData) GetMovieShortInfo(id int) (*tmdb.Movie, bool, error) {
	movie, err := m.mediaClient.GetMovieShort(id)
	if err != nil {
		return nil, false, err
	}
	voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(id)
	if err == nil {
		movie.VoteAverage = voteAverage
		movie.VoteCount = voteCount
	}
	return movie, m.mediaRepository.IsMediaPresent(id), nil
}

// GetEpisodeInfo returns an episode info given the tvID (TMDB ID), season and episode number
func (m *MediaData) GetEpisodeInfo(tvID, season, episodeNumber int) (*tmdb.TVEpisode, bool, error) {
	episode, err := m.mediaClient.GetTVEpisode(tvID, season, episodeNumber)
	if err != nil {
		return nil, false, err
	}
	return episode, m.mediaRepository.IsMediaPresent(episode.ID), nil
}

// GetEpisodeInfoByID returns an episode info given the episodeID (TMDB ID)
func (m *MediaData) GetEpisodeInfoByID(episodeID int) (*tmdb.TVEpisode, bool, error) {
	episode, err := m.mediaRepository.GetEpisode(episodeID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, ErrMediaNotFound
		}
		return nil, false, err
	}
	return m.GetEpisodeInfo(episode.TvShow.Media.ID, episode.NbSeason, episode.NbEpisode)
}

// GetTvShowInfo returns a tv show given the mediaID (TMDB ID)
func (m *MediaData) GetTvShowInfo(mediaID int) (*tmdb.TVShow, bool, error) {
	tvShow, err := m.mediaClient.GetTVShow(mediaID)
	if err != nil {
		return nil, false, err
	}
	voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(mediaID)
	if err == nil {
		tvShow.VoteAverage = voteAverage
		tvShow.VoteCount = voteCount
	}
	return tvShow, m.mediaRepository.IsMediaPresent(mediaID), nil
}

// GetTvShowShortInfo returns a tv show given the mediaID (TMDB ID)
func (m *MediaData) GetTvShowShortInfo(mediaID int) (*tmdb.TVShow, bool, error) {
	tvShow, err := m.mediaClient.GetTVShowShort(mediaID)
	if err != nil {
		return nil, false, err
	}
	voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(mediaID)
	if err == nil {
		tvShow.VoteAverage = voteAverage
		tvShow.VoteCount = voteCount
	}
	return tvShow, m.mediaRepository.IsMediaPresent(mediaID), nil
}

// GetSeasonEpisodes returns a list of episodes given the tvID (TMDB ID) and season number
func (m *MediaData) GetSeasonEpisodes(tvID, season int) ([]*tmdb.TVEpisode, *[]bool, error) {
	episodes, err := m.mediaClient.GetTVSeasonEpisodes(tvID, season)
	presence := make([]bool, len(episodes))
	if err != nil {
		return nil, nil, err
	}
	for i, episode := range episodes {
		presence[i] = m.mediaRepository.IsMediaPresent(episode.ID)
	}
	return episodes, &presence, nil
}
