package handlers

import (
    "strings"
    "encoding/json"
    "net/http"

    "ViDrop/internal/service"
)

type InfoRequest struct {
    URL string `json:"url"`
}

type InfoResponse struct {
    Title     string `json:"title"`
	Thumbnail string `json:"thumbnail"`
}

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

func (h *DownloadHandler) GetInfo(w http.ResponseWriter, r *http.Request) {
    var req InfoRequest

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

    res := InfoResponse{
        Title:     info.Title,
        Thumbnail: info.Thumbnail,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(res)
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
