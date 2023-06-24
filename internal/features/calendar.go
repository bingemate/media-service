package features

import (
	"github.com/bingemate/media-go-pkg/tmdb"
	"github.com/bingemate/media-service/internal/repository"
	"log"
	"time"
)

type CalendarService struct {
	mediaClient     tmdb.MediaClient
	mediaRepository *repository.MediaRepository
}

func NewCalendarService(mediaClient tmdb.MediaClient, mediaRepository *repository.MediaRepository) *CalendarService {
	return &CalendarService{mediaClient, mediaRepository}
}

func (s *CalendarService) GetMoviesCalendar(userID string, month int, year int) ([]*tmdb.Movie, *[]bool, error) {
	startOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endOfMonth := startOfMonth.AddDate(0, 1, 0)
	followedReleases, err := s.mediaRepository.GetFollowedMoviesReleases(userID)
	if err != nil {
		return nil, nil, err
	}
	var presence []bool
	movies, err := s.mediaClient.GetMoviesReleases(*followedReleases, startOfMonth, endOfMonth)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}
	presence = make([]bool, len(movies))
	for i, movie := range movies {
		presence[i] = s.mediaRepository.IsMovieFilePresent(movie.ID)
	}
	return movies, &presence, nil
}

func (s *CalendarService) GetTvShowCalendar(userID string, month int, year int) ([]*tmdb.TVEpisode, []*tmdb.TVShow, *[]bool, error) {
	startOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endOfMonth := startOfMonth.AddDate(0, 1, 0)
	followedReleases, err := s.mediaRepository.GetFollowedTvShowsReleases(userID)
	if err != nil {
		return nil, nil, nil, err
	}
	episodes, tvShows, err := s.mediaClient.GetTVShowsReleases(*followedReleases, startOfMonth, endOfMonth)
	if err != nil {
		log.Println(err)
		return nil, nil, nil, err
	}
	presence := make([]bool, len(episodes))
	for i, episode := range episodes {
		presence[i] = s.mediaRepository.IsEpisodeFilePresent(episode.ID)
	}

	return episodes, tvShows, &presence, nil
}

func (s *CalendarService) GetMoviesCalendarInRange(userID string, start time.Time, end time.Time) ([]*tmdb.Movie, *[]bool, error) {
	followedReleases, err := s.mediaRepository.GetFollowedMoviesReleases(userID)
	if err != nil {
		return nil, nil, err
	}
	var presence []bool
	movies, err := s.mediaClient.GetMoviesReleases(*followedReleases, start, end)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}
	presence = make([]bool, len(movies))
	for i, movie := range movies {
		presence[i] = s.mediaRepository.IsMovieFilePresent(movie.ID)
	}
	return movies, &presence, nil
}

func (s *CalendarService) GetTvShowCalendarInRange(userID string, start time.Time, end time.Time) ([]*tmdb.TVEpisode, []*tmdb.TVShow, *[]bool, error) {
	followedReleases, err := s.mediaRepository.GetFollowedTvShowsReleases(userID)
	if err != nil {
		return nil, nil, nil, err
	}
	episodes, tvShows, err := s.mediaClient.GetTVShowsReleases(*followedReleases, start, end)
	if err != nil {
		log.Println(err)
		return nil, nil, nil, err
	}
	presence := make([]bool, len(episodes))
	for i, episode := range episodes {
		presence[i] = s.mediaRepository.IsEpisodeFilePresent(episode.ID)
	}

	return episodes, tvShows, &presence, nil
}
