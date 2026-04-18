package yt

import (
	"strings"
	"os/exec"
    "encoding/json"
	"path/filepath"
)

type YtDlp struct {
	outputDir string
}

func New(outputDir string) *YtDlp {
	return &YtDlp{outputDir: outputDir}
}

type VideoInfo struct {
    Title     string `json:"title"`
    Thumbnail string `json:"thumbnail"`
}

func (y *YtDlp) GetInfo(url string) (VideoInfo, error) {
    cmd := exec.Command("yt-dlp", "-j", url)

    out, err := cmd.Output()
    if err != nil {
        return VideoInfo{}, err
    }

    var data VideoInfo

    if err := json.Unmarshal(out, &data); err != nil {
        return VideoInfo{}, err
    }

    return VideoInfo{
        Title:     data.Title,
        Thumbnail: data.Thumbnail,
    }, nil
}

func (y *YtDlp) Download(url string) (string, error) {
	cmdID := exec.Command("yt-dlp", "--print", "id", url)
	idOut, err := cmdID.Output()
	if err != nil {
		return "", err
	}

	videoID := strings.TrimSpace(string(idOut))

	outputTemplate := filepath.Join(y.outputDir, videoID+".mp4")

	cmd := exec.Command(
        "yt-dlp",
        "-o", outputTemplate,
        "--merge-output-format", "mp4",
        url,
    )

	_, err = cmd.CombinedOutput()
    if err != nil {
        return "", err
    }

	filePath := filepath.Join(y.outputDir, videoID+".mp4")

	return filePath, nil
}
