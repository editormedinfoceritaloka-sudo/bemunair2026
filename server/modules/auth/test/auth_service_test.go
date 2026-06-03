package test

import (
	"testing"

	"bemunair2026/server/database/entities"
	"bemunair2026/server/modules/auth/dto"
	authService "bemunair2026/server/modules/auth/service"
	"bemunair2026/server/pkg/constants"
)

func TestRegisterAndLogin(t *testing.T) {
	userRepo := newFakeUserRepository()
	service := authService.NewAuthService(userRepo, "secret")
	if _, err := service.Register(dto.RegisterRequest{Name: "Admin", Email: "admin@test", Password: "password", Role: constants.RoleAdmin}); err != nil {
		t.Fatal(err)
	}
	res, err := service.Login(dto.LoginRequest{Email: "admin@test", Password: "password"})
	if err != nil || res.Token == "" || res.User.Email != "admin@test" {
		t.Fatalf("login failed response=%+v err=%v", res, err)
	}
	if _, err := service.Login(dto.LoginRequest{Email: "admin@test", Password: "bad"}); err == nil {
		t.Fatal("expected bad password error")
	}
}

type fakeUserRepository struct {
	nextID uint64
	users  map[uint64]*entities.User
	emails map[string]*entities.User
}

func newFakeUserRepository() *fakeUserRepository {
	return &fakeUserRepository{
		nextID: 1,
		users:  map[uint64]*entities.User{},
		emails: map[string]*entities.User{},
	}
}

func (r *fakeUserRepository) Create(user *entities.User) error {
	user.ID = r.nextID
	r.nextID++
	r.users[user.ID] = user
	r.emails[user.Email] = user
	return nil
}

func (r *fakeUserRepository) FindByID(id uint64) (*entities.User, error) {
	return r.users[id], nil
}

func (r *fakeUserRepository) FindByEmail(email string) (*entities.User, error) {
	return r.emails[email], nil
}

func (r *fakeUserRepository) List() ([]entities.User, error) {
	users := make([]entities.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, *user)
	}
	return users, nil
}

func (r *fakeUserRepository) Update(user *entities.User) error {
	r.users[user.ID] = user
	r.emails[user.Email] = user
	return nil
}

func (r *fakeUserRepository) Delete(id uint64) error {
	delete(r.users, id)
	return nil
}
