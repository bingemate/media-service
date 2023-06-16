package features

import (
	"errors"
	repository2 "github.com/bingemate/media-go-pkg/repository"
	"github.com/bingemate/media-service/internal/repository"
	"gorm.io/gorm"
)

type RatingService struct {
	mediaRepository *repository.MediaRepository
}

func NewRatingService(mediaRepository *repository.MediaRepository) *RatingService {
	return &RatingService{mediaRepository}
}

//func (s *RatingService) GetMediaRating(mediaID, page int) ([]*repository2.Rating, int, error) {
//	return s.mediaRepository.GetMediaRatings(mediaID, 10, page)
//}

func (s *RatingService) GetMovieRatings(movieID, page int) ([]*repository2.MovieRating, int, error) {
	return s.mediaRepository.GetMovieRatings(movieID, 10, page)
}

func (s *RatingService) GetTvShowRatings(tvShowID, page int) ([]*repository2.TvShowRating, int, error) {
	return s.mediaRepository.GetTvShowRatings(tvShowID, 10, page)
}

//func (s *RatingService) GetUsersRating(userID string, page int) ([]*repository2.Rating, int, error) {
//	return s.mediaRepository.GetUserRatings(userID, 10, page)
//}

func (s *RatingService) GetUsersMovieRatings(userID string, page int) ([]*repository2.MovieRating, int, error) {
	return s.mediaRepository.GetUserMovieRatings(userID, 10, page)
}

func (s *RatingService) GetUsersTvShowRatings(userID string, page int) ([]*repository2.TvShowRating, int, error) {
	return s.mediaRepository.GetUserTvShowRatings(userID, 10, page)
}

//func (s *RatingService) GetUserMediaRating(userID string, mediaID int) (*repository2.Rating, error) {
//	rating, err := s.mediaRepository.GetUserMediaRating(userID, mediaID)
//	if err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return &repository2.Rating{Rating: 0, UserID: userID, MediaID: mediaID}, nil
//		}
//		return nil, err
//	}
//	return rating, nil
//}

func (s *RatingService) GetUserMovieRating(userID string, movieID int) (*repository2.MovieRating, error) {
	rating, err := s.mediaRepository.GetUserMovieRating(userID, movieID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &repository2.MovieRating{Rating: 0, UserID: userID, MovieID: movieID}, nil
		}
		return nil, err
	}
	return rating, nil
}

func (s *RatingService) GetUserTvShowRating(userID string, tvShowID int) (*repository2.TvShowRating, error) {
	rating, err := s.mediaRepository.GetUserTvShowRating(userID, tvShowID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &repository2.TvShowRating{Rating: 0, UserID: userID, TvShowID: tvShowID}, nil
		}
		return nil, err
	}
	return rating, nil
}

//func (s *RatingService) RateMedia(userID string, mediaID, rating int) (*repository2.Rating, error) {
//	return s.mediaRepository.SaveMediaRating(mediaID, userID, rating)
//}

func (s *RatingService) RateMovie(userID string, movieID, rating int) (*repository2.MovieRating, error) {
	return s.mediaRepository.SaveMovieRating(movieID, userID, rating)
}

func (s *RatingService) RateTvShow(userID string, tvShowID, rating int) (*repository2.TvShowRating, error) {
	return s.mediaRepository.SaveTvShowRating(tvShowID, userID, rating)
}

func (s *RatingService) CountUserRatings(userID string) (int, error) {
	movieRatingCount, err := s.mediaRepository.CountUserMovieRatings(userID)
	if err != nil {
		return 0, err
	}
	tvShowRatingCount, err := s.mediaRepository.CountUserTvShowRatings(userID)
	if err != nil {
		return 0, err
	}
	return movieRatingCount + tvShowRatingCount, nil
}

func (s *RatingService) CountRatings() (int, error) {
	movieRatingCount, err := s.mediaRepository.CountMovieRatings()
	if err != nil {
		return 0, err
	}
	tvShowRatingCount, err := s.mediaRepository.CountTvShowRatings()
	if err != nil {
		return 0, err
	}
	return movieRatingCount + tvShowRatingCount, nil
}
