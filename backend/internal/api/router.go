package api

import (
    "net/http"

    "ViDrop/internal/api/handlers"
)

func RegisterRoutes(mux *http.ServeMux, downloadHandler *handlers.DownloadHandler) {
    mux.HandleFunc("/download", downloadHandler.Download)
}
