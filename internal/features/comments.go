package features

import (
	"fmt"
	repository2 "github.com/bingemate/media-go-pkg/repository"
	"github.com/bingemate/media-service/internal/repository"
	"time"
)

type CommentService struct {
	mediaRepository *repository.MediaRepository
}

func NewCommentService(mediaRepository *repository.MediaRepository) *CommentService {
	return &CommentService{mediaRepository}
}

//func (s *CommentService) GetComments(mediaID, page int) ([]*repository2.Comment, int, error) {
//	return s.mediaRepository.GetMediaComments(mediaID, 5, page)
//}

func (s *CommentService) GetMovieComments(movieID, page int) ([]*repository2.MovieComment, int, error) {
	return s.mediaRepository.GetMovieComments(movieID, 5, page)
}

func (s *CommentService) GetTvShowComments(tvShowID, page int) ([]*repository2.TvShowComment, int, error) {
	return s.mediaRepository.GetTvShowComments(tvShowID, 5, page)
}

//func (s *CommentService) GetUserComments(userID string, page int) ([]*repository2.Comment, int, error) {
//	return s.mediaRepository.GetUserComments(userID, 5, page)
//}

func (s *CommentService) GetMovieUserComments(userID string, page int) ([]*repository2.MovieComment, int, error) {
	return s.mediaRepository.GetUserMovieComments(userID, 5, page)
}

func (s *CommentService) GetTvShowUserComments(userID string, page int) ([]*repository2.TvShowComment, int, error) {
	return s.mediaRepository.GetUserTvShowComments(userID, 5, page)
}

//func (s *CommentService) AddComment(userID string, mediaID int, comment string) (*repository2.Comment, error) {
//	return s.mediaRepository.AddComment(userID, mediaID, comment)
//}

func (s *CommentService) AddMovieComment(userID string, movieID int, comment string) (*repository2.MovieComment, error) {
	return s.mediaRepository.AddMovieComment(userID, movieID, comment)
}

func (s *CommentService) AddTvShowComment(userID string, tvShowID int, comment string) (*repository2.TvShowComment, error) {
	return s.mediaRepository.AddTvShowComment(userID, tvShowID, comment)
}

//func (s *CommentService) DeleteComment(commentID, userID string, isAdmin bool) error {
//	comment, err := s.mediaRepository.GetComment(commentID)
//	if err != nil {
//		return err
//	}
//	if comment.UserID != userID && !isAdmin {
//		return fmt.Errorf("you are not allowed to delete this comment")
//	}
//
//	return s.mediaRepository.DeleteComment(commentID)
//}

func (s *CommentService) DeleteMovieComment(commentID, userID string, isAdmin bool) error {
	comment, err := s.mediaRepository.GetMovieComment(commentID)
	if err != nil {
		return err
	}
	if comment.UserID != userID && !isAdmin {
		return fmt.Errorf("you are not allowed to delete this comment")
	}

	return s.mediaRepository.DeleteMovieComment(commentID)
}

func (s *CommentService) DeleteTvShowComment(commentID, userID string, isAdmin bool) error {
	comment, err := s.mediaRepository.GetTvShowComment(commentID)
	if err != nil {
		return err
	}
	if comment.UserID != userID && !isAdmin {
		return fmt.Errorf("you are not allowed to delete this comment")
	}

	return s.mediaRepository.DeleteTvShowComment(commentID)
}

//func (s *CommentService) UpdateComment(commentID, userID string, isAdmin bool, content string) (*repository2.Comment, error) {
//	comment, err := s.mediaRepository.GetComment(commentID)
//	if err != nil {
//		return nil, err
//	}
//	if comment.UserID != userID && !isAdmin {
//		return nil, fmt.Errorf("you are not allowed to update this comment")
//	}
//
//	return s.mediaRepository.UpdateComment(commentID, content)
//}

func (s *CommentService) UpdateMovieComment(commentID, userID string, isAdmin bool, content string) (*repository2.MovieComment, error) {
	comment, err := s.mediaRepository.GetMovieComment(commentID)
	if err != nil {
		return nil, err
	}
	if comment.UserID != userID && !isAdmin {
		return nil, fmt.Errorf("you are not allowed to update this comment")
	}

	return s.mediaRepository.UpdateMovieComment(commentID, content)
}

func (s *CommentService) UpdateTvShowComment(commentID, userID string, isAdmin bool, content string) (*repository2.TvShowComment, error) {
	comment, err := s.mediaRepository.GetTvShowComment(commentID)
	if err != nil {
		return nil, err
	}
	if comment.UserID != userID && !isAdmin {
		return nil, fmt.Errorf("you are not allowed to update this comment")
	}

	return s.mediaRepository.UpdateTvShowComment(commentID, content)
}

func (s *CommentService) GetUserMovieCommentsByRange(userID string, start, end string) ([]*repository2.MovieComment, error) {
	startTime, err := time.Parse("2006-01-02", start)
	if err != nil {
		return nil, err
	}
	endTime, err := time.Parse("2006-01-02", end)
	if err != nil {
		return nil, err
	}
	return s.mediaRepository.GetUserMovieCommentsByRange(userID, startTime, endTime)
}

func (s *CommentService) GetUserTvShowCommentsByRange(userID string, start, end string) ([]*repository2.TvShowComment, error) {
	startTime, err := time.Parse("2006-01-02", start)
	if err != nil {
		return nil, err
	}
	endTime, err := time.Parse("2006-01-02", end)
	if err != nil {
		return nil, err
	}
	return s.mediaRepository.GetUserTvShowCommentsByRange(userID, startTime, endTime)
}

func (s *CommentService) CountUserComments(userID string) (int, error) {
	movieCommentsCount, err := s.mediaRepository.CountUserMovieComments(userID)
	if err != nil {
		return 0, err
	}
	tvShowCommentsCount, err := s.mediaRepository.CountUserTvShowComments(userID)
	if err != nil {
		return 0, err
	}
	return movieCommentsCount + tvShowCommentsCount, nil
}

func (s *CommentService) GetMovieCommentsByRange(start, end string) ([]*repository2.MovieComment, error) {
	startTime, err := time.Parse("2006-01-02", start)
	if err != nil {
		return nil, err
	}
	endTime, err := time.Parse("2006-01-02", end)
	if err != nil {
		return nil, err
	}
	return s.mediaRepository.GetMovieCommentsByRange(startTime, endTime)
}

func (s *CommentService) GetTvShowCommentsByRange(start, end string) ([]*repository2.TvShowComment, error) {
	startTime, err := time.Parse("2006-01-02", start)
	if err != nil {
		return nil, err
	}
	endTime, err := time.Parse("2006-01-02", end)
	if err != nil {
		return nil, err
	}
	return s.mediaRepository.GetTvShowCommentsByRange(startTime, endTime)
}

func (s *CommentService) CountComments() (int, error) {
	movieCommentsCount, err := s.mediaRepository.CountMovieComments()
	if err != nil {
		return 0, err
	}
	tvShowCommentsCount, err := s.mediaRepository.CountTvShowComments()
	if err != nil {
		return 0, err
	}
	return movieCommentsCount + tvShowCommentsCount, nil
}
