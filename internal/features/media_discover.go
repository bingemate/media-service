package features

import (
	"github.com/bingemate/media-go-pkg/tmdb"
	"github.com/bingemate/media-service/internal/repository"
	"log"
	"math"
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

func (m *MediaDiscovery) SearchMovie(query string, page int, adult, available bool) (*tmdb.PaginatedMovieResults, *[]bool, error) {
	if available {
		return m.searchAvailableMovie(query, page)
	}
	return m.searchAllMovie(query, page, adult)
}

func (m *MediaDiscovery) searchAllMovie(query string, page int, adult bool) (*tmdb.PaginatedMovieResults, *[]bool, error) {
	movies, err := m.mediaClient.SearchMovies(query, page, adult)
	if err != nil {
		return nil, nil, err
	}
	presence := make([]bool, len(movies.Results))
	for i, movie := range movies.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetMovieRating(movie.ID)
		if err == nil {
			movie.VoteAverage = voteAverage
			movie.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsMovieFilePresent(movie.ID)
	}
	return movies, &presence, nil
}

func (m *MediaDiscovery) searchAvailableMovie(query string, page int) (*tmdb.PaginatedMovieResults, *[]bool, error) {
	movies, total, err := m.mediaRepository.SearchAvailableMovies(page, 20, query)
	if err != nil {
		return nil, nil, err
	}
	presence := make([]bool, len(movies))
	results := make([]*tmdb.Movie, len(movies))
	for i, movie := range movies {
		result, err := m.mediaClient.GetMovieShort(movie.ID)
		if err != nil {
			log.Println("error getting movie", movie.ID, err)
			return nil, nil, err
		}
		voteAverage, voteCount, err := m.mediaRepository.GetMovieRating(movie.ID)
		if err == nil {
			result.VoteAverage = voteAverage
			result.VoteCount = voteCount
		}
		results[i] = result
		presence[i] = true
	}
	return &tmdb.PaginatedMovieResults{
		Results:     results,
		TotalResult: total,
		TotalPage:   int(math.Round(float64(total) / 20)),
	}, &presence, nil
}

func (m *MediaDiscovery) SearchShow(query string, page int, adult, available bool) (*tmdb.PaginatedTVShowResults, *[]bool, error) {
	if available {
		return m.searchAvailableShow(query, page)
	}
	return m.searchAllShow(query, page, adult)
}

func (m *MediaDiscovery) searchAllShow(query string, page int, adult bool) (*tmdb.PaginatedTVShowResults, *[]bool, error) {
	shows, err := m.mediaClient.SearchTVShows(query, page, adult)
	if err != nil {
		return nil, nil, err
	}
	presence := make([]bool, len(shows.Results))
	for i, show := range shows.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetTvShowRating(show.ID)
		if err == nil {
			show.VoteAverage = voteAverage
			show.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsTvShowHasEpisodeFiles(show.ID)
	}
	return shows, &presence, nil
}

func (m *MediaDiscovery) SearchActor(query string, page int, adult bool) (*tmdb.PaginatedActorResults, error) {
	return m.mediaClient.SearchActors(query, page, adult)
}

func (m *MediaDiscovery) searchAvailableShow(query string, page int) (*tmdb.PaginatedTVShowResults, *[]bool, error) {
	shows, total, err := m.mediaRepository.SearchAvailableTvShows(page, 20, query)
	if err != nil {
		return nil, nil, err
	}
	presence := make([]bool, len(shows))
	results := make([]*tmdb.TVShow, len(shows))
	for i, show := range shows {

		result, err := m.mediaClient.GetTVShowShort(show.ID)
		if err != nil {
			log.Println("error getting show", show.ID, err)
			return nil, nil, err
		}
		voteAverage, voteCount, err := m.mediaRepository.GetTvShowRating(show.ID)
		if err == nil {
			result.VoteAverage = voteAverage
			result.VoteCount = voteCount
		}
		results[i] = result
		presence[i] = true
	}
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
	if err != nil {
		return nil, nil, err
	}
	presence := make([]bool, len(movies.Results))
	for i, movie := range movies.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetMovieRating(movie.ID)
		if err == nil {
			movie.VoteAverage = voteAverage
			movie.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsMovieFilePresent(movie.ID)
	}
	return movies, &presence, nil
}

func (m *MediaDiscovery) getAvailablePopularMovies(page int) (*tmdb.PaginatedMovieResults, *[]bool, error) {
	movies, total, err := m.mediaRepository.GetAvailableMoviesByRating(page, 20, 30)
	if err != nil {
		return nil, nil, err
	}
	presence := make([]bool, len(movies))
	results := make([]*tmdb.Movie, len(movies))
	for i, movie := range movies {
		result, err := m.mediaClient.GetMovieShort(movie.ID)
		if err != nil {
			log.Println("error getting movie", movie.ID, err)
			return nil, nil, err
		}
		voteAverage, voteCount, err := m.mediaRepository.GetMovieRating(movie.ID)
		if err == nil {
			result.VoteAverage = voteAverage
			result.VoteCount = voteCount
		}
		results[i] = result
		presence[i] = true
	}
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
	if err != nil {
		return nil, nil, err
	}
	presence := make([]bool, len(shows.Results))
	for i, show := range shows.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetTvShowRating(show.ID)
		if err == nil {
			show.VoteAverage = voteAverage
			show.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsTvShowHasEpisodeFiles(show.ID)
	}
	return shows, &presence, nil
}

func (m *MediaDiscovery) getAvailablePopularTVShows(page int) (*tmdb.PaginatedTVShowResults, *[]bool, error) {
	shows, total, err := m.mediaRepository.GetAvailableTvShowsByRating(page, 20, 30)
	if err != nil {
		return nil, nil, err
	}
	presence := make([]bool, len(shows))
	results := make([]*tmdb.TVShow, len(shows))
	for i, show := range shows {
		result, err := m.mediaClient.GetTVShowShort(show.ID)
		if err != nil {
			log.Println("error getting show", show.ID, err)
			return nil, nil, err
		}
		voteAverage, voteCount, err := m.mediaRepository.GetTvShowRating(show.ID)
		if err == nil {
			result.VoteAverage = voteAverage
			result.VoteCount = voteCount
		}
		results[i] = result
		presence[i] = true
	}
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
	if err != nil {
		return nil, nil, err
	}
	presence := make([]bool, len(movies))
	for i, movie := range movies {
		voteAverage, voteCount, err := m.mediaRepository.GetMovieRating(movie.ID)
		if err == nil {
			movie.VoteAverage = voteAverage
			movie.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsMovieFilePresent(movie.ID)
	}
	return movies, &presence, nil
}

func (m *MediaDiscovery) getAvailableRecentMovies() ([]*tmdb.Movie, *[]bool, error) {
	movies, _, err := m.mediaRepository.GetAvailableRecentMovies(1, 20)
	if err != nil {
		return nil, nil, err
	}
	presence := make([]bool, len(movies))
	results := make([]*tmdb.Movie, len(movies))
	for i, movie := range movies {
		result, err := m.mediaClient.GetMovieShort(movie.ID)
		if err != nil {
			log.Println("error getting movie", movie.ID, err)
			return nil, nil, err
		}
		voteAverage, voteCount, err := m.mediaRepository.GetMovieRating(movie.ID)
		if err == nil {
			result.VoteAverage = voteAverage
			result.VoteCount = voteCount
		}
		results[i] = result
		presence[i] = true
	}
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
	if err != nil {
		return nil, nil, err
	}
	presence := make([]bool, len(shows))
	for i, show := range shows {
		voteAverage, voteCount, err := m.mediaRepository.GetTvShowRating(show.ID)
		if err == nil {
			show.VoteAverage = voteAverage
			show.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsTvShowHasEpisodeFiles(show.ID)
	}
	return shows, &presence, nil
}

func (m *MediaDiscovery) getAvailableRecentShows() ([]*tmdb.TVShow, *[]bool, error) {
	shows, _, err := m.mediaRepository.GetAvailableRecentTvShows(1, 20)
	if err != nil {
		return nil, nil, err
	}
	presence := make([]bool, len(shows))
	results := make([]*tmdb.TVShow, len(shows))
	for i, show := range shows {
		result, err := m.mediaClient.GetTVShowShort(show.ID)
		if err != nil {
			log.Println("error getting show", show.ID, err)
			return nil, nil, err
		}
		voteAverage, voteCount, err := m.mediaRepository.GetTvShowRating(show.ID)
		if err == nil {
			result.VoteAverage = voteAverage
			result.VoteCount = voteCount
		}
		results[i] = result
		presence[i] = true
	}
	return results, &presence, nil
}

func (m *MediaDiscovery) GetMoviesByGenre(genreID int, page int) (*tmdb.PaginatedMovieResults, *[]bool, error) {
	movies, err := m.mediaClient.GetMoviesByGenre(genreID, page)
	if err != nil {
		return nil, nil, err
	}
	presence := make([]bool, len(movies.Results))
	for i, movie := range movies.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetMovieRating(movie.ID)
		if err == nil {
			movie.VoteAverage = voteAverage
			movie.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsMovieFilePresent(movie.ID)
	}
	return movies, &presence, nil
}

func (m *MediaDiscovery) GetShowsByGenre(genreID int, page int) (*tmdb.PaginatedTVShowResults, *[]bool, error) {
	shows, err := m.mediaClient.GetTVShowsByGenre(genreID, page)
	if err != nil {
		return nil, nil, err
	}
	presence := make([]bool, len(shows.Results))
	for i, show := range shows.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetTvShowRating(show.ID)
		if err == nil {
			show.VoteAverage = voteAverage
			show.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsTvShowHasEpisodeFiles(show.ID)
	}
	return shows, &presence, nil
}

func (m *MediaDiscovery) GetMoviesByActor(actorID int, page int) (*tmdb.PaginatedMovieResults, *[]bool, error) {
	movies, err := m.mediaClient.GetMoviesByActor(actorID, page)
	if err != nil {
		return nil, nil, err
	}
	presence := make([]bool, len(movies.Results))
	for i, movie := range movies.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetMovieRating(movie.ID)
		if err == nil {
			movie.VoteAverage = voteAverage
			movie.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsMovieFilePresent(movie.ID)
	}
	return movies, &presence, nil
}

func (m *MediaDiscovery) GetShowsByActor(actorID int, page int) (*tmdb.PaginatedTVShowResults, *[]bool, error) {
	shows, err := m.mediaClient.GetTVShowsByActor(actorID, page)
	if err != nil {
		return nil, nil, err
	}
	presence := make([]bool, len(shows.Results))
	for i, show := range shows.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetTvShowRating(show.ID)
		if err == nil {
			show.VoteAverage = voteAverage
			show.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsTvShowHasEpisodeFiles(show.ID)
	}
	return shows, &presence, nil
}

func (m *MediaDiscovery) GetMoviesByDirector(directorID int, page int) (*tmdb.PaginatedMovieResults, *[]bool, error) {
	movies, err := m.mediaClient.GetMoviesByDirector(directorID, page)
	if err != nil {
		return nil, nil, err
	}
	presence := make([]bool, len(movies.Results))
	for i, movie := range movies.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetMovieRating(movie.ID)
		if err == nil {
			movie.VoteAverage = voteAverage
			movie.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsMovieFilePresent(movie.ID)
	}
	return movies, &presence, nil
}

func (m *MediaDiscovery) GetMoviesByStudio(studioID int, page int) (*tmdb.PaginatedMovieResults, *[]bool, error) {
	movies, err := m.mediaClient.GetMoviesByStudio(studioID, page)
	if err != nil {
		return nil, nil, err
	}
	presence := make([]bool, len(movies.Results))
	for i, movie := range movies.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetMovieRating(movie.ID)
		if err == nil {
			movie.VoteAverage = voteAverage
			movie.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsMovieFilePresent(movie.ID)
	}
	return movies, &presence, nil
}

func (m *MediaDiscovery) GetShowsByNetwork(networkID int, page int) (*tmdb.PaginatedTVShowResults, *[]bool, error) {
	shows, err := m.mediaClient.GetTVShowsByNetwork(networkID, page)
	if err != nil {
		return nil, nil, err
	}
	presence := make([]bool, len(shows.Results))
	for i, show := range shows.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetTvShowRating(show.ID)
		if err == nil {
			show.VoteAverage = voteAverage
			show.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsTvShowHasEpisodeFiles(show.ID)
	}
	return shows, &presence, nil
}

func (m *MediaDiscovery) GetMovieRecommendations(movieID int) ([]*tmdb.Movie, *[]bool, error) {
	movies, err := m.mediaClient.GetMovieRecommendations(movieID)
	if err != nil {
		return nil, nil, err
	}
	presence := make([]bool, len(movies))
	for i, movie := range movies {
		voteAverage, voteCount, err := m.mediaRepository.GetMovieRating(movie.ID)
		if err == nil {
			movie.VoteAverage = voteAverage
			movie.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsMovieFilePresent(movie.ID)
	}
	return movies, &presence, nil
}

func (m *MediaDiscovery) GetShowRecommendations(showID int) ([]*tmdb.TVShow, *[]bool, error) {
	shows, err := m.mediaClient.GetTVShowRecommendations(showID)
	if err != nil {
		return nil, nil, err
	}
	presence := make([]bool, len(shows))
	for i, show := range shows {
		voteAverage, voteCount, err := m.mediaRepository.GetTvShowRating(show.ID)
		if err == nil {
			show.VoteAverage = voteAverage
			show.VoteCount = voteCount
		}
		presence[i] = m.mediaRepository.IsTvShowHasEpisodeFiles(show.ID)
	}
	return shows, &presence, nil
}

//func (m *MediaDiscovery) GetMediasByComments(present bool) (*[]int, error) {
//	return m.mediaRepository.GetMediasByComments(present)
//}

func (m *MediaDiscovery) GetMoviesByComments(present bool) (*[]int, error) {
	return m.mediaRepository.GetMoviesByComments(present)
}

func (m *MediaDiscovery) GetShowsByComments(present bool) (*[]int, error) {
	return m.mediaRepository.GetTvShowsByComments(present)
}
