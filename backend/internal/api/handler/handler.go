package handler

import (
	"ViDrop/internal/job"
	"ViDrop/internal/service"
)

type Handler struct {
	downloader *service.Downloader
	jobs       *job.Manager
	broker     *job.Broker
}

func NewHandler(d *service.Downloader, jobs *job.Manager, broker *job.Broker) *Handler {
	return &Handler{
		downloader: d,
		jobs:       jobs,
		broker:     broker,
	}
}
