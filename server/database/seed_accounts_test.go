package database

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestProvisionAccounts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(filepath.Join(t.TempDir(), "accounts.db")), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		t.Fatal(err)
	}
	for _, statement := range []string{
		"CREATE TABLE ministries (id INTEGER PRIMARY KEY, code TEXT UNIQUE)",
		"CREATE TABLE users (id INTEGER PRIMARY KEY, email TEXT UNIQUE, name TEXT, phone TEXT, ministry TEXT, ministry_id INTEGER, role TEXT, password_hash TEXT, created_at datetime, updated_at datetime)",
	} {
		if err := db.Exec(statement).Error; err != nil {
			t.Fatal(err)
		}
	}
	for _, account := range seedAccounts {
		if account.unit != "" {
			if err := db.Exec("INSERT OR IGNORE INTO ministries(code) VALUES (?)", account.unit).Error; err != nil {
				t.Fatal(err)
			}
		}
	}

	legacy, err := bcrypt.GenerateFromPassword([]byte("legacy"), bcrypt.MinCost)
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Exec("INSERT INTO users(email, password_hash, role, ministry_id) VALUES (?, ?, 'ADMIN', (SELECT id FROM ministries WHERE code = 'ADKESMA'))", "menteri.adkesma@bem.unair.ac.id", string(legacy)).Error; err != nil {
		t.Fatal(err)
	}

	dir := filepath.Join(t.TempDir(), "credentials")
	manager := &SeederManager{db: db, credentialDir: dir, legacyPassword: "legacy"}
	if err := manager.provisionAccounts(false); err != nil {
		t.Fatal(err)
	}

	passwords := map[string]string{}
	for _, account := range seedAccounts {
		file := filepath.Join(dir, account.email)
		info, err := os.Stat(file)
		if err != nil {
			t.Fatalf("missing credential for %s: %v", account.email, err)
		}
		if info.Mode().Perm() != 0600 {
			t.Fatalf("unsafe credential mode for %s: %v", account.email, info.Mode().Perm())
		}
		password, err := os.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}
		if len(password) < 20 {
			t.Fatalf("short password for %s", account.email)
		}
		if previous, ok := passwords[string(password)]; ok {
			t.Fatalf("duplicate password for %s and %s", account.email, previous)
		}
		passwords[string(password)] = account.email

		var hash string
		if err := db.Raw("SELECT password_hash FROM users WHERE email = ?", account.email).Scan(&hash).Error; err != nil {
			t.Fatal(err)
		}
		if bcrypt.CompareHashAndPassword([]byte(hash), bytes.TrimSpace(password)) != nil {
			t.Fatalf("hash does not match credential for %s", account.email)
		}
	}

	before := map[string]string{}
	for _, account := range seedAccounts {
		var hash string
		if err := db.Raw("SELECT password_hash FROM users WHERE email = ?", account.email).Scan(&hash).Error; err != nil {
			t.Fatal(err)
		}
		before[account.email] = hash
	}
	if err := manager.provisionAccounts(false); err != nil {
		t.Fatal(err)
	}
	for email, hash := range before {
		var current string
		if err := db.Raw("SELECT password_hash FROM users WHERE email = ?", email).Scan(&current).Error; err != nil {
			t.Fatal(err)
		}
		if current != hash {
			t.Fatalf("rerun rotated %s", email)
		}
	}

	// A customized password survives both the full dev seed and the manual command.
	custom, err := bcrypt.GenerateFromPassword([]byte("custom"), bcrypt.MinCost)
	if err != nil {
		t.Fatal(err)
	}
	email := "menteri.adkesma@bem.unair.ac.id"
	if err := db.Exec("UPDATE users SET password_hash = ? WHERE email = ?", string(custom), email).Error; err != nil {
		t.Fatal(err)
	}
	if err := db.Exec("DELETE FROM users WHERE email IN (?, ?)", "menteri.psdm@bem.unair.ac.id", "ocha.medinfo@bem.unair.ac.id").Error; err != nil {
		t.Fatal(err)
	}
	if err := manager.provisionAccounts(true); err != nil {
		t.Fatal(err)
	}
	var hash string
	if err := db.Raw("SELECT password_hash FROM users WHERE email = ?", email).Scan(&hash).Error; err != nil || hash != string(custom) {
		t.Fatal("changed custom password", err)
	}
	var count int
	if err := db.Raw("SELECT COUNT(*) FROM users WHERE email = ?", "ocha.medinfo@bem.unair.ac.id").Scan(&count).Error; err != nil || count != 0 {
		t.Fatal("provisioned non-leader", err)
	}
	if err := db.Raw("SELECT COUNT(*) FROM users WHERE email = ?", "menteri.psdm@bem.unair.ac.id").Scan(&count).Error; err != nil || count != 1 {
		t.Fatal("missing leader", err)
	}
	if err := db.Exec("UPDATE users SET password_hash = ? WHERE email = ?", string(legacy), email).Error; err != nil {
		t.Fatal(err)
	}
	// A credential write/read failure must leave the current hash usable.
	if err := os.Chmod(filepath.Join(dir, email), 0644); err != nil {
		t.Fatal(err)
	}
	if err := manager.provisionAccounts(true); err == nil {
		t.Fatal("accepted insecure credential file")
	}
	if err := db.Raw("SELECT password_hash FROM users WHERE email = ?", email).Scan(&hash).Error; err != nil || hash != string(legacy) {
		t.Fatal("updated password before durable credential", err)
	}
}
