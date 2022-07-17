package feedback

type FeedbackService struct {
	repo *FeedbackRepository
}

func NewFeedbackService(repo *FeedbackRepository) *FeedbackService {
	return &FeedbackService{repo}
}
