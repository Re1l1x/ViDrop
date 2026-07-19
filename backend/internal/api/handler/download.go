package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"ViDrop/internal/api/handler/dto"
	j "ViDrop/internal/job"
)

func (h *Handler) StartDownload(w http.ResponseWriter, r *http.Request) {
	var req dto.StartDownloadRequest

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

	res := dto.StartDownloadResponse{
		JobID: job.ID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) GetDownloadStatus(w http.ResponseWriter, r *http.Request) {
	jobID := r.PathValue("id")

	job, ok := h.jobs.Get(jobID)
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	res := dto.DownloadStatusResponse{
		Status:   string(job.Status),
		Progress: job.Progress,
	}

	if job.Status == j.Done {
		res.FileID = job.FileID
	}

	if job.Status == j.Error {
		res.Error = job.Error
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) GetDownloadProgress(w http.ResponseWriter, r *http.Request) {
	jobID := r.PathValue("id")

	job, ok := h.jobs.Get(jobID)
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

	send := func(event dto.DownloadProgressEvent) error {
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

	if err := send(dto.DownloadProgressEvent{
		Status:   string(job.Status),
		Progress: job.Progress,
		FileID:   job.FileID,
		Error:    job.Error,
	}); err != nil {
		return
	}

	for {
		select {
		case event := <-ch:
			if err := send(dto.DownloadProgressEvent{
				Status:   string(event.Status),
				Progress: event.Progress,
				FileID:   event.FileID,
				Error:    event.Error,
			}); err != nil {
				return
			}

		case <-r.Context().Done():
			return
		}
	}
}
