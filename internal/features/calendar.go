package features

import (
	repository2 "github.com/bingemate/media-go-pkg/repository"
	"github.com/bingemate/media-go-pkg/tmdb"
	"github.com/bingemate/media-service/internal/repository"
	"log"
	"sync"
	"time"
)

type CalendarService struct {
	mediaClient     tmdb.MediaClient
	mediaRepository *repository.MediaRepository
}

func NewCalendarService(mediaClient tmdb.MediaClient, mediaRepository *repository.MediaRepository) *CalendarService {
	return &CalendarService{mediaClient, mediaRepository}
}

func (s *CalendarService) GetMoviesCalendar(userID string) ([]*tmdb.Movie, *[]bool, error) {
	month := int(time.Now().Month())
	followedMovies, err := s.mediaRepository.GetFollowedMovieReleases(userID, month)
	if err != nil {
		return nil, nil, err
	}
	var movies = make([]*tmdb.Movie, len(*followedMovies))
	var presence = make([]bool, len(*followedMovies))
	var locks = make([]sync.Mutex, len(*followedMovies))
	var wg sync.WaitGroup
	for i, followedMovie := range *followedMovies {
		wg.Add(1)
		go func(i int, followedMovie repository2.Movie) {
			defer wg.Done()
			movie, err := s.mediaClient.GetMovieShort(followedMovie.MediaID)
			if err != nil {
				log.Println("error getting movie", followedMovie.MediaID, err)
				return
			}
			locks[i].Lock()
			defer locks[i].Unlock()
			movies[i] = movie
			presence[i] = s.mediaRepository.IsMediaPresent(followedMovie.MediaID)
		}(i, followedMovie)
	}
	wg.Wait()
	return movies, &presence, nil
}

func (s *CalendarService) GetTvShowCalendar(userID string) ([]*tmdb.TVEpisode, *[]bool, error) {
	month := int(time.Now().Month())
	startOfMonth := time.Date(time.Now().Year(), time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endOfMonth := time.Date(time.Now().Year(), time.Month(month), 1, 0, 0, 0, 0, time.UTC).AddDate(0, 1, 0)
	followedTvShows, err := s.mediaRepository.GetFollowedTvShowReleases(userID, month)
	if err != nil {
		return nil, nil, err
	}
	showIds := make([]int, len(*followedTvShows))
	presence := make([]bool, len(*followedTvShows))
	for i, followedTvShow := range *followedTvShows {
		showIds[i] = followedTvShow.MediaID
		presence[i] = s.mediaRepository.IsMediaPresent(followedTvShow.MediaID)
	}

	tvShows, err := s.mediaClient.GetTVShowsReleases(showIds, startOfMonth, endOfMonth)
	if err != nil {
		return nil, nil, err
	}

	return tvShows, &presence, nil
}
