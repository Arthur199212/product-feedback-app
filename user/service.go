package user

type UserService interface {
	Create(user User) (int, error)
	GetByEmail(email string) (User, error)
	GetById(id int) (User, error)
	GetAll(userIds []int) ([]User, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Create(user User) (int, error) {
	return s.repo.Create(user)
}

func (s *userService) GetByEmail(email string) (User, error) {
	return s.repo.GetByEmail(email)
}

func (s *userService) GetById(id int) (User, error) {
	return s.repo.GetById(id)
}

func (s *userService) GetAll(userIds []int) ([]User, error) {
	return s.repo.GetAll(userIds)
}
