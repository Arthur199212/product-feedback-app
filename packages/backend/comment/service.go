package comment

type CommentService struct {
	repo *CommentRepository
}

func NewCommentService(repo *CommentRepository) *CommentService {
	return &CommentService{repo}
}
