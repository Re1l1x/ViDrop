package handler

import (
	"net/http"
	"path/filepath"
)

func (h *Handler) GetFile(w http.ResponseWriter, r *http.Request) {
	fileID := r.PathValue("id")

	filePath, err := h.downloader.GetFilePath(fileID)
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}

	fileName := filepath.Base(filePath)

	w.Header().Set("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	http.ServeFile(w, r, filePath)
}
