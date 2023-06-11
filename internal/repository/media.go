package repository

import (
	"errors"
	"github.com/bingemate/media-go-pkg/repository"
	"github.com/bingemate/media-go-pkg/tmdb"
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

//// GetMediaRating returns the average rating and the number of ratings for a media given the mediaID (TMDB ID)
//func (r *MediaRepository) GetMediaRating(mediaID int) (float32, int, error) {
//	var (
//		sum   float32
//		count int64
//	)
//
//	err := r.db.Model(&repository.Rating{}).Where("media_id = ?", mediaID).Count(&count).Error
//	if err != nil {
//		return 0, 0, err
//	}
//	if count == 0 {
//		return 0, 0, errors.New("no rating found")
//	}
//	err = r.db.Table("ratings").Select("SUM(rating) as total").Where("media_id = ?", mediaID).Scan(&sum).Error
//	if err != nil {
//		return 0, 0, err
//	}
//	return sum / float32(count), int(count), nil
//}

// GetMovieRating returns the average rating and the number of ratings for a movie given the mediaID (TMDB ID)
func (r *MediaRepository) GetMovieRating(movieID int) (float32, int, error) {
	var (
		sum   float32
		count int64
	)

	err := r.db.Model(&repository.MovieRating{}).Where("movie_id = ?", movieID).Count(&count).Error
	if err != nil {
		return 0, 0, err
	}
	if count == 0 {
		return 0, 0, errors.New("no rating found")
	}
	err = r.db.Table("movie_ratings").Select("SUM(rating) as total").Where("movie_id = ?", movieID).Scan(&sum).Error
	if err != nil {
		return 0, 0, err
	}
	return sum / float32(count), int(count), nil
}

// GetTvShowRating returns the average rating and the number of ratings for a tv show given the mediaID (TMDB ID)
func (r *MediaRepository) GetTvShowRating(tvShowID int) (float32, int, error) {
	var (
		sum   float32
		count int64
	)

	err := r.db.Model(&repository.TvShowRating{}).Where("tv_show_id = ?", tvShowID).Count(&count).Error
	if err != nil {
		return 0, 0, err
	}
	if count == 0 {
		return 0, 0, errors.New("no rating found")
	}
	err = r.db.Table("tv_show_ratings").Select("SUM(rating) as total").Where("tv_show_id = ?", tvShowID).Scan(&sum).Error
	if err != nil {
		return 0, 0, err
	}
	return sum / float32(count), int(count), nil
}

//// GetMedia returns a media given the mediaID (TMDB ID)
//func (r *MediaRepository) GetMedia(mediaID int) (*repository.Media, error) {
//	var media repository.Media
//	err := r.db.Where("id = ?", mediaID).First(&media).Error
//	if err != nil {
//		return nil, err
//	}
//	return &media, nil
//}

// GetMovie returns a movie given the movieID (TMDB ID)
func (r *MediaRepository) GetMovie(movieID int) (*repository.Movie, error) {
	var movie repository.Movie
	err := r.db.
		Where("id = ?", movieID).First(&movie).Error
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

// GetTvShow returns a tv show given the tvShowID (TMDB ID)
func (r *MediaRepository) GetTvShow(tvShowID int) (*repository.TvShow, error) {
	var tvShow repository.TvShow
	err := r.db.
		Where("id = ?", tvShowID).First(&tvShow).Error
	if err != nil {
		return nil, err
	}
	return &tvShow, nil
}

// GetEpisode returns an episode given the episodeID (TMDB ID)
func (r *MediaRepository) GetEpisode(episodeID int) (*repository.Episode, error) {
	var episode repository.Episode
	err := r.db.
		Joins("TvShow").
		Where(`episodes.id = ?`, episodeID).First(&episode).Error
	if err != nil {
		return nil, err
	}
	return &episode, nil
}

// GetEpisodeFileInfo returns the file info for an episode given the episodeID (TMDB ID)
func (r *MediaRepository) GetEpisodeFileInfo(episodeID int) (*repository.MediaFile, error) {
	var episode repository.Episode
	err := r.db.
		Joins("MediaFile").
		Preload("MediaFile.Audios").
		Preload("MediaFile.Subtitles").
		Where("episodes.id = ?", episodeID).
		First(&episode).Error
	if err != nil {
		return nil, err
	}
	return episode.MediaFile, nil
}

// GetMovieFileInfo returns the file info for a movie given the movieID (TMDB ID)
func (r *MediaRepository) GetMovieFileInfo(movieID int) (*repository.MediaFile, error) {
	var mediaFile repository.Movie
	err := r.db.
		Joins("MediaFile").
		Preload("MediaFile.Audios").
		Preload("MediaFile.Subtitles").
		Where("movies.id = ?", movieID).
		First(&mediaFile).Error
	if err != nil {
		return nil, err
	}
	return mediaFile.MediaFile, nil
}

//// IsMediaPresent returns true if the media is present in the database
//func (r *MediaRepository) IsMediaPresent(mediaID int) bool {
//	var count int64
//	r.db.Model(&repository.Media{}).Where("id = ?", mediaID).Count(&count)
//	return count > 0
//}

// IsMoviePresent returns true if the movie is present in the database
func (r *MediaRepository) IsMoviePresent(movieID int) bool {
	var count int64
	r.db.Model(&repository.Movie{}).Where("id = ?", movieID).Count(&count)
	return count > 0
}

// IsTvShowPresent returns true if the tv show is present in the database
func (r *MediaRepository) IsTvShowPresent(tvShowID int) bool {
	var count int64
	r.db.Model(&repository.TvShow{}).Where("id = ?", tvShowID).Count(&count)
	return count > 0
}

// IsEpisodePresent returns true if the episode is present in the database
func (r *MediaRepository) IsEpisodePresent(episodeID int) bool {
	var count int64
	r.db.Model(&repository.Episode{}).Where("id = ?", episodeID).Count(&count)
	return count > 0
}

// IsMovieFilePresent returns true if the movie file is present in the database
func (r *MediaRepository) IsMovieFilePresent(movieID int) bool {
	var count int64
	r.db.Model(&repository.Movie{}).Where("id = ? AND media_file_id IS NOT NULL", movieID).Count(&count)
	return count > 0
}

// IsEpisodeFilePresent returns true if the episode file is present in the database
func (r *MediaRepository) IsEpisodeFilePresent(episodeID int) bool {
	var count int64
	r.db.Model(&repository.Episode{}).Where("id = ? AND media_file_id IS NOT NULL", episodeID).Count(&count)
	return count > 0
}

// IsTvShowHasEpisodeFiles returns true if the tv show has episode files in the database
func (r *MediaRepository) IsTvShowHasEpisodeFiles(tvShowID int) bool {
	var count int64
	r.db.Model(&repository.Episode{}).Where("tv_show_id = ? AND media_file_id IS NOT NULL", tvShowID).Count(&count)
	return count > 0
}

// GetAvailableMoviesByRating returns a list of movies ordered by rating
// Return also the total number of results
func (r *MediaRepository) GetAvailableMoviesByRating(page, limit, days int) ([]repository.Movie, int, error) {
	var movies []repository.Movie
	offset := (page - 1) * limit

	var count int64
	result := r.db.Table("movies").
		Select("movies.*, AVG(movie_ratings.rating) as average_rating").
		Joins("LEFT JOIN movie_ratings ON movie_ratings.movie_id = movies.id AND movie_ratings.created_at > ?", time.Now().AddDate(0, 0, -days)).
		Where("movies.media_file_id IS NOT NULL").
		Group("movies.id").
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

// GetAvailableTvShowsByRating returns a list of tv shows ordered by rating
// Return also the total number of results
func (r *MediaRepository) GetAvailableTvShowsByRating(page, limit, days int) ([]repository.TvShow, int, error) {
	var tvShows []repository.TvShow
	offset := (page - 1) * limit
	var count int64
	result := r.db.Table("tv_shows").
		Select("tv_shows.*, AVG(tv_show_ratings.rating) as average_rating").
		Joins("LEFT JOIN tv_show_ratings ON tv_show_ratings.tv_show_id = tv_shows.id AND tv_show_ratings.created_at > ?", time.Now().AddDate(0, 0, -days)).
		Joins("JOIN episodes ON episodes.tv_show_id = tv_shows.id").
		Where("episodes.media_file_id IS NOT NULL").
		Group("tv_shows.id").
		Having("COUNT(DISTINCT episodes.id) > 0").
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

// SearchAvailableMovies returns a list of movies matching the search query
// Return also the total number of results
// Results are ordered by pertinence and / or rating
func (r *MediaRepository) SearchAvailableMovies(page, limit int, query string) ([]repository.Movie, int, error) {
	var movies []repository.Movie
	offset := (page - 1) * limit

	var count int64
	result := r.db.Table("movies").
		Select("movies.*, AVG(movie_ratings.rating) as average_rating").
		Joins("LEFT JOIN movie_ratings ON movie_ratings.movie_id = movies.id").
		Where("movies.media_file_id IS NOT NULL AND movies.name ILIKE ?", "%"+query+"%").
		Group("movies.id").
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

// SearchAvailableTvShows returns a list of tv shows matching the search query
// Return also the total number of results
// Results are ordered by pertinence and / or rating
func (r *MediaRepository) SearchAvailableTvShows(page, limit int, query string) ([]repository.TvShow, int, error) {
	var tvShows []repository.TvShow
	offset := (page - 1) * limit

	var count int64
	result := r.db.Table("tv_shows").
		Select("tv_shows.*, AVG(tv_show_ratings.rating) as average_rating").
		Joins("LEFT JOIN tv_show_ratings ON tv_show_ratings.tv_show_id = tv_shows.id").
		Joins("JOIN episodes ON episodes.tv_show_id = tv_shows.id").
		Where("tv_shows.name ILIKE ?", "%"+query+"%").
		Where("episodes.media_file_id IS NOT NULL").
		Group("tv_shows.id").
		Having("COUNT(DISTINCT episodes.id) > 0").
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

// GetAvailableRecentMovies returns a list of recently added movies
func (r *MediaRepository) GetAvailableRecentMovies(page, limit int) ([]repository.Movie, int, error) {
	var movies []repository.Movie
	offset := (page - 1) * limit
	var count int64
	result := r.db.Table("movies").
		Select("*").
		Where("movies.media_file_id IS NOT NULL").
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

// GetAvailableRecentTvShows returns a list of recently added tv shows
func (r *MediaRepository) GetAvailableRecentTvShows(page, limit int) ([]repository.TvShow, int, error) {
	var tvShows []repository.TvShow
	offset := (page - 1) * limit
	var count int64
	result := r.db.Table("tv_shows").
		Joins("JOIN episodes ON episodes.tv_show_id = tv_shows.id").
		Where("episodes.media_file_id IS NOT NULL").
		Group("tv_shows.id").
		Having("COUNT(DISTINCT episodes.id) > 0").
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

//// GetMediasByComments returns a list of medias ordered by number of comments
//func (r *MediaRepository) GetMediasByComments(present bool) (*[]int, error) {
//	var mediaIds []int
//
//	query := r.db.Model(&repository.Comment{}).
//		Select("media_id").
//		Group("media_id").
//		Order("COUNT(comments.id) DESC").
//		Limit(20)
//
//	if present {
//		// Add WHERE clause to check if media_id exists in the "media" table
//		query = query.Joins("JOIN media ON comments.media_id = media.id").Where("media.id = comments.media_id")
//	}
//
//	err := query.Find(&mediaIds).Error
//
//	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
//		return nil, err
//	}
//
//	return &mediaIds, nil
//}

// GetMoviesByComments returns a list of movies ordered by number of comments
func (r *MediaRepository) GetMoviesByComments(present bool) (*[]int, error) {
	var movieIds []int

	query := r.db.Model(&repository.MovieComment{}).
		Select("movie_id").
		Group("movie_id").
		Order("COUNT(movie_comments.id) DESC").
		Limit(20)

	if present {
		// Add WHERE clause to check if movie_id exists in the "movie" table
		query = query.Joins("JOIN movies ON movie_comments.movie_id = movies.id").Where("movies.id = movie_comments.movie_id AND movies.media_file_id IS NOT NULL")
	}

	err := query.Find(&movieIds).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &movieIds, nil
}

// GetTvShowsByComments returns a list of tv shows ordered by number of comments
func (r *MediaRepository) GetTvShowsByComments(present bool) (*[]int, error) {
	var tvShowIds []int

	query := r.db.Model(&repository.TvShowComment{}).
		Select("tv_show_comments.tv_show_id").
		Group("tv_show_comments.tv_show_id").
		Order("COUNT(tv_show_comments.id) DESC").
		Limit(20)

	if present {
		// Add WHERE clause to check if tv_show_id exists in the "tv_show" table
		query = query.Joins("JOIN tv_shows ON tv_show_comments.tv_show_id = tv_shows.id").
			Joins("JOIN episodes ON episodes.tv_show_id = tv_shows.id").
			Having("COUNT(DISTINCT episodes.id) > 0").
			Where("tv_shows.id = tv_show_comments.tv_show_id")
	}

	err := query.Find(&tvShowIds).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &tvShowIds, nil
}

//
//func (r *MediaRepository) GetFollowedReleases(userID string, month int) (*[]int, error) {
//	var followedReleases []int
//	result := r.db.Table("watch_list_item").
//		Select("media_id").
//		Where("user_id = ? AND status != ?", userID, repository.WatchListStatusAbandoned).
//		Find(&followedReleases)
//
//	if result.Error != nil {
//		return nil, result.Error
//	}
//
//	return &followedReleases, nil
//}

// GetFollowedMoviesReleases returns a list of followed movies releases
func (r *MediaRepository) GetFollowedMoviesReleases(userID string, month int) (*[]int, error) {
	var followedMoviesReleases []int
	result := r.db.Table("movie_watch_list_item").
		Select("movie_id").
		Where("user_id = ? AND status != ?", userID, repository.WatchListStatusAbandoned).
		Find(&followedMoviesReleases)

	if result.Error != nil {
		return nil, result.Error
	}

	return &followedMoviesReleases, nil
}

// GetFollowedTvShowsReleases returns a list of followed tv shows releases
func (r *MediaRepository) GetFollowedTvShowsReleases(userID string, month int) (*[]int, error) {
	var followedTvShowsReleases []int
	result := r.db.Table("tv_show_watch_list_item").
		Select("tv_show_id").
		Where("user_id = ? AND status != ?", userID, repository.WatchListStatusAbandoned).
		Find(&followedTvShowsReleases)

	if result.Error != nil {
		return nil, result.Error
	}

	return &followedTvShowsReleases, nil
}

//func (r *MediaRepository) GetMediaComments(mediaID, size, page int) ([]*repository.Comment, int, error) {
//	var comments []*repository.Comment
//	var count int64
//	offset := (page - 1) * size
//	result := r.db.Model(&repository.Comment{}).
//		Where("media_id = ?", mediaID).
//		Count(&count).
//		Order("created_at DESC").
//		Offset(offset).
//		Limit(size).
//		Find(&comments)
//
//	if result.Error != nil {
//		return nil, 0, result.Error
//	}
//	return comments, int(count), nil
//}

// GetMovieComments returns a list of comments for a movie
func (r *MediaRepository) GetMovieComments(movieID, size, page int) ([]*repository.MovieComment, int, error) {
	var comments []*repository.MovieComment
	var count int64
	offset := (page - 1) * size
	result := r.db.Model(&repository.MovieComment{}).
		Where("movie_id = ?", movieID).
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

// GetTvShowComments returns a list of comments for a tv show
func (r *MediaRepository) GetTvShowComments(tvShowID, size, page int) ([]*repository.TvShowComment, int, error) {
	var comments []*repository.TvShowComment
	var count int64
	offset := (page - 1) * size
	result := r.db.Model(&repository.TvShowComment{}).
		Where("tv_show_id = ?", tvShowID).
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

//func (r *MediaRepository) GetUserComments(userID string, size, page int) ([]*repository.Comment, int, error) {
//	var comments []*repository.Comment
//	var count int64
//	offset := (page - 1) * size
//	result := r.db.Model(&repository.Comment{}).
//		Where("user_id = ?", userID).
//		Count(&count).
//		Order("created_at DESC").
//		Offset(offset).
//		Limit(size).
//		Find(&comments)
//
//	if result.Error != nil {
//		return nil, 0, result.Error
//	}
//	return comments, int(count), nil
//}

// GetUserMovieComments returns a list of comments for a movie
func (r *MediaRepository) GetUserMovieComments(userID string, size, page int) ([]*repository.MovieComment, int, error) {
	var comments []*repository.MovieComment
	var count int64
	offset := (page - 1) * size
	result := r.db.Model(&repository.MovieComment{}).
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

// GetUserTvShowComments returns a list of comments for a tv show
func (r *MediaRepository) GetUserTvShowComments(userID string, size, page int) ([]*repository.TvShowComment, int, error) {
	var comments []*repository.TvShowComment
	var count int64
	offset := (page - 1) * size
	result := r.db.Model(&repository.TvShowComment{}).
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

// GetUserMovieCommentsByRange returns a list of comments for a movie
func (r *MediaRepository) GetUserMovieCommentsByRange(userID string, start, end time.Time) ([]*repository.MovieComment, error) {
	var comments []*repository.MovieComment
	result := r.db.Model(&repository.MovieComment{}).
		Where("user_id = ? AND created_at BETWEEN ? AND ?", userID, start, end).
		Order("created_at DESC").
		Find(&comments)

	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}

// GetUserTvShowCommentsByRange returns a list of comments for a tv show
func (r *MediaRepository) GetUserTvShowCommentsByRange(userID string, start, end time.Time) ([]*repository.TvShowComment, error) {
	var comments []*repository.TvShowComment
	result := r.db.Model(&repository.TvShowComment{}).
		Where("user_id = ? AND created_at BETWEEN ? AND ?", userID, start, end).
		Order("created_at DESC").
		Find(&comments)

	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}

// GetMovieCommentsByRange returns a list of comments for a movie
func (r *MediaRepository) GetMovieCommentsByRange(start, end time.Time) ([]*repository.MovieComment, error) {
	var comments []*repository.MovieComment
	result := r.db.Model(&repository.MovieComment{}).
		Where("created_at BETWEEN ? AND ?", start, end).
		Order("created_at DESC").
		Find(&comments)

	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}

// GetTvShowCommentsByRange returns a list of comments for a tv show
func (r *MediaRepository) GetTvShowCommentsByRange(start, end time.Time) ([]*repository.TvShowComment, error) {
	var comments []*repository.TvShowComment
	result := r.db.Model(&repository.TvShowComment{}).
		Where("created_at BETWEEN ? AND ?", start, end).
		Order("created_at DESC").
		Find(&comments)

	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}

// CountUserMovieComments returns the number of comments for a movie
func (r *MediaRepository) CountUserMovieComments(userID string) (int, error) {
	var count int64
	result := r.db.Model(&repository.MovieComment{}).
		Where("user_id = ?", userID).
		Count(&count)

	if result.Error != nil {
		return 0, result.Error
	}
	return int(count), nil
}

// CountUserTvShowComments returns the number of comments for a tv show
func (r *MediaRepository) CountUserTvShowComments(userID string) (int, error) {
	var count int64
	result := r.db.Model(&repository.TvShowComment{}).
		Where("user_id = ?", userID).
		Count(&count)

	if result.Error != nil {
		return 0, result.Error
	}
	return int(count), nil
}

// CountMovieComments returns the number of comments for a movie
func (r *MediaRepository) CountMovieComments() (int, error) {
	var count int64
	result := r.db.Model(&repository.MovieComment{}).
		Count(&count)

	if result.Error != nil {
		return 0, result.Error
	}
	return int(count), nil
}

// CountTvShowComments returns the number of comments for a tv show
func (r *MediaRepository) CountTvShowComments() (int, error) {
	var count int64
	result := r.db.Model(&repository.TvShowComment{}).
		Count(&count)

	if result.Error != nil {
		return 0, result.Error
	}
	return int(count), nil
}

//func (r *MediaRepository) AddComment(userID string, mediaID int, content string) (*repository.Comment, error) {
//	comment := repository.Comment{
//		UserID:  userID,
//		MediaID: mediaID,
//		Content: content,
//	}
//	result := r.db.Create(&comment)
//	if result.Error != nil {
//		return nil, result.Error
//	}
//	return &comment, nil
//}

// AddMovieComment adds a comment to a movie
func (r *MediaRepository) AddMovieComment(userID string, movieID int, content string) (*repository.MovieComment, error) {
	comment := repository.MovieComment{
		UserID:  userID,
		MovieID: movieID,
		Content: content,
	}
	result := r.db.Create(&comment)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comment, nil
}

// AddTvShowComment adds a comment to a tv show
func (r *MediaRepository) AddTvShowComment(userID string, tvShowID int, content string) (*repository.TvShowComment, error) {
	comment := repository.TvShowComment{
		UserID:   userID,
		TvShowID: tvShowID,
		Content:  content,
	}
	result := r.db.Create(&comment)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comment, nil
}

//func (r *MediaRepository) GetComment(commentID string) (*repository.Comment, error) {
//	var comment repository.Comment
//	result := r.db.Model(&repository.Comment{}).Where("id = ?", commentID).First(&comment)
//	if result.Error != nil {
//		return nil, result.Error
//	}
//	return &comment, nil
//}

// GetMovieComment returns a movie comment
func (r *MediaRepository) GetMovieComment(commentID string) (*repository.MovieComment, error) {
	var comment repository.MovieComment
	result := r.db.Model(&repository.MovieComment{}).Where("id = ?", commentID).First(&comment)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comment, nil
}

// GetTvShowComment returns a tv show comment
func (r *MediaRepository) GetTvShowComment(commentID string) (*repository.TvShowComment, error) {
	var comment repository.TvShowComment
	result := r.db.Model(&repository.TvShowComment{}).Where("id = ?", commentID).First(&comment)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comment, nil
}

//func (r *MediaRepository) DeleteComment(commentID string) error {
//	return r.db.Where("id = ?", commentID).Delete(&repository.Comment{}).Error
//}

// DeleteMovieComment deletes a movie comment
func (r *MediaRepository) DeleteMovieComment(commentID string) error {
	return r.db.Where("id = ?", commentID).Delete(&repository.MovieComment{}).Error
}

// DeleteTvShowComment deletes a tv show comment
func (r *MediaRepository) DeleteTvShowComment(commentID string) error {
	return r.db.Where("id = ?", commentID).Delete(&repository.TvShowComment{}).Error
}

//func (r *MediaRepository) UpdateComment(commentID string, content string) (*repository.Comment, error) {
//	var comment repository.Comment
//	result := r.db.Model(&comment).Where("id = ?", commentID).Update("content", content)
//	if result.Error != nil {
//		return nil, result.Error
//	}
//	return &comment, nil
//}

// UpdateMovieComment updates a movie comment
func (r *MediaRepository) UpdateMovieComment(commentID string, content string) (*repository.MovieComment, error) {
	var comment repository.MovieComment
	result := r.db.Model(&comment).Where("id = ?", commentID).Update("content", content)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comment, nil
}

// UpdateTvShowComment updates a tv show comment
func (r *MediaRepository) UpdateTvShowComment(commentID string, content string) (*repository.TvShowComment, error) {
	var comment repository.TvShowComment
	result := r.db.Model(&comment).Where("id = ?", commentID).Update("content", content)
	if result.Error != nil {
		return nil, result.Error
	}
	return &comment, nil
}

//func (r *MediaRepository) GetMediaRatings(mediaID, limit, page int) ([]*repository.Rating, int, error) {
//	offset := (page - 1) * limit
//
//	var ratings []*repository.Rating
//	var count int64
//	result := r.db.
//		Model(&repository.Rating{}).
//		Where("media_id = ?", mediaID).
//		Count(&count).
//		Order("created_at DESC").
//		Offset(offset).
//		Limit(limit).
//		Find(&ratings)
//
//	if result.Error != nil {
//		return nil, 0, result.Error
//	}
//	return ratings, int(count), nil
//}

// GetMovieRatings returns movie ratings
func (r *MediaRepository) GetMovieRatings(movieID, limit, page int) ([]*repository.MovieRating, int, error) {
	offset := (page - 1) * limit

	var ratings []*repository.MovieRating
	var count int64
	result := r.db.
		Model(&repository.MovieRating{}).
		Where("movie_id = ?", movieID).
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

// GetTvShowRatings returns tv show ratings
func (r *MediaRepository) GetTvShowRatings(tvShowID, limit, page int) ([]*repository.TvShowRating, int, error) {
	offset := (page - 1) * limit

	var ratings []*repository.TvShowRating
	var count int64
	result := r.db.
		Model(&repository.TvShowRating{}).
		Where("tv_show_id = ?", tvShowID).
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

//func (r *MediaRepository) GetUserMediaRating(userID string, mediaID int) (*repository.Rating, error) {
//	var rating repository.Rating
//	result := r.db.Model(&repository.Rating{}).Where("user_id = ? AND media_id = ?", userID, mediaID).First(&rating)
//	if result.Error != nil {
//		return nil, result.Error
//	}
//	return &rating, nil
//}

// GetUserMovieRating returns a user's movie rating
func (r *MediaRepository) GetUserMovieRating(userID string, movieID int) (*repository.MovieRating, error) {
	var rating repository.MovieRating
	result := r.db.Model(&repository.MovieRating{}).Where("user_id = ? AND movie_id = ?", userID, movieID).First(&rating)
	if result.Error != nil {
		return nil, result.Error
	}
	return &rating, nil
}

// GetUserTvShowRating returns a user's tv show rating
func (r *MediaRepository) GetUserTvShowRating(userID string, tvShowID int) (*repository.TvShowRating, error) {
	var rating repository.TvShowRating
	result := r.db.Model(&repository.TvShowRating{}).Where("user_id = ? AND tv_show_id = ?", userID, tvShowID).First(&rating)
	if result.Error != nil {
		return nil, result.Error
	}
	return &rating, nil
}

//func (r *MediaRepository) GetUserRatings(userID string, limit, page int) ([]*repository.Rating, int, error) {
//	offset := (page - 1) * limit
//
//	var ratings []*repository.Rating
//	var count int64
//	result := r.db.
//		Model(&repository.Rating{}).
//		Where("user_id = ?", userID).
//		Count(&count).
//		Limit(limit).
//		Offset(offset).
//		Find(&ratings)
//	if result.Error != nil {
//		return nil, 0, result.Error
//	}
//	return ratings, int(count), nil
//}

// GetUserMovieRatings returns a user's movie ratings
func (r *MediaRepository) GetUserMovieRatings(userID string, limit, page int) ([]*repository.MovieRating, int, error) {
	offset := (page - 1) * limit

	var ratings []*repository.MovieRating
	var count int64
	result := r.db.
		Model(&repository.MovieRating{}).
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

// GetUserTvShowRatings returns a user's tv show ratings
func (r *MediaRepository) GetUserTvShowRatings(userID string, limit, page int) ([]*repository.TvShowRating, int, error) {
	offset := (page - 1) * limit

	var ratings []*repository.TvShowRating
	var count int64
	result := r.db.
		Model(&repository.TvShowRating{}).
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

//func (r *MediaRepository) SaveMediaRating(mediaID int, userID string, rating int) (*repository.Rating, error) {
//	ratingEntity, err := r.GetUserMediaRating(userID, mediaID)
//	if err != nil {
//		ratingEntity = &repository.Rating{
//			UserID:  userID,
//			MediaID: mediaID,
//			Rating:  rating,
//		}
//	} else {
//		ratingEntity.Rating = rating
//	}
//	result := r.db.Save(ratingEntity)
//	if result.Error != nil {
//		return nil, result.Error
//	}
//	return ratingEntity, nil
//}

// SaveMovieRating saves a movie rating
func (r *MediaRepository) SaveMovieRating(movieID int, userID string, rating int) (*repository.MovieRating, error) {
	ratingEntity, err := r.GetUserMovieRating(userID, movieID)
	if err != nil {
		ratingEntity = &repository.MovieRating{
			UserID:  userID,
			MovieID: movieID,
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

// SaveTvShowRating saves a tv show rating
func (r *MediaRepository) SaveTvShowRating(tvShowID int, userID string, rating int) (*repository.TvShowRating, error) {
	ratingEntity, err := r.GetUserTvShowRating(userID, tvShowID)
	if err != nil {
		ratingEntity = &repository.TvShowRating{
			UserID:   userID,
			TvShowID: tvShowID,
			Rating:   rating,
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

func (r *MediaRepository) SaveMovie(movie *tmdb.Movie) error {
	if r.IsMoviePresent(movie.ID) {
		return nil
	}
	releaseDate, err := time.Parse("2006-01-02", movie.ReleaseDate)
	if err != nil {
		return err
	}

	movieEntity := &repository.Movie{
		ID:          movie.ID,
		Name:        movie.Title,
		ReleaseDate: releaseDate,
	}
	err = r.db.Save(movieEntity).Error
	if err != nil {
		return err
	}
	movieEntity.Categories = *r.extractCategories(&movie.Genres)
	err = r.db.Save(movieEntity).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *MediaRepository) SaveTvShow(tvShow *tmdb.TVShow) error {
	if r.IsTvShowPresent(tvShow.ID) {
		return nil
	}
	releaseDate, err := time.Parse("2006-01-02", tvShow.ReleaseDate)
	if err != nil {
		return err
	}

	tvShowEntity := &repository.TvShow{
		ID:          tvShow.ID,
		Name:        tvShow.Title,
		ReleaseDate: releaseDate,
	}
	err = r.db.Save(tvShowEntity).Error
	if err != nil {
		return err
	}
	tvShowEntity.Categories = *r.extractCategories(&tvShow.Genres)
	err = r.db.Save(tvShowEntity).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *MediaRepository) SaveEpisode(episode *tmdb.TVEpisode) error {
	if r.IsEpisodePresent(episode.ID) {
		return nil
	}
	if !r.IsTvShowPresent(episode.TVShowID) {
		return nil
	}
	releaseDate, err := time.Parse("2006-01-02", episode.AirDate)
	if err != nil {
		return err
	}
	episodeEntity := &repository.Episode{
		ID:          episode.ID,
		Name:        episode.Name,
		TvShowID:    episode.TVShowID,
		NbSeason:    episode.SeasonNumber,
		NbEpisode:   episode.EpisodeNumber,
		ReleaseDate: releaseDate,
	}
	return r.db.Save(episodeEntity).Error
}

func (r *MediaRepository) extractCategories(pkgCategories *[]tmdb.Genre) *[]repository.Category {
	var categories = make([]repository.Category, len(*pkgCategories))
	for i, c := range *pkgCategories {
		InDB, err := r.getOrCreateCategory(c.Name)
		if err != nil {
			categories[i] = repository.Category{
				Name: c.Name,
			}
		} else {
			categories[i] = *InDB
		}
	}
	return &categories
}

func (r *MediaRepository) getOrCreateCategory(name string) (*repository.Category, error) {
	var category repository.Category
	db := r.db.Where("name = ?", name).First(&category)
	if db.Error != nil && !errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return nil, db.Error
	}
	if category.ID != "" {
		return &category, nil
	}
	category = repository.Category{
		Name: name,
	}
	db = r.db.Save(&category)
	if db.Error != nil {
		return nil, db.Error
	}
	return &category, nil
}

func (r *MediaRepository) AvailableEpisodes(tvShowID int) (*[]int, error) {
	var episodeIDs []int

	result := r.db.Model(&repository.Episode{}).
		Select("id").
		Where("tv_show_id = ? AND media_file_id IS NOT NULL", tvShowID).
		Order("nb_season, nb_episode").
		Find(&episodeIDs)
	if result.Error != nil {
		return nil, result.Error
	}
	return &episodeIDs, nil
}

func (r *MediaRepository) CountUserMovieRatings(userID string) (int, error) {
	var count int64
	result := r.db.Model(&repository.MovieRating{}).
		Where("user_id = ?", userID).
		Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(count), nil
}

func (r *MediaRepository) CountUserTvShowRatings(userID string) (int, error) {
	var count int64
	result := r.db.Model(&repository.TvShowRating{}).
		Where("user_id = ?", userID).
		Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(count), nil
}

func (r *MediaRepository) CountMovieRatings() (int, error) {
	var count int64
	result := r.db.Model(&repository.MovieRating{}).
		Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(count), nil
}

func (r *MediaRepository) CountTvShowRatings() (int, error) {
	var count int64
	result := r.db.Model(&repository.TvShowRating{}).
		Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(count), nil
}
