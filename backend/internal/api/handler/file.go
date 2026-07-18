package handler

import (
	"net/http"
)

func (h *Handler) GetFile(w http.ResponseWriter, r *http.Request) {
	fileID := r.PathValue("id")

	path, err := h.downloader.GetFilePath(fileID)
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=\""+fileID+"\"")
	http.ServeFile(w, r, path)
}
