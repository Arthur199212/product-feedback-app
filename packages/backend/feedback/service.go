package feedback

type FeedbackService interface {
	Create(userId int, f createFeedbackInput) (int, error)
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
