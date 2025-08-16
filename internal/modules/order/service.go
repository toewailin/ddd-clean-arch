package order

// OrderService is the interface for order business logic
type OrderService interface {
	GetAllOrders() ([]Order, error)
	GetOrderByID(id uint) (*Order, error)
	CreateOrder(order *Order) error
	UpdateOrder(id uint, order *Order) error
	DeleteOrder(id uint) error
}

// Service is the concrete implementation of OrderService
type Service struct {
	repo OrderRepository
}

func NewService(repo OrderRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllOrders() ([]Order, error) {
	return s.repo.FindAll()
}

func (s *Service) GetOrderByID(id uint) (*Order, error) {
	return s.repo.FindByID(id)
}

func (s *Service) CreateOrder(order *Order) error {
	return s.repo.Create(order)
}

func (s *Service) UpdateOrder(id uint, order *Order) error {
	return s.repo.Update(id, order)
}

func (s *Service) DeleteOrder(id uint) error {
	return s.repo.Delete(id)
}
