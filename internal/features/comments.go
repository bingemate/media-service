package features

import (
	"fmt"
	repository2 "github.com/bingemate/media-go-pkg/repository"
	"github.com/bingemate/media-service/internal/repository"
)

type CommentService struct {
	mediaRepository *repository.MediaRepository
}

func NewCommentService(mediaRepository *repository.MediaRepository) *CommentService {
	return &CommentService{mediaRepository}
}

func (s *CommentService) GetComments(mediaID, page int) ([]*repository2.Comment, int, error) {
	return s.mediaRepository.GetMediaComments(mediaID, 5, page)
}

func (s *CommentService) GetUserComments(userID string, page int) ([]*repository2.Comment, int, error) {
	return s.mediaRepository.GetUserComments(userID, 5, page)
}

func (s *CommentService) AddComment(userID string, mediaID int, comment string) (*repository2.Comment, error) {
	return s.mediaRepository.AddComment(userID, mediaID, comment)
}

func (s *CommentService) DeleteComment(commentID, userID string, isAdmin bool) error {
	comment, err := s.mediaRepository.GetComment(commentID)
	if err != nil {
		return err
	}
	if comment.UserID != userID && !isAdmin {
		return fmt.Errorf("you are not allowed to delete this comment")
	}

	return s.mediaRepository.DeleteComment(commentID)
}

func (s *CommentService) UpdateComment(commentID, userID string, isAdmin bool, content string) (*repository2.Comment, error) {
	comment, err := s.mediaRepository.GetComment(commentID)
	if err != nil {
		return nil, err
	}
	if comment.UserID != userID && !isAdmin {
		return nil, fmt.Errorf("you are not allowed to update this comment")
	}

	return s.mediaRepository.UpdateComment(commentID, content)
}
