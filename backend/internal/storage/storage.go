package storage

type Storage interface {
	Save(filePath string) (string, error)
	Get(fileID string) (string, error)
	Delete(fileID string) error
}
