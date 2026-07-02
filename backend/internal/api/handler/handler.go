package handler

import (
	"ViDrop/internal/job"
	"ViDrop/internal/service"
)

type Handler struct {
	downloader *service.Downloader
	jobs       *job.Manager
}

func NewHandler(d *service.Downloader, jobs *job.Manager) *Handler {
	return &Handler{
		downloader: d,
		jobs:       jobs,
	}
}
