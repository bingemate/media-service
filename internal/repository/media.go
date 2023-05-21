package repository

import (
	"errors"
	"github.com/bingemate/media-go-pkg/repository"
	"gorm.io/gorm"
	"log"
)

type MediaRepository struct {
	db *gorm.DB
}

func NewMediaRepository(db *gorm.DB) *MediaRepository {
	if db == nil {
		log.Fatal("db is nil")
	}
	return &MediaRepository{db}
}

// GetMediaRating returns the average rating and the number of ratings for a media given the mediaID (TMDB ID)
func (r *MediaRepository) GetMediaRating(mediaID int) (float32, int, error) {
	var (
		media repository.Media
		sum   float32
		count int64
	)
	err := r.db.Where("id = ?", mediaID).First(&media).Error
	if err != nil {
		return 0, 0, errors.New("media not found")
	}

	err = r.db.Model(&repository.Rating{}).Where("media_id = ?", mediaID).Count(&count).Error
	if err != nil {
		return 0, 0, err
	}
	if count == 0 {
		return 0, 0, errors.New("no rating found")
	}
	err = r.db.Table("ratings").Select("SUM(rating) as total").Where("media_id = ?", mediaID).Scan(&sum).Error
	if err != nil {
		return 0, 0, err
	}
	return sum / float32(count), int(count), nil

}

// GetMedia returns a media given the mediaID (TMDB ID)
func (r *MediaRepository) GetMedia(mediaID int) (*repository.Media, error) {
	var media repository.Media
	err := r.db.Where("id = ?", mediaID).First(&media).Error
	if err != nil {
		return nil, err
	}
	return &media, nil
}

// GetEpisode returns an episode given the mediaID (TMDB ID)
func (r *MediaRepository) GetEpisode(mediaID int) (*repository.Episode, error) {
	var episode repository.Episode
	err := r.db.
		Joins("Media").
		Joins("TvShow").
		Preload("TvShow.Media").
		Where(`"Media".id = ?`, mediaID).First(&episode).Error
	if err != nil {
		return nil, err
	}
	return &episode, nil
}

// GetTvShow returns a tv show given the mediaID (TMDB ID)
func (r *MediaRepository) GetTvShow(mediaID int) (*repository.TvShow, error) {
	var tvShow repository.TvShow
	err := r.db.
		Joins("Media").
		Where("media_id = ?", mediaID).First(&tvShow).Error
	if err != nil {
		return nil, err
	}
	return &tvShow, nil
}

// GetEpisodeFileInfo returns the file info for an episode given the mediaID (TMDB ID)
func (r *MediaRepository) GetEpisodeFileInfo(mediaID int) (*repository.MediaFile, error) {
	var mediaFile repository.Episode
	err := r.db.
		Joins("Media").
		Joins("MediaFile").
		Preload("MediaFile.Audio").
		Preload("MediaFile.Subtitles").
		Where("media_id = ?", mediaID).
		First(&mediaFile).Error
	if err != nil {
		return nil, err
	}
	return &mediaFile.MediaFile, nil
}

// GetMovieFileInfo returns the file info for a movie given the mediaID (TMDB ID)
func (r *MediaRepository) GetMovieFileInfo(mediaID int) (*repository.MediaFile, error) {
	var mediaFile repository.Movie
	err := r.db.
		Joins("Media").
		Joins("MediaFile").
		Preload("MediaFile.Audio").
		Preload("MediaFile.Subtitles").
		Where("media_id = ?", mediaID).
		First(&mediaFile).Error
	if err != nil {
		return nil, err
	}
	return &mediaFile.MediaFile, nil
}

func (r *MediaRepository) IsMediaPresent(mediaID int) bool {
	var count int64
	r.db.Model(&repository.Media{}).Where("id = ?", mediaID).Count(&count)
	return count > 0
}
