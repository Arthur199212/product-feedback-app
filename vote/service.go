package vote

import (
	"database/sql"
	"errors"
)

type VoteService interface {
	Create(userId int, v createVoteInput) (int, error)
	Delete(userId, voteId int) error
	GetAll(feedbackId *int) ([]Vote, error)
}

type voteService struct {
	repo VoteRepository
}

func NewVoteService(repo VoteRepository) VoteService {
	return &voteService{repo}
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

	return s.repo.Create(userId, v)
}

func (s *voteService) Delete(userId, voteId int) error {
	return s.repo.Delete(userId, voteId)
}

func (s *voteService) GetAll(feedbackId *int) ([]Vote, error) {
	return s.repo.GetAll(feedbackId)
}
