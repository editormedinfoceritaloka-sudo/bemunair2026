package repository

import (
	"errors"

	"bemunair2026/server/database/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entities.User) error
	FindByID(id uint64) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	List() ([]entities.User, error)
	Update(user *entities.User) error
	Delete(id uint64) error
}

type userRepository struct {
	db *gorm.DB
}

var _ UserRepository = (*userRepository)(nil)

func NewUserRepository(db *gorm.DB) UserRepository { return &userRepository{db: db} }

func (r *userRepository) Create(u *entities.User) error {
	return r.db.Create(u).Error
}

func (r *userRepository) FindByID(id uint64) (*entities.User, error) {
	var u entities.User
	err := r.db.First(&u, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &u, err
}

func (r *userRepository) FindByEmail(email string) (*entities.User, error) {
	var u entities.User
	err := r.db.Where("email = ?", email).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &u, err
}

func (r *userRepository) List() ([]entities.User, error) {
	var users []entities.User
	return users, r.db.Order("id ASC").Find(&users).Error
}

func (r *userRepository) Update(u *entities.User) error {
	return r.db.Save(u).Error
}

func (r *userRepository) Delete(id uint64) error {
	return r.db.Delete(&entities.User{}, id).Error
}
