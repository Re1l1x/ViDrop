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

func (d *Downloader) Download(url string) (string, error) {
    filePath, err := d.yt.Download(url)
    if err != nil {
        return "", fmt.Errorf("download failed: %w", err)
    }

    fileID, err := d.storage.Save(filePath)
    if err != nil {
        return "", fmt.Errorf("save failed: %w", err)
    }

    return fileID, nil
}
