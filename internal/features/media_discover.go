package features

import (
	"github.com/bingemate/media-go-pkg/tmdb"
	"github.com/bingemate/media-service/internal/repository"
	"time"
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

func (m *MediaDiscovery) SearchMovie(query string, page int) (*tmdb.PaginatedMovieResults, error) {
	movies, err := m.mediaClient.SearchMovies(query, page)
	if err != nil {
		return nil, err
	}
	for _, movie := range movies.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(movie.ID)
		if err == nil {
			movie.VoteAverage = voteAverage
			movie.VoteCount = voteCount
		}
	}
	return movies, nil
}

func (m *MediaDiscovery) SearchShow(query string, page int) (*tmdb.PaginatedTVShowResults, error) {
	shows, err := m.mediaClient.SearchTVShows(query, page)
	if err != nil {
		return nil, err
	}
	for _, show := range shows.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(show.ID)
		if err == nil {
			show.VoteAverage = voteAverage
			show.VoteCount = voteCount
		}
	}
	return shows, nil
}

func (m *MediaDiscovery) GetPopularMovies(page int) (*tmdb.PaginatedMovieResults, error) {
	movies, err := m.mediaClient.GetPopularMovies(page)
	if err != nil {
		return nil, err
	}
	for _, movie := range movies.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(movie.ID)
		if err == nil {
			movie.VoteAverage = voteAverage
			movie.VoteCount = voteCount
		}
	}
	return movies, nil
}

func (m *MediaDiscovery) GetPopularShows(page int) (*tmdb.PaginatedTVShowResults, error) {
	shows, err := m.mediaClient.GetPopularTVShows(page)
	if err != nil {
		return nil, err
	}
	for _, show := range shows.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(show.ID)
		if err == nil {
			show.VoteAverage = voteAverage
			show.VoteCount = voteCount
		}
	}
	return shows, nil
}

func (m *MediaDiscovery) GetRecentMovies() ([]*tmdb.Movie, error) {
	movies, err := m.mediaClient.GetRecentMovies()
	if err != nil {
		return nil, err
	}
	for _, movie := range movies {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(movie.ID)
		if err == nil {
			movie.VoteAverage = voteAverage
			movie.VoteCount = voteCount
		}
	}
	return movies, nil
}

func (m *MediaDiscovery) GetRecentShows() ([]*tmdb.TVShow, error) {
	shows, err := m.mediaClient.GetRecentTVShows()
	if err != nil {
		return nil, err
	}
	for _, show := range shows {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(show.ID)
		if err == nil {
			show.VoteAverage = voteAverage
			show.VoteCount = voteCount
		}
	}
	return shows, nil
}

func (m *MediaDiscovery) GetMoviesByGenre(genreID int, page int) (*tmdb.PaginatedMovieResults, error) {
	movies, err := m.mediaClient.GetMoviesByGenre(genreID, page)
	if err != nil {
		return nil, err
	}
	for _, movie := range movies.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(movie.ID)
		if err == nil {
			movie.VoteAverage = voteAverage
			movie.VoteCount = voteCount
		}
	}
	return movies, nil
}

func (m *MediaDiscovery) GetShowsByGenre(genreID int, page int) (*tmdb.PaginatedTVShowResults, error) {
	shows, err := m.mediaClient.GetTVShowsByGenre(genreID, page)
	if err != nil {
		return nil, err
	}
	for _, show := range shows.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(show.ID)
		if err == nil {
			show.VoteAverage = voteAverage
			show.VoteCount = voteCount
		}
	}
	return shows, nil
}

func (m *MediaDiscovery) GetMoviesByActor(actorID int, page int) (*tmdb.PaginatedMovieResults, error) {
	movies, err := m.mediaClient.GetMoviesByActor(actorID, page)
	if err != nil {
		return nil, err
	}
	for _, movie := range movies.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(movie.ID)
		if err == nil {
			movie.VoteAverage = voteAverage
			movie.VoteCount = voteCount
		}
	}
	return movies, nil
}

func (m *MediaDiscovery) GetShowsByActor(actorID int, page int) (*tmdb.PaginatedTVShowResults, error) {
	shows, err := m.mediaClient.GetTVShowsByActor(actorID, page)
	if err != nil {
		return nil, err
	}
	for _, show := range shows.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(show.ID)
		if err == nil {
			show.VoteAverage = voteAverage
			show.VoteCount = voteCount
		}
	}
	return shows, nil
}

func (m *MediaDiscovery) GetMoviesByDirector(directorID int, page int) (*tmdb.PaginatedMovieResults, error) {
	movies, err := m.mediaClient.GetMoviesByDirector(directorID, page)
	if err != nil {
		return nil, err
	}
	for _, movie := range movies.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(movie.ID)
		if err == nil {
			movie.VoteAverage = voteAverage
			movie.VoteCount = voteCount
		}
	}
	return movies, nil
}

func (m *MediaDiscovery) GetShowsByDirector(directorID int, page int) (*tmdb.PaginatedTVShowResults, error) {
	shows, err := m.mediaClient.GetTVShowsByDirector(directorID, page)
	if err != nil {
		return nil, err
	}
	for _, show := range shows.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(show.ID)
		if err == nil {
			show.VoteAverage = voteAverage
			show.VoteCount = voteCount
		}
	}
	return shows, nil
}

func (m *MediaDiscovery) GetMoviesByStudio(studioID int, page int) (*tmdb.PaginatedMovieResults, error) {
	movies, err := m.mediaClient.GetMoviesByStudio(studioID, page)
	if err != nil {
		return nil, err
	}
	for _, movie := range movies.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(movie.ID)
		if err == nil {
			movie.VoteAverage = voteAverage
			movie.VoteCount = voteCount
		}
	}
	return movies, nil
}

func (m *MediaDiscovery) GetShowsByNetwork(networkID int, page int) (*tmdb.PaginatedTVShowResults, error) {
	shows, err := m.mediaClient.GetTVShowsByNetwork(networkID, page)
	if err != nil {
		return nil, err
	}
	for _, show := range shows.Results {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(show.ID)
		if err == nil {
			show.VoteAverage = voteAverage
			show.VoteCount = voteCount
		}
	}
	return shows, nil
}

func (m *MediaDiscovery) GetMovieRecommendations(movieID int) ([]*tmdb.Movie, error) {
	movies, err := m.mediaClient.GetMovieRecommendations(movieID)
	if err != nil {
		return nil, err
	}
	for _, movie := range movies {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(movie.ID)
		if err == nil {
			movie.VoteAverage = voteAverage
			movie.VoteCount = voteCount
		}
	}
	return movies, nil
}

func (m *MediaDiscovery) GetShowRecommendations(showID int) ([]*tmdb.TVShow, error) {
	shows, err := m.mediaClient.GetTVShowRecommendations(showID)
	if err != nil {
		return nil, err
	}
	for _, show := range shows {
		voteAverage, voteCount, err := m.mediaRepository.GetMediaRating(show.ID)
		if err == nil {
			show.VoteAverage = voteAverage
			show.VoteCount = voteCount
		}
	}
	return shows, nil
}

func (m *MediaDiscovery) GetTVShowsReleases(tvIds []int, startDate string, endDate string) ([]*tmdb.TVEpisodeRelease, error) {
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
}