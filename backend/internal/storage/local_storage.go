package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type LocalStorage struct {
	basePath string
}

func NewLocalStorage(basePath string) *LocalStorage {
	return &LocalStorage{basePath: basePath}
}

func (s *LocalStorage) Save(tempPath string, fileName string) (string, error) {
	destPath := filepath.Join(s.basePath, fileName)

	if err := os.Rename(tempPath, destPath); err != nil {
		return "", fmt.Errorf("storage: move file: %w", err)
	}

	return destPath, nil
}

func (s *LocalStorage) Get(fileID string) (string, error) {
	entries, err := os.ReadDir(s.basePath)
	if err != nil {
		return "", fmt.Errorf("storage: read dir: %w", err)
	}

	prefix := fileID + "."

	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), prefix) {
			return filepath.Join(s.basePath, entry.Name()), nil
		}
	}

	return "", fmt.Errorf("storage: file not found")
}

func (s *LocalStorage) Delete(fileID string) error {
	path := filepath.Join(s.basePath, fileID)

	if err := os.Remove(path); err != nil {
		return fmt.Errorf("storage: delete file: %w", err)
	}

	return nil
}
