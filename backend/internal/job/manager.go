package job

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"
)

type Manager struct {
	jobs   map[string]*DownloadJob
	runner *Runner
	mu     sync.RWMutex
}

func NewManager(runner *Runner) *Manager {
	return &Manager{
		jobs:   make(map[string]*DownloadJob),
		runner: runner,
	}
}

func (m *Manager) StartDownload(url string, resolution int, audioBitrate int, format string) *DownloadJob {
	jobID := generateID()

	job := &DownloadJob{
		ID:           jobID,
		URL:          url,
		Resolution:   resolution,
		AudioBitrate: audioBitrate,
		Format:       format,
		Status:       Pending,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	m.mu.Lock()
	m.jobs[jobID] = job
	m.mu.Unlock()

	go m.runner.Run(job)

	return job
}

func generateID() string {
	b := make([]byte, 8)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func (m *Manager) Get(id string) (*DownloadJob, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	j, ok := m.jobs[id]
	return j, ok
}
