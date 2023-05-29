package features

import "github.com/bingemate/media-go-pkg/tmdb"

type MediaAssetsData struct {
	mediaClient tmdb.MediaClient
}

func NewMediaAssetsData(mediaClient tmdb.MediaClient) *MediaAssetsData {
	return &MediaAssetsData{
		mediaClient: mediaClient,
	}
}

func (m *MediaAssetsData) GetMovieGenre(id int) (*tmdb.Genre, error) {
	genres, err := m.mediaClient.GetMovieGenre(id)
	if err != nil {
		return nil, err
	}
	return genres, nil
}

func (m *MediaAssetsData) GetMovieGenres() ([]*tmdb.Genre, error) {
	genres, err := m.mediaClient.GetMovieGenres()
	if err != nil {
		return nil, err
	}
	return genres, nil
}

func (m *MediaAssetsData) GetTVGenre(id int) (*tmdb.Genre, error) {
	genres, err := m.mediaClient.GetTVGenre(id)
	if err != nil {
		return nil, err
	}
	return genres, nil
}

func (m *MediaAssetsData) GetTVGenres() ([]*tmdb.Genre, error) {
	genres, err := m.mediaClient.GetTVShowGenres()
	if err != nil {
		return nil, err
	}
	return genres, nil
}

func (m *MediaAssetsData) GetStudio(id int) (*tmdb.Studio, error) {
	studio, err := m.mediaClient.GetStudio(id)
	if err != nil {
		return nil, err
	}
	return studio, nil
}

func (m *MediaAssetsData) GetNetwork(id int) (*tmdb.Studio, error) {
	network, err := m.mediaClient.GetNetwork(id)
	if err != nil {
		return nil, err
	}
	return network, nil
}

func (m *MediaAssetsData) GetActor(id int) (*tmdb.Actor, error) {
	actor, err := m.mediaClient.GetActor(id)
	if err != nil {
		return nil, err
	}
	return actor, nil
}
