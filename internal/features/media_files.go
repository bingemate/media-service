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

func (m *MediaFile) GetMediaFileInfo(mediaID string) (*repository2.MediaFile, error) {
	media, err := m.mediaRepository.GetMedia(mediaID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMediaNotFound
		}
		return nil, err
	}
	if media.MediaType == repository2.MediaTypeTvShow {
		return nil, ErrInvalidMediaType
	}
	if media.MediaType == repository2.MediaTypeMovie {
		return m.mediaRepository.GetMovieFileInfo(mediaID)
	}
	if media.MediaType == repository2.MediaTypeEpisode {
		return m.mediaRepository.GetEpisodeFileInfo(mediaID)
	}
	return nil, ErrInvalidMediaType
}
