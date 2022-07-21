package user

type UserService interface {
	Create(user User) (int, error)
	GetByEmail(email string) (User, error)
	GetById(id int) (User, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Create(user User) (int, error) {
	// todo: validation in handler
	return s.repo.Create(user)
}

func (s *userService) GetByEmail(email string) (User, error) {
	var user User
	// todo
	return user, nil
}

func (s *userService) GetById(id int) (User, error) {
	var user User
	// todo
	return user, nil
}
