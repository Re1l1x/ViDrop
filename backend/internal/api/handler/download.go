package handler

import (
	"encoding/json"
	"net/http"

	"ViDrop/internal/api/handler/dto"
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

	json.NewEncoder(w).Encode(map[string]string{
		"job_id": job.ID,
	})
}
