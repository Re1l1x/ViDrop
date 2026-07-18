package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ViDrop/internal/api/handler/dto"
	"ViDrop/internal/job"
)

func (h *Handler) StartDownload(w http.ResponseWriter, r *http.Request) {
	var req dto.DownloadRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	job := h.jobs.StartDownload(
		req.URL,
		req.Resolution,
		req.AudioBitrate,
		req.Format,
	)

	res := map[string]string{
		"task_id": job.ID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) GetDownloadProgress(w http.ResponseWriter, r *http.Request) {
	jobID := r.PathValue("id")

	task, ok := h.jobs.Get(jobID)
	if !ok {
		http.Error(w, "job not found", http.StatusNotFound)
		return
	}

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	ch, unsubscribe := h.broker.Subscribe(jobID)
	defer unsubscribe()

	send := func(event job.DownloadEvent) error {
		data, err := json.Marshal(event)
		if err != nil {
			return err
		}

		if _, err := fmt.Fprintf(w, "data: %s\n\n", data); err != nil {
			return err
		}

		flusher.Flush()

		return nil
	}

	if err := send(job.DownloadEvent{
		Status:   task.Status,
		Progress: task.Progress,
		FileID:   task.FileID,
		Error:    task.Error,
	}); err != nil {
		return
	}

	for {
		select {
		case event := <-ch:
			if err := send(event); err != nil {
				return
			}

		case <-r.Context().Done():
			return
		}
	}
}

func (h *Handler) GetDownloadStatus(w http.ResponseWriter, r *http.Request) {
	jobID := r.PathValue("id")

	task, ok := h.jobs.Get(jobID)
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	res := map[string]interface{}{
		"status":   task.Status,
		"progress": task.Progress,
	}

	if task.Status == job.Done {
		res["file_id"] = task.FileID
		res["download_url"] = "/file/" + task.FileID
	}

	if task.Status == job.Error {
		res["error"] = task.Error
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
