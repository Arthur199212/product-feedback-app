package vote

type VoteService interface {
}

type voteService struct {
	repo VoteRepository
}

func NewVoteService(repo VoteRepository) VoteService {
	return &voteService{repo}
}
