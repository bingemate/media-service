package features

import (
	repository2 "github.com/bingemate/media-go-pkg/repository"
	"github.com/bingemate/media-service/internal/repository"
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
	return s.mediaRepository.GetUserMediaRating(userID, mediaID)
}

func (s *RatingService) RateMedia(userID string, mediaID, rating int) (*repository2.Rating, error) {
	return s.mediaRepository.SaveMediaRating(mediaID, userID, rating)
}
