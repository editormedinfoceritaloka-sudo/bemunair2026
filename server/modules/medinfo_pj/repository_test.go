package medinfo_pj

import (
	"testing"

	"bemunair2026/server/database/entities"
	medinfoRepository "bemunair2026/server/modules/medinfo_pj/repository"
	"bemunair2026/server/pkg/constants"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestAssignNextRoundRobin(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:round_robin?mode=memory&cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := createTestTables(db); err != nil {
		t.Fatal(err)
	}
	users := []entities.User{{Name: "A", Email: "a@test", PasswordHash: "x", Role: constants.RoleMentri}, {Name: "B", Email: "b@test", PasswordHash: "x", Role: constants.RoleMentri}}
	if err := db.Create(&users).Error; err != nil {
		t.Fatal(err)
	}
	rows := []entities.MedinfoPJQueue{{UserID: users[0].ID, Position: 1, IsCurrent: true}, {UserID: users[1].ID, Position: 2}}
	if err := db.Create(&rows).Error; err != nil {
		t.Fatal(err)
	}

	var first, second *entities.User
	if err := db.Transaction(func(tx *gorm.DB) error { var err error; first, err = medinfoRepository.AssignNext(tx); return err }); err != nil {
		t.Fatal(err)
	}
	if err := db.Transaction(func(tx *gorm.DB) error { var err error; second, err = medinfoRepository.AssignNext(tx); return err }); err != nil {
		t.Fatal(err)
	}
	if first.ID != users[0].ID || second.ID != users[1].ID {
		t.Fatalf("rotation got %d then %d", first.ID, second.ID)
	}
}

func TestAssignNextEmptyQueue(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:empty_queue?mode=memory&cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := createTestTables(db); err != nil {
		t.Fatal(err)
	}
	var pj *entities.User
	if err := db.Transaction(func(tx *gorm.DB) error { var err error; pj, err = medinfoRepository.AssignNext(tx); return err }); err != nil {
		t.Fatal(err)
	}
	if pj != nil {
		t.Fatalf("expected nil PJ")
	}
}

func createTestTables(db *gorm.DB) error {
	if err := db.Exec(`CREATE TABLE users (id integer primary key autoincrement, name text, email text unique, password_hash text, role text, ministry text, phone text, created_at datetime, updated_at datetime)`).Error; err != nil {
		return err
	}
	return db.Exec(`CREATE TABLE medinfo_pj_queues (id integer primary key autoincrement, user_id integer unique, position integer, is_current boolean, created_at datetime, updated_at datetime)`).Error
}
