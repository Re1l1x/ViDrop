package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	TempDir  string
	MediaDir string
}

func New() *Config {
	wd, _ := os.Getwd()

	return &Config{
		TempDir:  filepath.Join(wd, "storage/temp"),
		MediaDir: filepath.Join(wd, "storage/media"),
	}
}
