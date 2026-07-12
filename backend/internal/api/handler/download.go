package handler

import (
	"encoding/json"
	"net/http"
	"strings"

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

func (h *Handler) GetDownloadStatus(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/download/")

	task, ok := h.jobs.Get(id)
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
