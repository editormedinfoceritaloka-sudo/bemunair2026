package test

import (
	"testing"

	"bemunair2026/server/modules/auth/dto"
	"bemunair2026/server/modules/auth/repository"
	"bemunair2026/server/modules/auth/service"
	"bemunair2026/server/modules/user"
	"bemunair2026/server/pkg/constants"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestRegisterAndLogin(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Exec(`CREATE TABLE users (id integer primary key autoincrement, name text, email text unique, password_hash text, role text, ministry text, phone text, created_at datetime, updated_at datetime)`).Error; err != nil {
		t.Fatal(err)
	}
	userRepo := user.NewRepository(db)
	authRepo := repository.NewRepository(userRepo)
	authService := service.NewService(authRepo, "secret")
	if _, err := authService.Register(dto.RegisterRequest{Name: "Admin", Email: "admin@test", Password: "secret", Role: constants.RoleAdmin}); err != nil {
		t.Fatal(err)
	}
	res, err := authService.Login(dto.LoginRequest{Email: "admin@test", Password: "secret"})
	if err != nil || res.Token == "" || res.User.Email != "admin@test" {
		t.Fatalf("login failed response=%+v err=%v", res, err)
	}
	if _, err := authService.Login(dto.LoginRequest{Email: "admin@test", Password: "bad"}); err == nil {
		t.Fatal("expected bad password error")
	}
}
