package main

import (
	"log"
	"net/http"
	"time"

	"ViDrop/internal/api"
	"ViDrop/internal/api/handler"
	"ViDrop/internal/config"
	"ViDrop/internal/job"
	"ViDrop/internal/middleware"
	"ViDrop/internal/service"
	"ViDrop/internal/storage"
	"ViDrop/internal/yt"
)

func main() {
	cfg := config.New()

	ytClient := yt.New(cfg.TempDir)
	store := storage.NewLocalStorage(cfg.DownloadDir)

	downloader := service.NewDownloader(ytClient, store)
	runner := job.NewRunner(downloader)
	jobs := job.NewManager(runner)

	handler := handler.NewHandler(downloader, jobs)

	mux := http.NewServeMux()
	api.RegisterRoutes(mux, handler)

	log.Println("Server running on http://localhost:8080")

	server := &http.Server{
		Addr:         ":8080",
		Handler:      middleware.Middleware(mux),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
