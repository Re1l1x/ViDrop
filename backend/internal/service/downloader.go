package service

import (
	"fmt"

	"ViDrop/internal/storage"
	"ViDrop/internal/yt"
)

type Downloader struct {
	yt      *yt.YtDlp
	storage storage.Storage
}

func NewDownloader(ytClient *yt.YtDlp, storage storage.Storage) *Downloader {
	return &Downloader{
		yt:      ytClient,
		storage: storage,
	}
}

func (d *Downloader) GetInfo(url string) (yt.VideoInfo, error) {
	return d.yt.GetInfo(url)
}

func (d *Downloader) Download(url string, resolution int, audioBitrate int, format string) (string, error) {
	filePath, err := d.yt.Download(url, resolution, audioBitrate, format)
	if err != nil {
		return "", fmt.Errorf("download failed: %w", err)
	}

	fileID, err := d.storage.Save(filePath)
	if err != nil {
		return "", fmt.Errorf("save failed: %w", err)
	}

	return fileID, nil
}

func (d *Downloader) GetFilePath(fileID string) (string, error) {
	return d.storage.Get(fileID)
}
