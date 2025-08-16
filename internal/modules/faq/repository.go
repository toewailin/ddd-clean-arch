package faq

import (
	"gorm.io/gorm"
)

// FAQRepository is the interface for our database abstraction
type FAQRepository interface {
	FindAll() ([]FAQ, error)
	FindByID(id uint) (*FAQ, error)
	Create(faq *FAQ) error
	Update(id uint, faq *FAQ) error
	Delete(id uint) error
}

// Repository is the concrete GORM implementation of FAQRepository
type Repository struct {
	DB *gorm.DB
}

func (r *Repository) FindAll() ([]FAQ, error) {
	var faqs []FAQ
	if err := r.DB.Find(&faqs).Error; err != nil {
		return nil, err
	}
	return faqs, nil
}

func (r *Repository) FindByID(id uint) (*FAQ, error) {
	var faq FAQ
	if err := r.DB.First(&faq, id).Error; err != nil {
		return nil, err
	}
	return &faq, nil
}

func (r *Repository) Create(faq *FAQ) error {
	return r.DB.Create(faq).Error
}

func (r *Repository) Update(id uint, faq *FAQ) error {
	var existingFAQ FAQ
	if err := r.DB.First(&existingFAQ, id).Error; err != nil {
		return err
	}
	return r.DB.Model(&existingFAQ).Updates(faq).Error
}

func (r *Repository) Delete(id uint) error {
	return r.DB.Delete(&FAQ{}, id).Error
}