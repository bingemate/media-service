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
	movies, presence, err := m.searchAllMovie(query, page)
	if err != nil {
		return nil, nil, err
	}
	return movies, presence, nil
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

func (m *MediaDiscovery) GetPopularMovies(page int, available bool) (*tmdb.PaginatedMovieResults, *[]bool, error) {
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

func (m *MediaDiscovery) GetPopularShows(page int, available bool) (*tmdb.PaginatedTVShowResults, *[]bool, error) {
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

func (m *MediaDiscovery) GetRecentMovies(available bool) ([]*tmdb.Movie, *[]bool, error) {
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

func (m *MediaDiscovery) GetRecentShows(available bool) ([]*tmdb.TVShow, *[]bool, error) {
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

// Add in another service
/*func (m *MediaDiscovery) GetTVShowsReleases(tvIds []int, startDate string, endDate string) ([]*tmdb.TVEpisodeRelease, error) {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, err
	}
	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, err
	}
	return m.mediaClient.GetTVShowsReleases(tvIds, start, end)
}

func (m *MediaDiscovery) GetMovieReleases(movieIds []int, startDate string, endDate string) ([]*tmdb.MovieRelease, error) {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, err
	}
	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, err
	}
	return m.mediaClient.GetMoviesReleases(movieIds, start, end)
}*/
