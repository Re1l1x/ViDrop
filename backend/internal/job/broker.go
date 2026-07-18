package job

import (
	"log/slog"
	"sync"
)

type DownloadEvent struct {
	FileID   string `json:"file_id,omitempty"`
	Status   Status `json:"status"`
	Progress int    `json:"progress"`
	Error    string `json:"error,omitempty"`
}

type Broker struct {
	mu sync.RWMutex

	subscribers map[string]map[chan DownloadEvent]struct{}
}

func NewBroker() *Broker {
	return &Broker{
		subscribers: make(map[string]map[chan DownloadEvent]struct{}),
	}
}

func (b *Broker) Subscribe(jobID string) (<-chan DownloadEvent, func()) {
	ch := make(chan DownloadEvent, 8)

	b.mu.Lock()
	defer b.mu.Unlock()

	if _, ok := b.subscribers[jobID]; !ok {
		b.subscribers[jobID] = make(map[chan DownloadEvent]struct{})
	}

	b.subscribers[jobID][ch] = struct{}{}

	slog.Info("sse client subscribed", "job_id", jobID, "subscribers", len(b.subscribers[jobID]))

	unsubscribe := func() {
		b.mu.Lock()
		defer b.mu.Unlock()

		delete(b.subscribers[jobID], ch)

		if len(b.subscribers[jobID]) == 0 {
			delete(b.subscribers, jobID)
		}
	}

	return ch, unsubscribe
}

func (b *Broker) Publish(jobID string, event DownloadEvent) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for ch := range b.subscribers[jobID] {
		select {
		case ch <- event:
		default:
		}
	}
}
