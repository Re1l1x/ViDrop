package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log/slog"

	"ViDrop/internal/storage/filestorage"
	"ViDrop/internal/yt"
)

type Downloader struct {
	yt      *yt.YtDlp
	storage filestorage.FileStorage
}

func NewDownloader(ytClient *yt.YtDlp, storage filestorage.FileStorage) *Downloader {
	return &Downloader{
		yt:      ytClient,
		storage: storage,
	}
}

func (d *Downloader) GetInfo(url string) (yt.VideoInfo, error) {
	slog.Debug("getting video info", "url", url)

	info, err := d.yt.GetInfo(url)

	if err != nil {
		slog.Error("failed to get video info", "error", err)

		return yt.VideoInfo{}, fmt.Errorf("service: get video info: %w", err)
	}

	slog.Debug("video info received", "title", info.Title)

	return info, nil
}

func (d *Downloader) Download(url string, resolution int, audioBitrate int, format string, onProgress func(int)) (string, error) {
	videoID, err := d.yt.GetVideoID(url)
	if err != nil {
		return "", fmt.Errorf("service: get video id: %w", err)
	}

	fileID := GenerateFileID(videoID, resolution, audioBitrate, format)
	fileName := fileID + "." + format

	tempPath, err := d.yt.Download(url, resolution, audioBitrate, format, fileName, onProgress)
	if err != nil {
		return "", fmt.Errorf("service: download video: %w", err)
	}

	if _, err := d.storage.Save(tempPath, fileName); err != nil {
		return "", fmt.Errorf("service: save file: %w", err)
	}

	return fileID, nil
}

func GenerateFileID(videoID string, resolution int, audioBitrate int, format string) string {
	raw := fmt.Sprintf("%s|%d|%d|%s", videoID, resolution, audioBitrate, format)

	hash := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(hash[:])[:16]
}

func (d *Downloader) GetFilePath(fileID string) (string, error) {
	filePath, err := d.storage.Get(fileID)
	if err != nil {
		slog.Error("failed to get file", "file_id", fileID, "error", err)

		return "", fmt.Errorf("service: get file: %w", err)
	}

	slog.Debug("file found", "file_id", fileID)

	return filePath, nil
}
