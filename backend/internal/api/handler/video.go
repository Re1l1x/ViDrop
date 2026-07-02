package handler

import (
	"encoding/json"
	"net/http"

	"ViDrop/internal/api/handler/dto"
)

func (h *Handler) GetVideoInfo(w http.ResponseWriter, r *http.Request) {
	var req dto.InfoRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	info, err := h.downloader.GetInfo(req.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := dto.InfoResponse{
		Title:         info.Title,
		Thumbnail:     info.Thumbnail,
		Resolutions:   info.Resolutions,
		AudioBitrates: info.AudioBitrates,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
