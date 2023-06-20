package features

import (
	"errors"
	repository2 "github.com/bingemate/media-go-pkg/repository"
	"github.com/bingemate/media-go-pkg/tmdb"
	"github.com/bingemate/media-service/internal/repository"
	"gorm.io/gorm"
	"sync"
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

//// GetMediaByID returns a media given the mediaID (TMDB ID)
//func (m *MediaData) GetMediaByID(id int) (*repository2.Media, error) {
//	media, err := m.mediaRepository.GetMedia(id)
//	if err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return nil, ErrMediaNotFound
//		}
//		return nil, err
//	}
//	return media, nil
//}

func (m *MediaData) GetMovieByID(id int) (*repository2.Movie, error) {
	movie, err := m.mediaRepository.GetMovie(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMediaNotFound
		}
		return nil, err
	}
	return movie, nil
}

func (m *MediaData) GetEpisodeByID(id int) (*repository2.Episode, error) {
	episode, err := m.mediaRepository.GetEpisode(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMediaNotFound
		}
		return nil, err
	}
	return episode, nil
}

func (m *MediaData) GetEpisodesByIDs(ids []int) ([]*repository2.Episode, error) {
	episodes := make([]*repository2.Episode, len(ids))
	wg := sync.WaitGroup{}
	wg.Add(len(ids))
	for i, id := range ids {
		go func(i, id int) {
			defer wg.Done()
			episode, err := m.GetEpisodeByID(id)
			if err != nil {
				episodes[i] = nil
			}
			episodes[i] = episode
		}(i, id)
	}
	wg.Wait()
	return episodes, nil
}

func (m *MediaData) GetTvShowByID(id int) (*repository2.TvShow, error) {
	tvShow, err := m.mediaRepository.GetTvShow(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMediaNotFound
		}
		return nil, err
	}
	return tvShow, nil
}

// GetMovieInfo returns a movie given the mediaID (TMDB ID)
func (m *MediaData) GetMovieInfo(id int) (*tmdb.Movie, bool, error) {
	movie, err := m.mediaClient.GetMovie(id)
	if err != nil {
		return nil, false, err
	}
	voteAverage, voteCount, err := m.mediaRepository.GetMovieRating(id)
	if err == nil {
		movie.VoteAverage = voteAverage
		movie.VoteCount = voteCount
	}
	err = m.mediaRepository.SaveMovie(movie)
	if err != nil {
		return nil, false, err
	}
	return movie, m.mediaRepository.IsMovieFilePresent(id), nil
}

// GetMovieShortInfo returns a movie given the mediaID (TMDB ID)
func (m *MediaData) GetMovieShortInfo(id int) (*tmdb.Movie, bool, error) {
	movie, err := m.mediaClient.GetMovieShort(id)
	if err != nil {
		return nil, false, err
	}
	voteAverage, voteCount, err := m.mediaRepository.GetMovieRating(id)
	if err == nil {
		movie.VoteAverage = voteAverage
		movie.VoteCount = voteCount
	}
	err = m.mediaRepository.SaveMovie(movie)
	if err != nil {
		return nil, false, err
	}
	return movie, m.mediaRepository.IsMovieFilePresent(id), nil
}

// GetMoviesShortInfo returns a list of movies given the mediaID (TMDB ID)
func (m *MediaData) GetMoviesShortInfo(ids []int) ([]*tmdb.Movie, *[]bool, error) {
	movies := make([]*tmdb.Movie, len(ids))
	presences := make([]bool, len(ids))
	wg := sync.WaitGroup{}
	wg.Add(len(ids))
	for i, id := range ids {
		go func(i, id int) {
			defer wg.Done()
			movie, present, err := m.GetMovieShortInfo(id)
			if err != nil {
				movies[i] = nil
				presences[i] = false
			}
			movies[i] = movie
			presences[i] = present
		}(i, id)
	}
	wg.Wait()
	return movies, &presences, nil
}

// GetEpisodeInfo returns an episode info given the tvID (TMDB ID), season and episode number
func (m *MediaData) GetEpisodeInfo(tvID, season, episodeNumber int) (*tmdb.TVEpisode, bool, error) {
	episode, err := m.mediaClient.GetTVEpisode(tvID, season, episodeNumber)
	if err != nil {
		return nil, false, err
	}
	err = m.mediaRepository.SaveEpisode(episode)
	if err != nil {
		return nil, false, err
	}
	return episode, m.mediaRepository.IsEpisodeFilePresent(episode.ID), nil
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
	return m.GetEpisodeInfo(episode.TvShow.ID, episode.NbSeason, episode.NbEpisode)
}

// GetEpisodesInfoByIDs returns a list of episodes info given the episodeIDs (TMDB ID)
func (m *MediaData) GetEpisodesInfoByIDs(episodeIDs []int) ([]*tmdb.TVEpisode, *[]bool, error) {
	episodes := make([]*tmdb.TVEpisode, len(episodeIDs))
	presences := make([]bool, len(episodeIDs))
	wg := sync.WaitGroup{}
	wg.Add(len(episodeIDs))
	for i, id := range episodeIDs {
		go func(i, id int) {
			defer wg.Done()
			episode, present, err := m.GetEpisodeInfoByID(id)
			if err != nil {
				episodes[i] = nil
				presences[i] = false
			}
			episodes[i] = episode
			presences[i] = present
		}(i, id)
	}
	wg.Wait()
	return episodes, &presences, nil
}

// GetTvShowInfo returns a tv show given the mediaID (TMDB ID)
func (m *MediaData) GetTvShowInfo(mediaID int) (*tmdb.TVShow, bool, error) {
	tvShow, err := m.mediaClient.GetTVShow(mediaID)
	if err != nil {
		return nil, false, err
	}
	voteAverage, voteCount, err := m.mediaRepository.GetTvShowRating(mediaID)
	if err == nil {
		tvShow.VoteAverage = voteAverage
		tvShow.VoteCount = voteCount
	}
	err = m.mediaRepository.SaveTvShow(tvShow)
	return tvShow, m.mediaRepository.IsTvShowHasEpisodeFiles(mediaID), nil
}

// GetTvShowShortInfo returns a tv show given the mediaID (TMDB ID)
func (m *MediaData) GetTvShowShortInfo(mediaID int) (*tmdb.TVShow, bool, error) {
	tvShow, err := m.mediaClient.GetTVShowShort(mediaID)
	if err != nil {
		return nil, false, err
	}
	voteAverage, voteCount, err := m.mediaRepository.GetTvShowRating(mediaID)
	if err == nil {
		tvShow.VoteAverage = voteAverage
		tvShow.VoteCount = voteCount
	}
	err = m.mediaRepository.SaveTvShow(tvShow)
	return tvShow, m.mediaRepository.IsTvShowHasEpisodeFiles(mediaID), nil
}

// GetTvShowsShortInfo returns a list of tv shows given the mediaID (TMDB ID)
func (m *MediaData) GetTvShowsShortInfo(ids []int) ([]*tmdb.TVShow, *[]bool, error) {
	tvShows := make([]*tmdb.TVShow, len(ids))
	presences := make([]bool, len(ids))
	wg := sync.WaitGroup{}
	wg.Add(len(ids))
	for i, id := range ids {
		go func(i, id int) {
			defer wg.Done()
			tvShow, present, err := m.GetTvShowShortInfo(id)
			if err != nil {
				tvShows[i] = nil
				presences[i] = false
			}
			tvShows[i] = tvShow
			presences[i] = present
		}(i, id)
	}
	wg.Wait()
	return tvShows, &presences, nil
}

// GetSeasonEpisodes returns a list of episodes given the tvID (TMDB ID) and season number
func (m *MediaData) GetSeasonEpisodes(tvID, season int) ([]*tmdb.TVEpisode, *[]bool, error) {
	episodes, err := m.mediaClient.GetTVSeasonEpisodes(tvID, season)
	presence := make([]bool, len(episodes))
	if err != nil {
		return nil, nil, err
	}
	for i, episode := range episodes {
		presence[i] = m.mediaRepository.IsEpisodeFilePresent(episode.ID)
		err = m.mediaRepository.SaveEpisode(episode)
		if err != nil {
			return nil, nil, err
		}
	}
	return episodes, &presence, nil
}

// GetTvShowEpisodes returns a list of episodes given the tvID (TMDB ID)
func (m *MediaData) GetTvShowEpisodes(tvID int) ([]*tmdb.TVEpisode, *[]bool, error) {
	tvShow, _, err := m.GetTvShowInfo(tvID)
	if err != nil {
		return nil, nil, err
	}
	episodes := make([]*tmdb.TVEpisode, 0)
	presence := make([]bool, 0)
	for i := 1; i <= tvShow.SeasonsCount; i++ {
		seasonEpisodes, seasonPresence, err := m.GetSeasonEpisodes(tvID, i)
		if err != nil {
			return nil, nil, err
		}
		episodes = append(episodes, seasonEpisodes...)
		presence = append(presence, *seasonPresence...)
	}
	return episodes, &presence, nil
}
