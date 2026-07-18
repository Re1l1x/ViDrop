package api

import (
	"net/http"

	"ViDrop/internal/api/handler"
)

func RegisterRoutes(mux *http.ServeMux, handler *handler.Handler) {
	mux.HandleFunc("POST /info", handler.GetVideoInfo)

	mux.HandleFunc("POST /download", handler.StartDownload)
	mux.HandleFunc("GET /download/{id}", handler.GetDownloadStatus)
	mux.HandleFunc("GET /download/{id}/events", handler.GetDownloadProgress)

	mux.HandleFunc("GET /file/{id}", handler.GetFile)
}
