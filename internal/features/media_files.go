package features

import (
	"errors"
	objectStorage "github.com/bingemate/media-go-pkg/object-storage"
	repository2 "github.com/bingemate/media-go-pkg/repository"
	"github.com/bingemate/media-service/internal/repository"
	"gorm.io/gorm"
	"path"
	"strconv"
	"syscall"
)

type MediaFile struct {
	moviePath       string
	tvPath          string
	mediaRepository *repository.MediaRepository
	objectStorage   objectStorage.ObjectStorage // Object storage object to upload the media files.
}

func NewMediaFile(moviePath string, tvPath string, mediaRepository *repository.MediaRepository, objectStorage objectStorage.ObjectStorage) *MediaFile {
	return &MediaFile{
		moviePath:       moviePath,
		tvPath:          tvPath,
		mediaRepository: mediaRepository,
		objectStorage:   objectStorage,
	}
}

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
	// Check if the file is present in movie or tv show folder
	// If it is, delete it
	episode, err := m.mediaRepository.GetEpisodeByFileID(fileID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if episode != nil {
		tvPath := path.Join("tv-shows", strconv.Itoa(episode.ID))
		go m.objectStorage.DeleteMediaFiles(tvPath)
	}

	movie, err := m.mediaRepository.GetMovieByFileID(fileID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if movie != nil {
		moviePath := path.Join("movies", strconv.Itoa(movie.ID))
		go m.objectStorage.DeleteMediaFiles(moviePath)
	}
	return m.mediaRepository.DeleteMediaFile(fileID)
}

// MediaFilesTotalSize returns the total size of all media files
func (m *MediaFile) MediaFilesTotalSize() (int64, error) {
	return m.mediaRepository.MediaFilesTotalSize()
}

// MediaFilesCount returns the total number of media files
func (m *MediaFile) MediaFilesCount() (int64, error) {
	return m.mediaRepository.MediaFilesCount()
}

// CountAvailableMovies returns the total number of available movies
func (m *MediaFile) CountAvailableMovies() (int64, error) {
	return m.mediaRepository.CountAvailableMovies()
}

// CountAvailableEpisodes returns the total number of available episodes
func (m *MediaFile) CountAvailableEpisodes() (int64, error) {
	return m.mediaRepository.CountAvailableEpisodes()
}

// CountAvailableTvShows returns the total number of available tv shows
func (m *MediaFile) CountAvailableTvShows() (int64, error) {
	return m.mediaRepository.CountAvailableTvShows()
}

// CountMoviesTotalDuration returns the total duration of all movies
func (m *MediaFile) CountMoviesTotalDuration() (int64, error) {
	return m.mediaRepository.CountMoviesTotalDuration()
}

// CountEpisodesTotalDuration returns the total duration of all episodes
func (m *MediaFile) CountEpisodesTotalDuration() (int64, error) {
	return m.mediaRepository.CountEpisodesTotalDuration()
}

// AvailableSpace returns the available space in the media folder
func (m *MediaFile) AvailableSpace() (uint64, error) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(m.moviePath, &fs)
	if err != nil {
		return 0, err
	}
	return fs.Bavail * uint64(fs.Bsize), nil
}
