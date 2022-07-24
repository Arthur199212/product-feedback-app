package vote

type VoteService interface {
	Create(userId int, v createVoteInput) (int, error)
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
