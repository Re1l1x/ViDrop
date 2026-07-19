package job

import (
	"log/slog"
	"sync"
)

type ProgressEvent struct {
	Status   Status
	Progress int
	FileID   string
	Error    string
}

type Broker struct {
	mu sync.RWMutex

	subscribers map[string]map[chan ProgressEvent]struct{}
}

func NewBroker() *Broker {
	return &Broker{
		subscribers: make(map[string]map[chan ProgressEvent]struct{}),
	}
}

func (b *Broker) Subscribe(jobID string) (<-chan ProgressEvent, func()) {
	ch := make(chan ProgressEvent, 8)

	b.mu.Lock()
	defer b.mu.Unlock()

	if _, ok := b.subscribers[jobID]; !ok {
		b.subscribers[jobID] = make(map[chan ProgressEvent]struct{})
	}

	b.subscribers[jobID][ch] = struct{}{}

	slog.Debug("sse client subscribed", "job_id", jobID, "subscribers", len(b.subscribers[jobID]))

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

func (b *Broker) Publish(jobID string, event ProgressEvent) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for ch := range b.subscribers[jobID] {
		select {
		case ch <- event:
		default:
		}
	}
}
