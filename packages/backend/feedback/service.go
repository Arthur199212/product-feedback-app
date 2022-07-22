package feedback

type FeedbackService interface {
	Create(userId int, f createFeedbackInput) (int, error)
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

func (s *feedbackService) GetById(userId, feedbackId int) (Feedback, error) {
	return s.repo.GetById(userId, feedbackId)
}
