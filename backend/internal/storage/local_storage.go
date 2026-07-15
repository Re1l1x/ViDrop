package storage

import (
	"fmt"
	"os"
	"path/filepath"
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

	return fileName, nil
}

func (s *LocalStorage) Get(fileID string) (string, error) {
	path := filepath.Join(s.basePath, fileID)

	if _, err := os.Stat(path); err != nil {
		return "", fmt.Errorf("storage: check file: %w", err)
	}

	return path, nil
}

func (s *LocalStorage) Delete(fileID string) error {
	path := filepath.Join(s.basePath, fileID)

	if err := os.Remove(path); err != nil {
		return fmt.Errorf("storage: delete file: %w", err)
	}

	return nil
}
