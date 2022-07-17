package vote

type VoteService struct {
	repo *VoteRepository
}

func NewVoteService(repo *VoteRepository) *VoteService {
	return &VoteService{repo}
}
