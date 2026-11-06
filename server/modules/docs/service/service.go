package service

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type DocsService interface {
	Index() ([]byte, error)
	Show(slug string) ([]byte, error)
}

type docsService struct {
	dir string
}

var _ DocsService = (*docsService)(nil)

func NewDocsService(dir string) DocsService {
	return &docsService{dir: dir}
}

func (s *docsService) Index() ([]byte, error) {
	return os.ReadFile(filepath.Join(s.dir, "index.json"))
}

func (s *docsService) Show(slug string) ([]byte, error) {
	slug = strings.Trim(slug, "/")
	if slug == "" || strings.ContainsRune(slug, 0x00) {
		return nil, errors.New("invalid docs slug")
	}
	// Force relative path: prefix with "/" so filepath.Clean collapses any leading "..".
	cleaned := filepath.Clean("/" + slug)
	if strings.Contains(cleaned, "..") {
		return nil, errors.New("invalid docs slug")
	}
	full := filepath.Join(s.dir, cleaned+".md")
	rel, err := filepath.Rel(s.dir, full)
	if err != nil || strings.HasPrefix(rel, "..") || filepath.IsAbs(rel) {
		return nil, errors.New("invalid docs slug")
	}
	return os.ReadFile(full)
}
