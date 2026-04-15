package config

import (
    "os"
    "path/filepath"
)

type Config struct {
    TempDir      string
    DownloadDir  string
}

func New() *Config {
    wd, _ := os.Getwd()

    return &Config{
        TempDir:     filepath.Join(wd, "temp"),
        DownloadDir: filepath.Join(wd, "downloads"),
    }
}
