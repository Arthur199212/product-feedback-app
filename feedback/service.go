package feedback

type FeedbackService interface {
	Create(userId int, f createFeedbackInput) (int, error)
	Delete(userId, feedbackId int) error
	GetAll() ([]Feedback, error)
	GetById(userId, feedbackId int) (Feedback, error)
	Update(userId, feedbackId int, f updateFeedbackInput) error
}

type feedbackService struct {
	repo FeedbackRepository
}

func NewFeedbackService(repo FeedbackRepository) FeedbackService {
	return &feedbackService{repo}
}

func (s *feedbackService) Create(userId int, f createFeedbackInput) (int, error) {
	return s.repo.Create(userId, f)
}

func (s *feedbackService) Delete(userId, feedbackId int) error {
	// check if feedback exists
	_, err := s.GetById(userId, feedbackId)
	if err != nil {
		return err
	}

	return s.repo.Delete(userId, feedbackId)
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
	f updateFeedbackInput,
) error {
	return s.repo.Update(userId, feedbackId, f)
}
