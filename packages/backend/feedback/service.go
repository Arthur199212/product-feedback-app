package feedback

type FeedbackService interface {
	Create(userId int, f createFeedbackInput) (int, error)
	Delete(userId, feedbackId int) error
	GetAll() ([]Feedback, error)
	GetById(userId, feedbackId int) (Feedback, error)
}

type feedbackService struct {
	repo FeedbackRepository
}

func NewFeedbackService(repo FeedbackRepository) FeedbackService {
	return &feedbackService{repo}
}

func (s *feedbackService) Create(userId int, f createFeedbackInput) (int, error) {
	if f.Status == "" {
		f.Status = "idea"
	}
	return s.repo.Create(userId, f)
}

func (s *feedbackService) Delete(userId, feedbackId int) error {
	return s.repo.Delete(userId, feedbackId)
}

func (s *feedbackService) GetAll() ([]Feedback, error) {
	return s.repo.GetAll()
}

func (s *feedbackService) GetById(userId, feedbackId int) (Feedback, error) {
	return s.repo.GetById(userId, feedbackId)
}
