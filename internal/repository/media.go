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

func (r *MediaRepository) GetMediaRating(tmdbID int) (float32, int, error) {
	var media repository.Media
	var rating []repository.Rating
	err := r.db.Where("tmdb_id = ?", tmdbID).First(&media).Error
	if err != nil {
		return 0, 0, errors.New("media not found")
	}

	err = r.db.Where("media_id = ?", media.ID).Find(&rating).Error
	if err != nil {
		return 0, 0, err
	}
	var count = len(rating)
	if count == 0 {
		return 0, 0, errors.New("no rating found")
	}
	var sum float32
	for _, r := range rating {
		sum += float32(r.Rating)
	}
	return sum / float32(count), count, nil
}

func (r *MediaRepository) GetMedia(id string) (*repository.Media, error) {
	var media repository.Media
	err := r.db.Where("id = ?", id).First(&media).Error
	if err != nil {
		return nil, err
	}
	return &media, nil
}

func (r *MediaRepository) GetMediaByTmdbID(tmdbID int) (*repository.Media, error) {
	var media repository.Media
	err := r.db.Where("tmdb_id = ?", tmdbID).First(&media).Error
	if err != nil {
		return nil, err
	}
	return &media, nil
}

func (r *MediaRepository) GetEpisode(mediaID string) (*repository.Episode, error) {
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

func (r *MediaRepository) GetTvShow(mediaId string) (*repository.TvShow, error) {
	var tvShow repository.TvShow
	err := r.db.
		Joins("Media").
		Where("media_id = ?", mediaId).First(&tvShow).Error
	if err != nil {
		return nil, err
	}
	return &tvShow, nil
}

func (r *MediaRepository) GetEpisodeFileInfo(mediaID string) (*repository.MediaFile, error) {
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

func (r *MediaRepository) GetMovieFileInfo(mediaID string) (*repository.MediaFile, error) {
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
