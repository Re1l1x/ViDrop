package job

import "time"

type Status string

const (
	Pending     Status = "pending"
	Downloading Status = "downloading"
	Done        Status = "done"
	Error       Status = "error"
)

type DownloadJob struct {
	ID           string
	FileID       string
	URL          string
	Resolution   int
	AudioBitrate int
	Format       string

	Status    Status
	Progress  int
	Error     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
