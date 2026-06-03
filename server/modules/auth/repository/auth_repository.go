package repository

import (
	"bemunair2026/server/database/entities"
	"bemunair2026/server/modules/user"
)

type Repository struct {
	users *user.Repository
}

func NewRepository(users *user.Repository) *Repository {
	return &Repository{users: users}
}

func (r *Repository) Create(user *entities.User) error {
	return r.users.Create(user)
}

func (r *Repository) FindByID(id uint64) (*entities.User, error) {
	return r.users.FindByID(id)
}

func (r *Repository) FindByEmail(email string) (*entities.User, error) {
	return r.users.FindByEmail(email)
}
