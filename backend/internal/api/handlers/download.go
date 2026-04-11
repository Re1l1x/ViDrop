package handlers

import (
    "strings"
    "encoding/json"
    "net/http"

    "ViDrop/internal/service"
)

type DownloadRequest struct {
    URL string `json:"url"`
}

type DownloadResponse struct {
    FileID      string `json:"file_id"`
    DownloadURL string `json:"download_url"`
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

    res := DownloadResponse{
        FileID:      fileID,
        DownloadURL: "/file/" + fileID,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(res)
}

func (h *DownloadHandler) GetFile(w http.ResponseWriter, r *http.Request) {
    fileID := strings.TrimPrefix(r.URL.Path, "/file/")

    path, err := h.downloader.GetFilePath(fileID)
    if err != nil {
        http.Error(w, "file not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Disposition", "attachment; filename=\""+fileID+"\"")
    http.ServeFile(w, r, path)
}
