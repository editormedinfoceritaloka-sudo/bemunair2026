package database

import (
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"
	"sort"

	"bemunair2026/server/config"
	"gorm.io/gorm"
)

//go:embed seeders/*.sql
var seedersFS embed.FS

type SeederManager struct {
	db *gorm.DB
}

func NewSeederManager(db *gorm.DB) *SeederManager {
	return &SeederManager{db: db}
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

	return s.db.Transaction(func(tx *gorm.DB) error {
		for _, statement := range splitSQLStatements(string(content)) {
			if err := tx.Exec(statement).Error; err != nil {
				return fmt.Errorf("run seeder %s: %w", name, err)
			}
		}
		return nil
	})
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
	return NewSeederManager(db).Run()
}
