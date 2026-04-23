package handler

import (
	"net/http"
	"strings"
)

func (h *Handler) GetFile(w http.ResponseWriter, r *http.Request) {
	fileID := strings.TrimPrefix(r.URL.Path, "/file/")

	path, err := h.downloader.GetFilePath(fileID)
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=\""+fileID+"\"")
	http.ServeFile(w, r, path)
}
