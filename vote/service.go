package vote

import (
	"database/sql"
	"errors"
	"product-feedback/notifier"
)

type VoteService interface {
	Create(userId, feedbackId int) (int, error)
	Delete(userId, voteId int) error
	GetAll(feedbackIds []int) ([]Vote, error)
	Toggle(userId int, feedbackId int) error
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

func (s *voteService) Create(userId int, feedbackId int) (int, error) {
	_, err := s.repo.GetByFeedbackId(userId, feedbackId)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	// if vote exists -> do nothing
	if err == nil {
		return 0, ErrVoteAlreadyExists
	}

	// if vote doesn't exists -> then create a vote
	id, err := s.repo.Create(userId, feedbackId)
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

func (s *voteService) GetAll(feedbackIds []int) ([]Vote, error) {
	return s.repo.GetAll(feedbackIds)
}

func (s *voteService) Toggle(userId int, feedbackId int) error {
	vote, err := s.repo.GetByFeedbackId(userId, feedbackId)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err == nil {
		return s.Delete(userId, vote.Id)
	}

	_, err = s.Create(userId, feedbackId)
	return err
}
