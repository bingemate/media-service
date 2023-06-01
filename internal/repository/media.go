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
		sum   float32
		count int64
	)

	err := r.db.Model(&repository.Rating{}).Where("media_id = ?", mediaID).Count(&count).Error
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

// GetMediasByComments returns a list of medias ordered by number of comments
func (r *MediaRepository) GetMediasByComments() (*[]int, error) {
	var mediaIds []int

	err := r.db.Model(&repository.Comment{}).
		Select("media_id").
		Group("media_id").
		Order("COUNT(comments.id) DESC").
		Limit(20).
		Find(&mediaIds).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &mediaIds, nil
}

func (r *MediaRepository) GetFollowedReleases(userID string, month int) (*[]int, error) {
	var followedReleases []int
	result := r.db.Table("watch_list_item").
		Select("media_id").
		Where("user_id = ? AND status != ?", userID, repository.WatchListStatusAbandoned).
		Find(&followedReleases)

	if result.Error != nil {
		return nil, result.Error
	}

	return &followedReleases, nil
}

func (r *MediaRepository) GetMediaComments(mediaID, size, page int) ([]*repository.Comment, int, error) {
	var comments []*repository.Comment
	var count int64
	offset := (page - 1) * size
	result := r.db.Model(&repository.Comment{}).
		Where("media_id = ?", mediaID).
		Count(&count).
		Order("created_at DESC").
		Offset(offset).
		Limit(size).
		Find(&comments)

	if result.Error != nil {
		return nil, 0, result.Error
	}
	return comments, int(count), nil
}

func (r *MediaRepository) GetUserComments(userID string, size, page int) ([]*repository.Comment, int, error) {
	var comments []*repository.Comment
	var count int64
	offset := (page - 1) * size
	result := r.db.Model(&repository.Comment{}).
		Where("user_id = ?", userID).
		Count(&count).
		Order("created_at DESC").
		Offset(offset).
		Limit(size).
		Find(&comments)

	if result.Error != nil {
		return nil, 0, result.Error
	}
	return comments, int(count), nil
}

func (r *MediaRepository) AddComment(userID string, mediaID int, content string) (*repository.Comment, error) {
	comment := repository.Comment{
		UserID:  userID,
		MediaID: mediaID,
		Content: content,
	}
	result := r.db.Create(&comment)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comment, nil
}

func (r *MediaRepository) GetComment(commentID string) (*repository.Comment, error) {
	var comment repository.Comment
	result := r.db.Model(&repository.Comment{}).Where("id = ?", commentID).First(&comment)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comment, nil
}

func (r *MediaRepository) DeleteComment(commentID string) error {
	return r.db.Where("id = ?", commentID).Delete(&repository.Comment{}).Error
}

func (r *MediaRepository) UpdateComment(commentID string, content string) (*repository.Comment, error) {
	var comment repository.Comment
	result := r.db.Model(&comment).Where("id = ?", commentID).Update("content", content)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comment, nil
}

func (r *MediaRepository) GetMediaRatings(mediaID, limit, page int) ([]*repository.Rating, int, error) {
	offset := (page - 1) * limit

	var ratings []*repository.Rating
	var count int64
	result := r.db.
		Model(&repository.Rating{}).
		Where("media_id = ?", mediaID).
		Count(&count).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&ratings)

	if result.Error != nil {
		return nil, 0, result.Error
	}
	return ratings, int(count), nil
}

func (r *MediaRepository) GetUserMediaRating(userID string, mediaID int) (*repository.Rating, error) {
	var rating repository.Rating
	result := r.db.Model(&repository.Rating{}).Where("user_id = ? AND media_id = ?", userID, mediaID).First(&rating)
	if result.Error != nil {
		return nil, result.Error
	}
	return &rating, nil
}

func (r *MediaRepository) GetUserRatings(userID string, limit, page int) ([]*repository.Rating, int, error) {
	offset := (page - 1) * limit

	var ratings []*repository.Rating
	var count int64
	result := r.db.
		Model(&repository.Rating{}).
		Where("user_id = ?", userID).
		Count(&count).
		Limit(limit).
		Offset(offset).
		Find(&ratings)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return ratings, int(count), nil
}

func (r *MediaRepository) SaveMediaRating(mediaID int, userID string, rating int) (*repository.Rating, error) {
	ratingEntity, err := r.GetUserMediaRating(userID, mediaID)
	if err != nil {
		ratingEntity = &repository.Rating{
			UserID:  userID,
			MediaID: mediaID,
			Rating:  rating,
		}
	} else {
		ratingEntity.Rating = rating
	}
	result := r.db.Save(ratingEntity)
	if result.Error != nil {
		return nil, result.Error
	}
	return ratingEntity, nil
}
