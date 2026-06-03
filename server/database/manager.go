package database

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"gorm.io/gorm"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

const migrationsTableName = "schema_migrations"

type MigrationManager struct {
	db *gorm.DB
}

type MigrationStatus struct {
	Name  string
	Ran   bool
	Batch int
	RanAt *time.Time
}

func NewMigrationManager(db *gorm.DB) *MigrationManager {
	return &MigrationManager{db: db}
}

func (m *MigrationManager) Run() error {
	if err := m.ensureMigrationsTable(); err != nil {
		return err
	}

	names, err := migrationNames()
	if err != nil {
		return err
	}

	ran, err := m.ranMigrations()
	if err != nil {
		return err
	}

	batch, err := m.nextBatch()
	if err != nil {
		return err
	}

	for _, name := range names {
		if ran[name] {
			continue
		}

		if err := m.runMigration(name, batch); err != nil {
			return err
		}
	}

	return nil
}

func (m *MigrationManager) Status() ([]MigrationStatus, error) {
	if err := m.ensureMigrationsTable(); err != nil {
		return nil, err
	}

	names, err := migrationNames()
	if err != nil {
		return nil, err
	}

	records, err := m.migrationRecords()
	if err != nil {
		return nil, err
	}

	statuses := make([]MigrationStatus, 0, len(names))
	for _, name := range names {
		record, ok := records[name]
		status := MigrationStatus{Name: name, Ran: ok}
		if ok {
			status.Batch = record.Batch
			status.RanAt = &record.RanAt
		}
		statuses = append(statuses, status)
	}

	return statuses, nil
}

func (m *MigrationManager) ensureMigrationsTable() error {
	return m.db.Exec(fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  batch INT NOT NULL,
  ran_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  PRIMARY KEY (id),
  UNIQUE KEY idx_schema_migrations_name (name),
  KEY idx_schema_migrations_batch (batch)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
`, migrationsTableName)).Error
}

func (m *MigrationManager) nextBatch() (int, error) {
	var lastBatch sql.NullInt64
	if err := m.db.Raw(fmt.Sprintf("SELECT MAX(batch) FROM %s", migrationsTableName)).Scan(&lastBatch).Error; err != nil {
		return 0, err
	}
	if !lastBatch.Valid {
		return 1, nil
	}
	return int(lastBatch.Int64) + 1, nil
}

func (m *MigrationManager) ranMigrations() (map[string]bool, error) {
	records, err := m.migrationRecords()
	if err != nil {
		return nil, err
	}

	ran := make(map[string]bool, len(records))
	for name := range records {
		ran[name] = true
	}

	return ran, nil
}

func (m *MigrationManager) migrationRecords() (map[string]migrationRecord, error) {
	var rows []migrationRecord
	if err := m.db.Raw(fmt.Sprintf("SELECT name, batch, ran_at FROM %s ORDER BY id ASC", migrationsTableName)).Scan(&rows).Error; err != nil {
		return nil, err
	}

	records := make(map[string]migrationRecord, len(rows))
	for _, row := range rows {
		records[row.Name] = row
	}

	return records, nil
}

func (m *MigrationManager) runMigration(name string, batch int) error {
	content, err := migrationsFS.ReadFile(filepath.ToSlash(filepath.Join("migrations", name)))
	if err != nil {
		return err
	}

	return m.db.Transaction(func(tx *gorm.DB) error {
		for _, statement := range splitSQLStatements(string(content)) {
			if err := tx.Exec(statement).Error; err != nil {
				return fmt.Errorf("run migration %s: %w", name, err)
			}
		}

		if err := tx.Exec(
			fmt.Sprintf("INSERT INTO %s (name, batch) VALUES (?, ?)", migrationsTableName),
			name,
			batch,
		).Error; err != nil {
			return fmt.Errorf("record migration %s: %w", name, err)
		}

		return nil
	})
}

type migrationRecord struct {
	Name  string
	Batch int
	RanAt time.Time
}

func migrationNames() ([]string, error) {
	entries, err := fs.ReadDir(migrationsFS, "migrations")
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

func splitSQLStatements(sqlText string) []string {
	var statements []string
	var current strings.Builder
	var quote rune
	escaped := false

	for _, char := range sqlText {
		current.WriteRune(char)

		if quote != 0 {
			if escaped {
				escaped = false
				continue
			}
			if char == '\\' {
				escaped = true
				continue
			}
			if char == quote {
				quote = 0
			}
			continue
		}

		switch char {
		case '\'', '"', '`':
			quote = char
		case ';':
			statement := strings.TrimSpace(strings.TrimSuffix(current.String(), ";"))
			if statement != "" {
				statements = append(statements, statement)
			}
			current.Reset()
		}
	}

	statement := strings.TrimSpace(current.String())
	if statement != "" {
		statements = append(statements, statement)
	}

	return statements
}
