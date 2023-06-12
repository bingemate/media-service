package features

import (
	"errors"
	repository2 "github.com/bingemate/media-go-pkg/repository"
	"github.com/bingemate/media-service/internal/repository"
	"gorm.io/gorm"
)

type MediaFile struct {
	moviePath       string
	tvPath          string
	mediaRepository *repository.MediaRepository
}

func NewMediaFile(moviePath string, tvPath string, mediaRepository *repository.MediaRepository) *MediaFile {
	return &MediaFile{
		moviePath:       moviePath,
		tvPath:          tvPath,
		mediaRepository: mediaRepository,
	}
}

//// GetMediaFileInfo returns a media file info given the mediaID (TMDB ID)
//func (m *MediaFile) GetMediaFileInfo(mediaID int) (*repository2.MediaFile, error) {
//	media, err := m.mediaRepository.GetMedia(mediaID)
//	if err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return nil, ErrMediaNotFound
//		}
//		return nil, err
//	}
//	if media.MediaType == repository2.MediaTypeTvShow {
//		return nil, ErrInvalidMediaType
//	}
//	if media.MediaType == repository2.MediaTypeMovie {
//		file, err := m.mediaRepository.GetMovieFileInfo(mediaID)
//		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
//			return nil, ErrMediaNotFound
//		}
//		return file, err
//	}
//	if media.MediaType == repository2.MediaTypeEpisode {
//		file, err := m.mediaRepository.GetEpisodeFileInfo(mediaID)
//		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
//			return nil, ErrMediaNotFound
//		}
//		return file, err
//	}
//	return nil, ErrInvalidMediaType
//}

// GetMovieFileInfo returns a movie file info given the movieID (TMDB ID)
func (m *MediaFile) GetMovieFileInfo(movieID int) (*repository2.MediaFile, error) {
	file, err := m.mediaRepository.GetMovieFileInfo(movieID)
	if (err != nil && errors.Is(err, gorm.ErrRecordNotFound)) || file == nil {
		return nil, ErrMediaNotFound
	}
	return file, err
}

// GetEpisodeFileInfo returns a episode file info given the episodeID (TMDB ID)
func (m *MediaFile) GetEpisodeFileInfo(episodeID int) (*repository2.MediaFile, error) {
	file, err := m.mediaRepository.GetEpisodeFileInfo(episodeID)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) || file == nil {
		return nil, ErrMediaNotFound
	}
	return file, err
}

// GetAvailableEpisode returns all available episodes id for a given tv show
func (m *MediaFile) GetAvailableEpisode(tvShowID int) (*[]int, error) {
	episodes, err := m.mediaRepository.AvailableEpisodes(tvShowID)
	if err != nil {
		return nil, err
	}
	return episodes, nil
}

// SearchEpisodeFiles returns all episodes that match the query
func (m *MediaFile) SearchEpisodeFiles(query string, page, limit int) ([]*repository2.Episode, int, error) {
	return m.mediaRepository.SearchEpisodeFiles(query, page, limit)
}

// SearchMovieFiles returns all movies that match the query
func (m *MediaFile) SearchMovieFiles(query string, page, limit int) ([]*repository2.Movie, int, error) {
	return m.mediaRepository.SearchMovieFiles(query, page, limit)
}

// DeleteMediaFile deletes a media file given the fileID
func (m *MediaFile) DeleteMediaFile(fileID string) error {
	return m.mediaRepository.DeleteMediaFile(fileID)
}
