package comment

import (
	"errors"
	"product-feedback/notifier"
)

type CommentService interface {
	Create(userId int, f createCommentInput) (int, error)
	GetAll(feedbackId *int) ([]Comment, error)
	GetById(commentId int) (Comment, error)
}

type commentService struct {
	repo     CommentRepository
	notifier notifier.NotifierService
}

func NewCommentService(
	repo CommentRepository,
	notifier notifier.NotifierService,
) CommentService {
	return &commentService{
		repo:     repo,
		notifier: notifier,
	}
}

func (s *commentService) Create(userId int, f createCommentInput) (int, error) {
	// check if parent comment has the same feedbackId
	if f.ParentId != nil {
		comment, err := s.repo.GetById(*f.ParentId)
		if err != nil {
			return 0, err
		}
		if comment.FeedbackId != f.FeedbackId {
			return 0, errors.New("feedbackId of a parent comment and a new comment are not the same")
		}
	}

	id, err := s.repo.Create(userId, f)
	if err != nil {
		return id, err
	}

	go s.notifier.BroadcastMessage(
		notifier.CreateEvent,
		notifier.SubjectComment,
		id,
	)

	return id, nil
}

func (s *commentService) GetAll(feedbackId *int) ([]Comment, error) {
	return s.repo.GetAll(feedbackId)
}

func (s *commentService) GetById(commentId int) (Comment, error) {
	return s.repo.GetById(commentId)
}
