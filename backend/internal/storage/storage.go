package storage

type Storage interface {
	Save(tempPath string, fileName string) (string, error)
	Get(fileID string) (string, error)
	Delete(fileID string) error
}
