package product

import "gorm.io/gorm"

// ProductRepository is the interface for the product repository
type ProductRepository interface {
	FindAll() ([]Product, error)
	FindByID(id uint) (*Product, error)
	Create(product *Product) error
	Update(id uint, product *Product) error
	Delete(id uint) error
}

// Repository is the concrete GORM implementation of ProductRepository
type Repository struct {
	DB *gorm.DB
}

func (r *Repository) FindAll() ([]Product, error) {
	var products []Product
	if err := r.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *Repository) FindByID(id uint) (*Product, error) {
	var product Product
	if err := r.DB.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *Repository) Create(product *Product) error {
	return r.DB.Create(product).Error
}

func (r *Repository) Update(id uint, product *Product) error {
	var existingProduct Product
	if err := r.DB.First(&existingProduct, id).Error; err != nil {
		return err
	}
	return r.DB.Model(&existingProduct).Updates(product).Error
}

func (r *Repository) Delete(id uint) error {
	return r.DB.Delete(&Product{}, id).Error
}
