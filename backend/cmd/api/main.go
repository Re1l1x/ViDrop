package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"ViDrop/internal/api"
	"ViDrop/internal/api/handler"
	"ViDrop/internal/config"
	"ViDrop/internal/job"
	"ViDrop/internal/logger"
	"ViDrop/internal/middleware"
	"ViDrop/internal/service"
	"ViDrop/internal/storage"
	"ViDrop/internal/yt"
)

func main() {
	logger.Init()

	cfg := config.New()

	if err := os.MkdirAll(cfg.TempDir, 0755); err != nil {
		slog.Error("failed to create temp directory", "error", err)
		return
	}

	if err := os.MkdirAll(cfg.DownloadDir, 0755); err != nil {
		slog.Error("failed to create download directory", "error", err)
		return
	}

	ytClient := yt.New(cfg.TempDir)
	store := storage.NewLocalStorage(cfg.DownloadDir)

	broker := job.NewBroker()
	downloader := service.NewDownloader(ytClient, store)
	runner := job.NewRunner(downloader, broker)
	jobs := job.NewManager(runner)

	handler := handler.NewHandler(downloader, jobs, broker)

	mux := http.NewServeMux()
	api.RegisterRoutes(mux, handler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      middleware.Middleware(mux),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	slog.Info("server started", "addr", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		slog.Error("server stopped", "error", err)
	}
}
