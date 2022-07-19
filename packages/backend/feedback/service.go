package feedback

type FeedbackService interface {
}

type feedbackService struct {
	repo FeedbackRepository
}

func NewFeedbackService(repo FeedbackRepository) FeedbackService {
	return &feedbackService{repo}
}
