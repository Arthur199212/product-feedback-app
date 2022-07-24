package vote

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

func (s *voteService) Create(userId int, v createVoteInput) (int, error) {
	return s.repo.Create(userId, v)
}

func (s *voteService) Delete(userId, voteId int) error {
	return s.repo.Delete(userId, voteId)
}

func (s *voteService) GetAll(feedbackId *int) ([]Vote, error) {
	return s.repo.GetAll(feedbackId)
}
