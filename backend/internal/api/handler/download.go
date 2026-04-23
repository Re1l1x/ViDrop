package handler

import (
	"encoding/json"
	"net/http"

	"ViDrop/internal/api/handler/dto"
)

func (h *Handler) Download(w http.ResponseWriter, r *http.Request) {
	var req dto.DownloadRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	fileID, err := h.downloader.Download(req.URL, req.Resolution, req.AudioBitrate, req.Format)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := dto.DownloadResponse{
		FileID:      fileID,
		DownloadURL: "/file/" + fileID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
