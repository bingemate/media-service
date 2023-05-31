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

func (s *RatingService) GetMediaRating(mediaID, page int) ([]*repository2.Rating, int, error) {
	return s.mediaRepository.GetMediaRatings(mediaID, 10, page)
}

func (s *RatingService) GetUsersRating(userID string, page int) ([]*repository2.Rating, int, error) {
	return s.mediaRepository.GetUserRatings(userID, 10, page)
}

func (s *RatingService) GetUserMediaRating(userID string, mediaID int) (*repository2.Rating, error) {
	rating, err := s.mediaRepository.GetUserMediaRating(userID, mediaID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &repository2.Rating{Rating: 0, UserID: userID, MediaID: mediaID}, nil
		}
		return nil, err
	}
	return rating, nil
}

func (s *RatingService) RateMedia(userID string, mediaID, rating int) (*repository2.Rating, error) {
	return s.mediaRepository.SaveMediaRating(mediaID, userID, rating)
}
