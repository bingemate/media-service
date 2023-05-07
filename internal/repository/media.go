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
		return 0, 0, err
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
