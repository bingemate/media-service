package pkg

import (
	"github.com/ryanbradynd05/go-tmdb"
)

const imageBaseURL = "https://image.tmdb.org/t/p/original"

type Genre struct {
	ID   int
	Name string
}

type Person struct {
	ID         int
	Character  string
	Name       string
	ProfileURL string
}

type Studio struct {
	ID      int
	Name    string
	LogoURL string
}

type Movie struct {
	ID          int
	Actors      []Person
	BackdropURL string
	Crew        []Person
	Genres      []Genre
	Overview    string
	PosterURL   string
	ReleaseDate string
	Studios     []Studio
	Title       string
	VoteAverage float32 // From our DB ?
	VoteCount   int     // From our DB ?
}

type TVEpisode struct {
}

type TVShow struct {
}

type MediaClient interface {
	GetMovie(id int) (*Movie, error)
	GetTVShow(id int) (*TVShow, error)
	GetTVEpisode(id int) (*TVEpisode, error)
	GetTVSeasonEpisodes(id int, season int) (*[]TVEpisode, error)
}

type mediaClient struct {
	tmdbClient *tmdb.TMDb
	options    map[string]string
}

func NewMediaClient(apiKey string) MediaClient {
	config := tmdb.Config{
		APIKey:   apiKey,
		Proxies:  nil,
		UseProxy: false,
	}
	return &mediaClient{
		tmdbClient: tmdb.Init(config),
		options: map[string]string{
			"language": "fr",
		},
	}
}

func (m *mediaClient) GetMovie(id int) (*Movie, error) {
	movie, err := m.tmdbClient.GetMovieInfo(id, m.options)
	if err != nil {
		return nil, err
	}
	credits, err := m.tmdbClient.GetMovieCredits(id, m.options)
	if err != nil {
		return nil, err
	}
	return &Movie{
		ID:          movie.ID,
		Actors:      *extractActors(credits),
		BackdropURL: imageBaseURL + movie.BackdropPath,
		Crew:        *extractCrew(credits),
		Genres:      *extractGenres(&movie.Genres),
		Overview:    movie.Overview,
		PosterURL:   imageBaseURL + movie.PosterPath,
		ReleaseDate: movie.ReleaseDate,
		Studios:     *extractStudios(&movie.ProductionCompanies),
		Title:       movie.Title,
		VoteAverage: movie.VoteAverage,    // TODO: From our DB ?
		VoteCount:   int(movie.VoteCount), // TODO: From our DB ?
	}, err
}

func (m *mediaClient) GetTVShow(id int) (*TVShow, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mediaClient) GetTVEpisode(id int) (*TVEpisode, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mediaClient) GetTVSeasonEpisodes(id int, season int) (*[]TVEpisode, error) {
	//TODO implement me
	panic("implement me")
}

func extractActors(credits *tmdb.MovieCredits) *[]Person {
	var actors = make([]Person, len(credits.Cast))
	for i, cast := range credits.Cast {
		actors[i] = Person{
			ID:         cast.ID,
			Character:  cast.Character,
			Name:       cast.Name,
			ProfileURL: imageBaseURL + cast.ProfilePath,
		}
	}
	return &actors
}

func extractCrew(credits *tmdb.MovieCredits) *[]Person {
	var crew = make([]Person, len(credits.Crew))
	for i, cast := range credits.Crew {
		crew[i] = Person{
			ID:         cast.ID,
			Character:  cast.Job,
			Name:       cast.Name,
			ProfileURL: imageBaseURL + cast.ProfilePath,
		}
	}
	return &crew
}

func extractGenres(genres *[]struct {
	ID   int
	Name string
}) *[]Genre {
	var extractedGenres = make([]Genre, len(*genres))
	for i, genre := range *genres {
		extractedGenres[i] = Genre{
			ID:   genre.ID,
			Name: genre.Name,
		}
	}
	return &extractedGenres
}

func extractStudios(studios *[]struct {
	ID        int
	Name      string
	LogoPath  string `json:"logo_path"`
	Iso3166_1 string `json:"origin_country"`
}) *[]Studio {
	var extractedStudios = make([]Studio, len(*studios))
	for i, studio := range *studios {
		extractedStudios[i] = Studio{
			ID:      studio.ID,
			Name:    studio.Name,
			LogoURL: imageBaseURL + studio.LogoPath,
		}
	}
	return &extractedStudios
}
