package comment

type CommentService interface {
	Create(userId int, f createCommentInput) (int, error)
	GetAll(feedbackId *int) ([]Comment, error)
	GetById(commentId int) (Comment, error)
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

func (s *commentService) GetAll(feedbackId *int) ([]Comment, error) {
	return s.repo.GetAll(feedbackId)
}

func (s *commentService) GetById(commentId int) (Comment, error) {
	return s.repo.GetById(commentId)
}