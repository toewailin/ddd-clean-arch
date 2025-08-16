package faq

// FAQService is the interface for our FAQ business logic
type FAQService interface {
	GetAllFAQs() ([]FAQ, error)
	GetFAQByID(id uint) (*FAQ, error)
	CreateFAQ(faq *FAQ) error
	UpdateFAQ(id uint, faq *FAQ) error
	DeleteFAQ(id uint) error
}

// Service is the concrete implementation of FAQService
type Service struct {
	repo FAQRepository
}

func NewService(repo FAQRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllFAQs() ([]FAQ, error) {
	return s.repo.FindAll()
}

func (s *Service) GetFAQByID(id uint) (*FAQ, error) {
	return s.repo.FindByID(id)
}

func (s *Service) CreateFAQ(faq *FAQ) error {
	return s.repo.Create(faq)
}

func (s *Service) UpdateFAQ(id uint, faq *FAQ) error {
	return s.repo.Update(id, faq)
}

func (s *Service) DeleteFAQ(id uint) error {
	return s.repo.Delete(id)
}