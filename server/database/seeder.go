package database

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"

	"bemunair2026/server/config"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//go:embed seeders/*.sql
var seedersFS embed.FS

// seedPasswordPlaceholder appears in seeders/001_users.sql instead of a
// literal bcrypt hash. No real password — plaintext or hashed — is ever
// committed for the seeded accounts; it is computed here from
// SEED_USER_PASSWORD (gitignored .env) at seed time instead.
const seedPasswordPlaceholder = "__SEED_PASSWORD_HASH__"

type SeederManager struct {
	db               *gorm.DB
	seedUserPassword string
}

func NewSeederManager(db *gorm.DB, seedUserPassword string) *SeederManager {
	return &SeederManager{db: db, seedUserPassword: seedUserPassword}
}

// Run executes every seeder in lexical order. Seeder SQL must be idempotent so
// this can safely run whenever the development stack starts.
func (s *SeederManager) Run() error {
	names, err := seederNames()
	if err != nil {
		return err
	}

	for _, name := range names {
		if err := s.runSeeder(name); err != nil {
			return err
		}
	}

	return nil
}

func (s *SeederManager) runSeeder(name string) error {
	content, err := seedersFS.ReadFile(filepath.ToSlash(filepath.Join("seeders", name)))
	if err != nil {
		return err
	}

	rendered, err := s.renderSeedPassword(string(content))
	if err != nil {
		return fmt.Errorf("run seeder %s: %w", name, err)
	}

	return s.db.Transaction(func(tx *gorm.DB) error {
		for _, statement := range splitSQLStatements(rendered) {
			if err := tx.Exec(statement).Error; err != nil {
				return fmt.Errorf("run seeder %s: %w", name, err)
			}
		}
		return nil
	})
}

// renderSeedPassword substitutes seedPasswordPlaceholder with a freshly
// computed bcrypt hash of SEED_USER_PASSWORD. A seeder with no placeholder is
// returned unchanged, so this never requires the env var unless a seeder
// actually needs it.
func (s *SeederManager) renderSeedPassword(content string) (string, error) {
	if !strings.Contains(content, seedPasswordPlaceholder) {
		return content, nil
	}
	if s.seedUserPassword == "" {
		return "", errors.New("SEED_USER_PASSWORD must be set (see server/.env.example) to run this seeder")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(s.seedUserPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("hash SEED_USER_PASSWORD: %w", err)
	}
	return strings.ReplaceAll(content, seedPasswordPlaceholder, string(hash)), nil
}

func seederNames() ([]string, error) {
	entries, err := fs.ReadDir(seedersFS, "seeders")
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".sql" {
			continue
		}
		names = append(names, entry.Name())
	}

	sort.Strings(names)
	return names, nil
}

func Seed() error {
	cfg := config.Load()
	db, err := Connect(cfg)
	if err != nil {
		return err
	}
	return NewSeederManager(db, cfg.SeedUserPassword).Run()
}
