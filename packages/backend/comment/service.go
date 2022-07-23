package comment

type CommentService interface {
	Create(userId int, f createCommentInput) (int, error)
}

type commentService struct {
	repo CommentRepository
}

func NewCommentService(repo CommentRepository) CommentService {
	return &commentService{repo}
}

func (s *commentService) Create(userId int, f createCommentInput) (int, error) {
	return s.repo.Create(userId, f)
}
