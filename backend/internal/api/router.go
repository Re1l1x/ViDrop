package api

import (
	"net/http"

	"ViDrop/internal/api/handler"
)

func RegisterRoutes(mux *http.ServeMux, handler *handler.Handler) {
	mux.HandleFunc("/info", handler.GetVideoInfo)
	mux.HandleFunc("/download", handler.Download)
	mux.HandleFunc("/file/", handler.GetFile)
}
