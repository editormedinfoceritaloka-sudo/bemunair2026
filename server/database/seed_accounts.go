package database

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"bemunair2026/server/config"
	"bemunair2026/server/database/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func ProvisionMinistryAccounts() error {
	cfg := config.Load()
	if cfg.SeedUserPassword == "" {
		return errors.New("set SEED_USER_PASSWORD to the previous shared password before provisioning ministry accounts")
	}
	db, err := Connect(cfg)
	if err != nil {
		return err
	}
	return NewSeederManager(db, cfg.SeedUserPassword).provisionAccounts(true)
}

func (s *SeederManager) provisionAccounts(leadersOnly bool) error {
	if err := os.MkdirAll(s.credentialDir, 0700); err != nil {
		return err
	}
	dir, err := os.Lstat(s.credentialDir)
	if err != nil {
		return err
	}
	if !dir.IsDir() || dir.Mode().Perm()&0077 != 0 {
		return fmt.Errorf("credential directory %s must be private (0700)", s.credentialDir)
	}
	for _, account := range seedAccounts {
		if leadersOnly && !account.leader {
			continue
		}
		if err := s.provisionAccount(account, leadersOnly); err != nil {
			return fmt.Errorf("provision %s: %w", account.email, err)
		}
	}
	return nil
}

func (s *SeederManager) provisionAccount(account seedAccount, leadersOnly bool) error {
	var ministryID *uint64
	if account.unit != "" {
		var ministry entities.Ministry
		err := s.db.Where("code = ?", account.unit).First(&ministry).Error
		if leadersOnly && errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("%s: skipped, unit %s is missing", account.email, account.unit)
			return nil
		}
		if err != nil {
			return fmt.Errorf("find unit %s: %w", account.unit, err)
		}
		ministryID = &ministry.ID
	}
	var user entities.User
	err := s.db.Where("email = ?", account.email).First(&user).Error
	exists := err == nil
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if exists {
		if leadersOnly && (user.MinistryID == nil || *user.MinistryID != *ministryID || user.Role != account.role) {
			return fmt.Errorf("existing account has unexpected role or ministry")
		}
		if s.legacyPassword == "" || bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(s.legacyPassword)) != nil {
			if leadersOnly {
				log.Printf("%s: kept existing password (does not match shared password)", account.email)
			}
			return nil
		}
	}
	password, err := seedCredential(s.credentialDir, account.email)
	if err != nil {
		return err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	if exists {
		// Keep a password changed after our initial read.
		result := s.db.Model(&user).Where("password_hash = ?", user.PasswordHash).Update("password_hash", string(hash))
		if result.Error != nil {
			return result.Error
		}
		if leadersOnly {
			if result.RowsAffected == 1 {
				log.Printf("%s: replaced shared password", account.email)
			} else {
				log.Printf("%s: kept concurrently changed password", account.email)
			}
		}
		return nil
	}
	user = entities.User{Name: account.name, Email: account.email, Role: account.role, MinistryID: ministryID, PasswordHash: string(hash)}
	if !leadersOnly {
		user.Phone = &account.phone
	}
	if account.unit != "" {
		user.Ministry = &account.unit
	}
	if err := s.db.Create(&user).Error; err != nil {
		return err
	}
	if leadersOnly {
		log.Printf("%s: created", account.email)
	}
	return nil
}

// Publish the complete credential atomically before writing its hash. Retries
// reuse the same credential, including after a failed database operation.
func seedCredential(dir, email string) (string, error) {
	path := filepath.Join(dir, email)
	info, err := os.Lstat(path)
	if err == nil {
		if !info.Mode().IsRegular() || info.Mode().Perm() != 0600 {
			return "", fmt.Errorf("credential %s must be a regular file with mode 0600", path)
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return "", err
		}
		password := strings.TrimSuffix(string(content), "\n")
		raw, err := base64.RawURLEncoding.DecodeString(password)
		if err != nil || len(raw) != 18 {
			return "", fmt.Errorf("invalid generated credential in %s", path)
		}
		return password, syncCredentialDir(dir)
	}
	if !errors.Is(err, os.ErrNotExist) {
		return "", err
	}
	var random [18]byte
	if _, err := rand.Read(random[:]); err != nil {
		return "", err
	}
	password := base64.RawURLEncoding.EncodeToString(random[:])
	tmp, err := os.CreateTemp(dir, ".credential-*")
	if err != nil {
		return "", err
	}
	defer os.Remove(tmp.Name())
	_, err = tmp.WriteString(password + "\n")
	if err == nil {
		err = tmp.Sync()
	}
	if closeErr := tmp.Close(); closeErr != nil {
		return "", closeErr
	}
	if err != nil {
		return "", err
	}
	if err := os.Link(tmp.Name(), path); err != nil {
		if errors.Is(err, os.ErrExist) {
			return seedCredential(dir, email)
		}
		return "", err
	}
	return password, syncCredentialDir(dir)
}

func syncCredentialDir(path string) error {
	dir, err := os.Open(path)
	if err != nil {
		return err
	}
	defer dir.Close()
	return dir.Sync()
}
