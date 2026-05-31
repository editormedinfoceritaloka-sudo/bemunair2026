package auth

import (
	"testing"

	"bemunair2026/server/database/entities"
	"bemunair2026/server/modules/user"
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
	service := NewService(user.NewRepository(db), "secret")
	if _, err := service.Register("Admin", "admin@test", "secret", entities.RoleAdmin, nil, nil); err != nil {
		t.Fatal(err)
	}
	token, u, err := service.Login("admin@test", "secret")
	if err != nil || token == "" || u.Email != "admin@test" {
		t.Fatalf("login failed token=%q user=%+v err=%v", token, u, err)
	}
	if _, _, err := service.Login("admin@test", "bad"); err == nil {
		t.Fatal("expected bad password error")
	}
}
