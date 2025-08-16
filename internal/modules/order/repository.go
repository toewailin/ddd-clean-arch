package order

import "gorm.io/gorm"

// OrderRepository is the interface for the order repository
type OrderRepository interface {
	FindAll() ([]Order, error)
	FindByID(id uint) (*Order, error)
	Create(order *Order) error
	Update(id uint, order *Order) error
	Delete(id uint) error
}

// Repository is the concrete GORM implementation of OrderRepository
type Repository struct {
	DB *gorm.DB
}

func (r *Repository) FindAll() ([]Order, error) {
	var orders []Order
	if err := r.DB.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *Repository) FindByID(id uint) (*Order, error) {
	var order Order
	if err := r.DB.First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *Repository) Create(order *Order) error {
	return r.DB.Create(order).Error
}

func (r *Repository) Update(id uint, order *Order) error {
	var existingOrder Order
	if err := r.DB.First(&existingOrder, id).Error; err != nil {
		return err
	}
	return r.DB.Model(&existingOrder).Updates(order).Error
}

func (r *Repository) Delete(id uint) error {
	return r.DB.Delete(&Order{}, id).Error
}
