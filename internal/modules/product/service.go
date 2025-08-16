package product

// ProductService is the interface for product business logic
type ProductService interface {
	GetAllProducts() ([]Product, error)
	GetProductByID(id uint) (*Product, error)
	CreateProduct(product *Product) error
	UpdateProduct(id uint, product *Product) error
	DeleteProduct(id uint) error
}

// Service is the concrete implementation of ProductService
type Service struct {
	repo ProductRepository
}

func NewService(repo ProductRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllProducts() ([]Product, error) {
	return s.repo.FindAll()
}

func (s *Service) GetProductByID(id uint) (*Product, error) {
	return s.repo.FindByID(id)
}

func (s *Service) CreateProduct(product *Product) error {
	return s.repo.Create(product)
}

func (s *Service) UpdateProduct(id uint, product *Product) error {
	return s.repo.Update(id, product)
}

func (s *Service) DeleteProduct(id uint) error {
	return s.repo.Delete(id)
}
