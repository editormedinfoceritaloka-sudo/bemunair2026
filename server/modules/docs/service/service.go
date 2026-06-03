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
	if strings.Contains(slug, "..") {
		return nil, errors.New("invalid docs slug")
	}
	return os.ReadFile(filepath.Join(s.dir, slug+".md"))
}
