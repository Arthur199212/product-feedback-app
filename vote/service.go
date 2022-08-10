package vote

import (
	"database/sql"
	"errors"
	"product-feedback/notifier"
)

type VoteService interface {
	Create(userId int, v createVoteInput) (int, error)
	Delete(userId, voteId int) error
	GetAll(feedbackId *int) ([]Vote, error)
}

type voteService struct {
	repo     VoteRepository
	notifier notifier.NotifierService
}

func NewVoteService(
	repo VoteRepository,
	notifier notifier.NotifierService,
) VoteService {
	return &voteService{
		repo:     repo,
		notifier: notifier,
	}
}

var ErrVoteAlreadyExists = errors.New("vote already exists")

func (s *voteService) Create(userId int, v createVoteInput) (int, error) {
	_, err := s.repo.GetByFeedbackId(userId, v.FeedbackId)
	// if vote doesn't exists -> then create a vote
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	// if vote exists -> do nothing
	if err == nil {
		return 0, ErrVoteAlreadyExists
	}

	id, err := s.repo.Create(userId, v)
	if err != nil {
		return id, err
	}

	go s.notifier.BroadcastMessage(
		notifier.CreateEvent,
		notifier.SubjectVote,
		id,
	)

	return id, nil
}

func (s *voteService) Delete(userId, voteId int) error {
	if err := s.repo.Delete(userId, voteId); err != nil {
		return err
	}

	go s.notifier.BroadcastMessage(
		notifier.DeleteEvent,
		notifier.SubjectVote,
		voteId,
	)

	return nil
}

func (s *voteService) GetAll(feedbackId *int) ([]Vote, error) {
	return s.repo.GetAll(feedbackId)
}
