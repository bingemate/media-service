package repository

import (
	"errors"
	"github.com/bingemate/media-go-pkg/repository"
	"gorm.io/gorm"
	"log"
	"time"
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

// IsMediaPresent returns true if the media is present in the database
func (r *MediaRepository) IsMediaPresent(mediaID int) bool {
	var count int64
	r.db.Model(&repository.Media{}).Where("id = ?", mediaID).Count(&count)
	return count > 0
}

// GetMoviesByRating returns a list of movies ordered by rating
// Return also the total number of results
func (r *MediaRepository) GetMoviesByRating(page, limit, days int) ([]repository.Movie, int, error) {
	var movies []repository.Movie
	offset := (page - 1) * limit

	var count int64
	result := r.db.Table("movies").
		Select("movies.*, AVG(ratings.rating) as average_rating").
		Joins("LEFT JOIN ratings ON ratings.media_id = movies.media_id AND ratings.created_at > ?", time.Now().AddDate(0, 0, -days)).
		Group("movies.id,movies.media_id").
		Count(&count).
		Order("average_rating DESC").
		Offset(offset).
		Limit(limit).
		Find(&movies)

	if result.Error != nil {
		return nil, 0, result.Error
	}
	return movies, int(count), nil
}

// GetTvShowsByRating returns a list of tv shows ordered by rating
// Return also the total number of results
func (r *MediaRepository) GetTvShowsByRating(page, limit, days int) ([]repository.TvShow, int, error) {
	var tvShows []repository.TvShow
	offset := (page - 1) * limit
	var count int64
	result := r.db.Table("tv_shows").
		Select("tv_shows.*, AVG(ratings.rating) as average_rating").
		Joins("LEFT JOIN ratings ON ratings.media_id = tv_shows.media_id AND ratings.created_at > ?", time.Now().AddDate(0, 0, -days)).
		Group("tv_shows.id,tv_shows.media_id").
		Count(&count).
		Order("average_rating DESC").
		Offset(offset).
		Limit(limit).
		Find(&tvShows)

	if result.Error != nil {
		return nil, 0, result.Error
	}
	return tvShows, int(count), nil
}

// SearchMovies returns a list of movies matching the search query
// Return also the total number of results
// Results are ordered by pertinence and / or rating
func (r *MediaRepository) SearchMovies(page, limit int, query string) ([]repository.Movie, int, error) {
	var movies []repository.Movie
	offset := (page - 1) * limit

	var count int64
	result := r.db.Table("movies").
		Select("movies.*, AVG(ratings.rating) as average_rating").
		Joins("LEFT JOIN ratings ON ratings.media_id = movies.media_id").
		Where("movies.name ILIKE ?", "%"+query+"%").
		Group("movies.id,movies.media_id").
		Count(&count).
		Order("average_rating DESC, movies.name ASC").
		Offset(offset).
		Limit(limit).
		Find(&movies)

	if result.Error != nil {
		return nil, 0, result.Error
	}
	return movies, int(count), nil
}

// SearchTvShows returns a list of tv shows matching the search query
// Return also the total number of results
// Results are ordered by pertinence and / or rating
func (r *MediaRepository) SearchTvShows(page, limit int, query string) ([]repository.TvShow, int, error) {
	var tvShows []repository.TvShow
	offset := (page - 1) * limit

	var count int64
	result := r.db.Table("tv_shows").
		Select("tv_shows.*, AVG(ratings.rating) as average_rating").
		Joins("LEFT JOIN ratings ON ratings.media_id = tv_shows.media_id").
		Where("tv_shows.name ILIKE ?", "%"+query+"%").
		Group("tv_shows.id,tv_shows.media_id").
		Count(&count).
		Order("average_rating DESC, tv_shows.name ASC").
		Offset(offset).
		Limit(limit).
		Find(&tvShows)

	if result.Error != nil {
		return nil, 0, result.Error
	}
	return tvShows, int(count), nil
}

// GetRecentMovies returns a list of recently added movies
func (r *MediaRepository) GetRecentMovies(page, limit int) ([]repository.Movie, int, error) {
	var movies []repository.Movie
	offset := (page - 1) * limit
	var count int64
	result := r.db.Table("movies").
		Select("*").
		Count(&count).
		Order("movies.created_at DESC, movies.updated_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&movies)

	if result.Error != nil {
		return nil, 0, result.Error
	}
	return movies, int(count), nil
}

// GetRecentTvShows returns a list of recently added tv shows
func (r *MediaRepository) GetRecentTvShows(page, limit int) ([]repository.TvShow, int, error) {
	var tvShows []repository.TvShow
	offset := (page - 1) * limit
	var count int64
	result := r.db.Table("tv_shows").
		Select("*").
		Count(&count).
		Order("tv_shows.created_at DESC, tv_shows.updated_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&tvShows)

	if result.Error != nil {
		return nil, 0, result.Error
	}
	return tvShows, int(count), nil
}
