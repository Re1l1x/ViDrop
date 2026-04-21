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
	os.MkdirAll(basePath, 0755)

	return &LocalStorage{basePath: basePath}
}

func (s *LocalStorage) Save(tempPath string) (string, error) {
	fileName := filepath.Base(tempPath)

	destPath := filepath.Join(s.basePath, fileName)

	err := os.Rename(tempPath, destPath)
	if err != nil {
		return "", fmt.Errorf("failed to move file: %w", err)
	}

	return fileName, nil
}

func (s *LocalStorage) Get(fileID string) (string, error) {
	path := filepath.Join(s.basePath, fileID)

	if _, err := os.Stat(path); err != nil {
		return "", err
	}

	return path, nil
}

func (s *LocalStorage) Delete(fileID string) error {
	path := filepath.Join(s.basePath, fileID)
	return os.Remove(path)
}
