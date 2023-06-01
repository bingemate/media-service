package features

import (
	repository2 "github.com/bingemate/media-go-pkg/repository"
	"github.com/bingemate/media-go-pkg/tmdb"
	"github.com/bingemate/media-service/internal/repository"
	"log"
	"math"
	"sync"
)

type MediaDiscovery struct {
	mediaClient     tmdb.MediaClient
	mediaRepository *repository.MediaRepository
}

func NewMediaDiscovery(mediaClient tmdb.MediaClient, mediaRepository *repository.MediaRepository) *MediaDiscovery {
	return &MediaDiscovery{
		mediaClient:     mediaClient,
		mediaRepository: mediaRepository,
	}
}

func (m *MediaDiscovery) SearchMovie(query string, page int, available bool) (*tmdb.PaginatedMovieResults, *[]bool, error) {
	if available {
		return m.searchAvailableMovie(query, page)
	}
	return m.searchAllMovie(query, page)
}

func (m *MediaDiscovery) searchAllMovie(query string, page int) (*tmdb.PaginatedMovieResults, *[]bool, error) {
	movies, err := m.mediaClient.SearchMovies(query, page)
	presence := make([]bool, len(movies.Results))
	if err != nil {
		return nil, nil, err
	}
	for i, movie := range movies.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(movie.ID)
		if err == nil {
			movie.VoteAverage = voteAverage
			movie.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsMediaPresent(movie.ID)
	}
	return movies, &presence, nil
}

func (m *MediaDiscovery) searchAvailableMovie(query string, page int) (*tmdb.PaginatedMovieResults, *[]bool, error) {
	movies, total, err := m.mediaRepository.SearchMovies(page, 20, query)
	if err != nil {
		return nil, nil, err
	}
	presence := make([]bool, len(movies))
	results := make([]*tmdb.Movie, len(movies))
	locker := make([]sync.Mutex, len(movies))
	wg := sync.WaitGroup{}
	for i, movie := range movies {
		wg.Add(1)
		go func(i int, movie repository2.Movie) {
			defer wg.Done()
			result, err := m.mediaClient.GetMovieShort(movie.MediaID)
			if err != nil {
				log.Println("error getting movie", movie.MediaID, err)
				return
			}
			voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(movie.MediaID)
			if err == nil {
				result.VoteAverage = voteAverage
				result.VoteCount = voteCount
			}
			locker[i].Lock()
			defer locker[i].Unlock()
			results[i] = result
			presence[i] = true
		}(i, movie)
	}
	wg.Wait()
	return &tmdb.PaginatedMovieResults{
		Results:     results,
		TotalResult: total,
		TotalPage:   int(math.Round(float64(total) / 20)),
	}, &presence, nil
}

func (m *MediaDiscovery) SearchShow(query string, page int, available bool) (*tmdb.PaginatedTVShowResults, *[]bool, error) {
	if available {
		return m.searchAvailableShow(query, page)
	}
	return m.searchAllShow(query, page)
}

func (m *MediaDiscovery) searchAllShow(query string, page int) (*tmdb.PaginatedTVShowResults, *[]bool, error) {
	shows, err := m.mediaClient.SearchTVShows(query, page)
	presence := make([]bool, len(shows.Results))
	if err != nil {
		return nil, nil, err
	}
	for i, show := range shows.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(show.ID)
		if err == nil {
			show.VoteAverage = voteAverage
			show.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsMediaPresent(show.ID)
	}
	return shows, &presence, nil
}

func (m *MediaDiscovery) searchAvailableShow(query string, page int) (*tmdb.PaginatedTVShowResults, *[]bool, error) {
	shows, total, err := m.mediaRepository.SearchTvShows(page, 20, query)
	if err != nil {
		return nil, nil, err
	}
	presence := make([]bool, len(shows))
	results := make([]*tmdb.TVShow, len(shows))
	locker := make([]sync.Mutex, len(shows))
	wg := sync.WaitGroup{}
	for i, show := range shows {
		wg.Add(1)
		go func(i int, show repository2.TvShow) {
			defer wg.Done()
			result, err := m.mediaClient.GetTVShowShort(show.MediaID)
			if err != nil {
				log.Println("error getting show", show.MediaID, err)
				return
			}
			voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(show.MediaID)
			if err == nil {
				result.VoteAverage = voteAverage
				result.VoteCount = voteCount
			}
			locker[i].Lock()
			defer locker[i].Unlock()
			results[i] = result
			presence[i] = true
		}(i, show)
	}
	wg.Wait()
	return &tmdb.PaginatedTVShowResults{
		Results:     results,
		TotalResult: total,
		TotalPage:   int(math.Round(float64(total) / 20)),
	}, &presence, nil
}

func (m *MediaDiscovery) GetPopularMovies(page int, available bool) (*tmdb.PaginatedMovieResults, *[]bool, error) {
	if available {
		return m.getAvailablePopularMovies(page)
	}
	return m.getAllPopularMovies(page)
}

func (m *MediaDiscovery) getAllPopularMovies(page int) (*tmdb.PaginatedMovieResults, *[]bool, error) {
	movies, err := m.mediaClient.GetPopularMovies(page)
	presence := make([]bool, len(movies.Results))
	if err != nil {
		return nil, nil, err
	}
	for i, movie := range movies.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(movie.ID)
		if err == nil {
			movie.VoteAverage = voteAverage
			movie.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsMediaPresent(movie.ID)
	}
	return movies, &presence, nil
}

func (m *MediaDiscovery) getAvailablePopularMovies(page int) (*tmdb.PaginatedMovieResults, *[]bool, error) {
	movies, total, err := m.mediaRepository.GetMoviesByRating(page, 20, 30)
	if err != nil {
		return nil, nil, err
	}
	presence := make([]bool, len(movies))
	results := make([]*tmdb.Movie, len(movies))
	locker := make([]sync.Mutex, len(movies))
	wg := sync.WaitGroup{}
	for i, movie := range movies {
		wg.Add(1)
		go func(i int, movie repository2.Movie) {
			defer wg.Done()
			result, err := m.mediaClient.GetMovieShort(movie.MediaID)
			if err != nil {
				log.Println("error getting movie", movie.MediaID, err)
				return
			}
			voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(movie.MediaID)
			if err == nil {
				result.VoteAverage = voteAverage
				result.VoteCount = voteCount
			}
			locker[i].Lock()
			defer locker[i].Unlock()
			results[i] = result
			presence[i] = true
		}(i, movie)
	}
	wg.Wait()
	return &tmdb.PaginatedMovieResults{
		Results:     results,
		TotalResult: total,
		TotalPage:   int(math.Round(float64(total) / 20)),
	}, &presence, nil
}

func (m *MediaDiscovery) GetPopularShows(page int, available bool) (*tmdb.PaginatedTVShowResults, *[]bool, error) {
	if available {
		return m.getAvailablePopularTVShows(page)
	}
	return m.getAllPopularTVShows(page)
}

func (m *MediaDiscovery) getAllPopularTVShows(page int) (*tmdb.PaginatedTVShowResults, *[]bool, error) {
	shows, err := m.mediaClient.GetPopularTVShows(page)
	presence := make([]bool, len(shows.Results))
	if err != nil {
		return nil, nil, err
	}
	for i, show := range shows.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(show.ID)
		if err == nil {
			show.VoteAverage = voteAverage
			show.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsMediaPresent(show.ID)
	}
	return shows, &presence, nil
}

func (m *MediaDiscovery) getAvailablePopularTVShows(page int) (*tmdb.PaginatedTVShowResults, *[]bool, error) {
	shows, total, err := m.mediaRepository.GetTvShowsByRating(page, 20, 30)
	if err != nil {
		return nil, nil, err
	}
	presence := make([]bool, len(shows))
	results := make([]*tmdb.TVShow, len(shows))
	locker := make([]sync.Mutex, len(shows))
	wg := sync.WaitGroup{}
	for i, show := range shows {
		wg.Add(1)
		go func(i int, show repository2.TvShow) {
			defer wg.Done()
			result, err := m.mediaClient.GetTVShowShort(show.MediaID)
			if err != nil {
				log.Println("error getting show", show.MediaID, err)
				return
			}
			voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(show.MediaID)
			if err == nil {
				result.VoteAverage = voteAverage
				result.VoteCount = voteCount
			}
			locker[i].Lock()
			defer locker[i].Unlock()
			results[i] = result
			presence[i] = true
		}(i, show)
	}
	wg.Wait()
	return &tmdb.PaginatedTVShowResults{
		Results:     results,
		TotalResult: total,
		TotalPage:   int(math.Round(float64(total) / 20)),
	}, &presence, nil
}

func (m *MediaDiscovery) GetRecentMovies(available bool) ([]*tmdb.Movie, *[]bool, error) {
	if available {
		return m.getAvailableRecentMovies()
	}
	return m.getAllRecentMovies()
}

func (m *MediaDiscovery) getAllRecentMovies() ([]*tmdb.Movie, *[]bool, error) {
	movies, err := m.mediaClient.GetRecentMovies()
	presence := make([]bool, len(movies))
	if err != nil {
		return nil, nil, err
	}
	for i, movie := range movies {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(movie.ID)
		if err == nil {
			movie.VoteAverage = voteAverage
			movie.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsMediaPresent(movie.ID)
	}
	return movies, &presence, nil
}

func (m *MediaDiscovery) getAvailableRecentMovies() ([]*tmdb.Movie, *[]bool, error) {
	movies, _, err := m.mediaRepository.GetRecentMovies(1, 20)
	if err != nil {
		return nil, nil, err
	}
	presence := make([]bool, len(movies))
	results := make([]*tmdb.Movie, len(movies))
	locker := make([]sync.Mutex, len(movies))
	wg := sync.WaitGroup{}
	for i, movie := range movies {
		wg.Add(1)
		go func(i int, movie repository2.Movie) {
			defer wg.Done()
			result, err := m.mediaClient.GetMovieShort(movie.MediaID)
			if err != nil {
				log.Println("error getting movie", movie.MediaID, err)
				return
			}
			voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(movie.MediaID)
			if err == nil {
				result.VoteAverage = voteAverage
				result.VoteCount = voteCount
			}
			locker[i].Lock()
			defer locker[i].Unlock()
			results[i] = result
			presence[i] = true
		}(i, movie)
	}
	wg.Wait()
	return results, &presence, nil
}

func (m *MediaDiscovery) GetRecentShows(available bool) ([]*tmdb.TVShow, *[]bool, error) {
	if available {
		return m.getAvailableRecentShows()
	}
	return m.getAllRecentShows()
}

func (m *MediaDiscovery) getAllRecentShows() ([]*tmdb.TVShow, *[]bool, error) {
	shows, err := m.mediaClient.GetRecentTVShows()
	presence := make([]bool, len(shows))
	if err != nil {
		return nil, nil, err
	}
	for i, show := range shows {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(show.ID)
		if err == nil {
			show.VoteAverage = voteAverage
			show.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsMediaPresent(show.ID)
	}
	return shows, &presence, nil
}

func (m *MediaDiscovery) getAvailableRecentShows() ([]*tmdb.TVShow, *[]bool, error) {
	shows, _, err := m.mediaRepository.GetRecentTvShows(1, 20)
	if err != nil {
		return nil, nil, err
	}
	presence := make([]bool, len(shows))
	results := make([]*tmdb.TVShow, len(shows))
	locker := make([]sync.Mutex, len(shows))
	wg := sync.WaitGroup{}
	for i, show := range shows {
		wg.Add(1)
		go func(i int, show repository2.TvShow) {
			defer wg.Done()
			result, err := m.mediaClient.GetTVShowShort(show.MediaID)
			if err != nil {
				log.Println("error getting show", show.MediaID, err)
				return
			}
			voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(show.MediaID)
			if err == nil {
				result.VoteAverage = voteAverage
				result.VoteCount = voteCount
			}
			locker[i].Lock()
			defer locker[i].Unlock()
			results[i] = result
			presence[i] = true
		}(i, show)
	}
	wg.Wait()
	return results, &presence, nil
}

func (m *MediaDiscovery) GetMoviesByGenre(genreID int, page int) (*tmdb.PaginatedMovieResults, *[]bool, error) {
	movies, err := m.mediaClient.GetMoviesByGenre(genreID, page)
	presence := make([]bool, len(movies.Results))
	if err != nil {
		return nil, nil, err
	}
	for i, movie := range movies.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(movie.ID)
		if err == nil {
			movie.VoteAverage = voteAverage
			movie.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsMediaPresent(movie.ID)
	}
	return movies, &presence, nil
}

func (m *MediaDiscovery) GetShowsByGenre(genreID int, page int) (*tmdb.PaginatedTVShowResults, *[]bool, error) {
	shows, err := m.mediaClient.GetTVShowsByGenre(genreID, page)
	presence := make([]bool, len(shows.Results))
	if err != nil {
		return nil, nil, err
	}
	for i, show := range shows.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(show.ID)
		if err == nil {
			show.VoteAverage = voteAverage
			show.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsMediaPresent(show.ID)
	}
	return shows, &presence, nil
}

func (m *MediaDiscovery) GetMoviesByActor(actorID int, page int) (*tmdb.PaginatedMovieResults, *[]bool, error) {
	movies, err := m.mediaClient.GetMoviesByActor(actorID, page)
	presence := make([]bool, len(movies.Results))
	if err != nil {
		return nil, nil, err
	}
	for i, movie := range movies.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(movie.ID)
		if err == nil {
			movie.VoteAverage = voteAverage
			movie.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsMediaPresent(movie.ID)
	}
	return movies, &presence, nil
}

func (m *MediaDiscovery) GetShowsByActor(actorID int, page int) (*tmdb.PaginatedTVShowResults, *[]bool, error) {
	shows, err := m.mediaClient.GetTVShowsByActor(actorID, page)
	presence := make([]bool, len(shows.Results))
	if err != nil {
		return nil, nil, err
	}
	for i, show := range shows.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(show.ID)
		if err == nil {
			show.VoteAverage = voteAverage
			show.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsMediaPresent(show.ID)
	}
	return shows, &presence, nil
}

func (m *MediaDiscovery) GetMoviesByDirector(directorID int, page int) (*tmdb.PaginatedMovieResults, *[]bool, error) {
	movies, err := m.mediaClient.GetMoviesByDirector(directorID, page)
	presence := make([]bool, len(movies.Results))
	if err != nil {
		return nil, nil, err
	}
	for i, movie := range movies.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(movie.ID)
		if err == nil {
			movie.VoteAverage = voteAverage
			movie.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsMediaPresent(movie.ID)
	}
	return movies, &presence, nil
}

func (m *MediaDiscovery) GetMoviesByStudio(studioID int, page int) (*tmdb.PaginatedMovieResults, *[]bool, error) {
	movies, err := m.mediaClient.GetMoviesByStudio(studioID, page)
	presence := make([]bool, len(movies.Results))
	if err != nil {
		return nil, nil, err
	}
	for i, movie := range movies.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(movie.ID)
		if err == nil {
			movie.VoteAverage = voteAverage
			movie.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsMediaPresent(movie.ID)
	}
	return movies, &presence, nil
}

func (m *MediaDiscovery) GetShowsByNetwork(networkID int, page int) (*tmdb.PaginatedTVShowResults, *[]bool, error) {
	shows, err := m.mediaClient.GetTVShowsByNetwork(networkID, page)
	presence := make([]bool, len(shows.Results))
	if err != nil {
		return nil, nil, err
	}
	for i, show := range shows.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(show.ID)
		if err == nil {
			show.VoteAverage = voteAverage
			show.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsMediaPresent(show.ID)
	}
	return shows, &presence, nil
}

func (m *MediaDiscovery) GetMovieRecommendations(movieID int) ([]*tmdb.Movie, *[]bool, error) {
	movies, err := m.mediaClient.GetMovieRecommendations(movieID)
	presence := make([]bool, len(movies))
	if err != nil {
		return nil, nil, err
	}
	for i, movie := range movies {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(movie.ID)
		if err == nil {
			movie.VoteAverage = voteAverage
			movie.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsMediaPresent(movie.ID)
	}
	return movies, &presence, nil
}

func (m *MediaDiscovery) GetShowRecommendations(showID int) ([]*tmdb.TVShow, *[]bool, error) {
	shows, err := m.mediaClient.GetTVShowRecommendations(showID)
	presence := make([]bool, len(shows))
	if err != nil {
		return nil, nil, err
	}
	for i, show := range shows {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(show.ID)
		if err == nil {
			show.VoteAverage = voteAverage
			show.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsMediaPresent(show.ID)
	}
	return shows, &presence, nil
}

func (m *MediaDiscovery) GetMediasByComments(present bool) (*[]int, error) {
	return m.mediaRepository.GetMediasByComments(present)
}
