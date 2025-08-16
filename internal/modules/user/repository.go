package user

import "gorm.io/gorm"

// UserRepository is the interface for user repository methods
type UserRepository interface {
	FindAll() ([]User, error)
	FindByID(id uint) (*User, error)
	Create(user *User) error
	Update(id uint, user *User) error
	Delete(id uint) error
}

// Repository is the concrete implementation of the UserRepository interface
type Repository struct {
	DB *gorm.DB
}

func (r *Repository) FindAll() ([]User, error) {
	var users []User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *Repository) FindByID(id uint) (*User, error) {
	var user User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) Create(user *User) error {
	return r.DB.Create(user).Error
}

func (r *Repository) Update(id uint, user *User) error {
	var existingUser User
	if err := r.DB.First(&existingUser, id).Error; err != nil {
		return err
	}
	return r.DB.Model(&existingUser).Updates(user).Error
}

func (r *Repository) Delete(id uint) error {
	return r.DB.Delete(&User{}, id).Error
}
