package user

// UserService is the interface for user-related business logic
type UserService interface {
	GetAllUsers() ([]User, error)
	GetUserByID(id uint) (*User, error)
	CreateUser(user *User) error
	UpdateUser(id uint, user *User) error
	DeleteUser(id uint) error
}

// Service is the concrete implementation of the UserService interface
type Service struct {
	repo UserRepository
}

func NewService(repo UserRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllUsers() ([]User, error) {
	return s.repo.FindAll()
}

func (s *Service) GetUserByID(id uint) (*User, error) {
	return s.repo.FindByID(id)
}

func (s *Service) CreateUser(user *User) error {
	return s.repo.Create(user)
}

func (s *Service) UpdateUser(id uint, user *User) error {
	return s.repo.Update(id, user)
}

func (s *Service) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
