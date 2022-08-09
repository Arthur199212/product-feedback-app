package feedback

import (
	"product-feedback/notifier"
)

//go:generate mockgen -source=service.go -destination=mocks/service.go

type FeedbackService interface {
	Create(userId int, f CreateFeedbackInput) (int, error)
	Delete(userId, feedbackId int) error
	GetAll() ([]Feedback, error)
	GetById(userId, feedbackId int) (Feedback, error)
	Update(userId, feedbackId int, f UpdateFeedbackInput) error
}

type feedbackService struct {
	repo     FeedbackRepository
	notifier *notifier.NotifierService
}

func NewFeedbackService(
	repo FeedbackRepository,
	hub *notifier.NotifierService,
) FeedbackService {
	return &feedbackService{repo: repo, notifier: hub}
}

func (s *feedbackService) Create(userId int, f CreateFeedbackInput) (int, error) {
	id, err := s.repo.Create(userId, f)
	if err != nil {
		return id, err
	}

	// notify about feedback create
	go s.notifier.BroadcastMessage(
		notifier.DeleteEvent,
		notifier.SubjectFeedback,
		id,
	)

	return id, nil
}

func (s *feedbackService) Delete(userId, feedbackId int) error {
	// check if feedback exists
	_, err := s.GetById(userId, feedbackId)
	if err != nil {
		return err
	}

	if err := s.repo.Delete(userId, feedbackId); err != nil {
		return err
	}

	// notify about feedback delete
	go s.notifier.BroadcastMessage(
		notifier.DeleteEvent,
		notifier.SubjectFeedback,
		feedbackId,
	)

	return nil
}

func (s *feedbackService) GetAll() ([]Feedback, error) {
	return s.repo.GetAll()
}

func (s *feedbackService) GetById(userId, feedbackId int) (Feedback, error) {
	return s.repo.GetById(userId, feedbackId)
}

func (s *feedbackService) Update(
	userId,
	feedbackId int,
	f UpdateFeedbackInput,
) error {
	if err := s.repo.Update(userId, feedbackId, f); err != nil {
		return err
	}

	// notify about feedback update
	go s.notifier.BroadcastMessage(
		notifier.UpdateEvent,
		notifier.SubjectFeedback,
		feedbackId,
	)

	return nil
}
