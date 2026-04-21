package api

import (
	"net/http"

	"ViDrop/internal/api/handlers"
)

func RegisterRoutes(mux *http.ServeMux, downloadHandler *handlers.DownloadHandler) {
	mux.HandleFunc("/info", downloadHandler.GetInfo)
	mux.HandleFunc("/download", downloadHandler.Download)
	mux.HandleFunc("/file/", downloadHandler.GetFile)
}
