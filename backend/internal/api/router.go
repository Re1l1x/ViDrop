package api

import (
	"net/http"

	"ViDrop/internal/api/handler"
)

func RegisterRoutes(mux *http.ServeMux, handler *handler.Handler) {
	mux.HandleFunc("/info", handler.GetVideoInfo)
	mux.HandleFunc("POST /download", handler.StartDownload)
	mux.HandleFunc("GET /download/", handler.GetDownloadStatus)
	mux.HandleFunc("/file/", handler.GetFile)
}
