package vote

type VoteService interface {
	Create(userId int, v createVoteInput) (int, error)
	GetAll(feedbackId *int) ([]Vote, error)
}

type voteService struct {
	repo VoteRepository
}

func NewVoteService(repo VoteRepository) VoteService {
	return &voteService{repo}
}

func (s *voteService) Create(userId int, v createVoteInput) (int, error) {
	return s.repo.Create(userId, v)
}

func (s *voteService) GetAll(feedbackId *int) ([]Vote, error) {
	return s.repo.GetAll(feedbackId)
}
