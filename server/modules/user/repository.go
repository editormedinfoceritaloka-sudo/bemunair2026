package user

import (
	"errors"

	"bemunair2026/server/database/entities"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository { return &Repository{db: db} }

func (r *Repository) Create(u *entities.User) error {
	return r.db.Create(u).Error
}

func (r *Repository) FindByID(id uint64) (*entities.User, error) {
	var u entities.User
	err := r.db.First(&u, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &u, err
}

func (r *Repository) FindByEmail(email string) (*entities.User, error) {
	var u entities.User
	err := r.db.Where("email = ?", email).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &u, err
}

func (r *Repository) List() ([]entities.User, error) {
	var users []entities.User
	return users, r.db.Order("id ASC").Find(&users).Error
}

func (r *Repository) Update(u *entities.User) error {
	return r.db.Save(u).Error
}

func (r *Repository) Delete(id uint64) error {
	return r.db.Delete(&entities.User{}, id).Error
}
