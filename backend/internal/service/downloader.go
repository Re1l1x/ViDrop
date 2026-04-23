package service

import (
	"crypto/sha256"
	"encoding/hex"
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
	videoID, err := d.yt.GetVideoID(url)
	if err != nil {
		return "", err
	}

	fileID := generateFileID(videoID, resolution, audioBitrate, format)
	fileName := fileID + "." + format

	tempPath, err := d.yt.Download(url, resolution, audioBitrate, format, fileName)
	if err != nil {
		return "", fmt.Errorf("download failed: %w", err)
	}

	_, err = d.storage.Save(tempPath, fileName)
	if err != nil {
		return "", fmt.Errorf("save failed: %w", err)
	}

	return fileID, nil
}

func generateFileID(videoID string, resolution int, audioBitrate int, format string) string {
	raw := fmt.Sprintf("%s|%d|%d|%s", videoID, resolution, audioBitrate, format)

	hash := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(hash[:])[:16]
}

func (d *Downloader) GetFilePath(fileID string) (string, error) {
	return d.storage.Get(fileID)
}
