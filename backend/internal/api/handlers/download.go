package handlers

import (
    "encoding/json"
    "net/http"

    "ViDrop/internal/service"
)

type DownloadRequest struct {
    URL string `json:"url"`
}

type DownloadResponse struct {
    FileID string `json:"file_id"`
}

type DownloadHandler struct {
    downloader *service.Downloader
}

func NewDownloadHandler(d *service.Downloader) *DownloadHandler {
    return &DownloadHandler{downloader: d}
}

func (h *DownloadHandler) Download(w http.ResponseWriter, r *http.Request) {
    var req DownloadRequest

    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "invalid request", http.StatusBadRequest)
        return
    }

    fileID, err := h.downloader.Download(req.URL)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    res := DownloadResponse{FileID: fileID}

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(res)
}
