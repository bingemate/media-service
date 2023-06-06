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

func (s *CalendarService) GetMoviesCalendar(userID string, month int) ([]*tmdb.Movie, *[]bool, error) {
	startOfMonth := time.Date(time.Now().Year(), time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endOfMonth := startOfMonth.AddDate(0, 1, 0)
	followedReleases, err := s.mediaRepository.GetFollowedReleases(userID, month)
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
		presence[i] = s.mediaRepository.IsMediaPresent(movie.ID)
	}
	return movies, &presence, nil
}

func (s *CalendarService) GetTvShowCalendar(userID string, month int) ([]*tmdb.TVEpisode, []*tmdb.TVShow, *[]bool, error) {
	startOfMonth := time.Date(time.Now().Year(), time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endOfMonth := startOfMonth.AddDate(0, 1, 0)
	followedReleases, err := s.mediaRepository.GetFollowedReleases(userID, month)
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
		presence[i] = s.mediaRepository.IsMediaPresent(episode.ID)
	}

	return episodes, tvShows, &presence, nil
}
