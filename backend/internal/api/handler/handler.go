package handler

import (
	"ViDrop/internal/service"
)

type Handler struct {
	downloader *service.Downloader
}

func NewHandler(d *service.Downloader) *Handler {
	return &Handler{downloader: d}
}
