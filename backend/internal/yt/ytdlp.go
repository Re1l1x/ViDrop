package yt

import (
	"strings"
	"os/exec"
	"path/filepath"
)

type YtDlp struct {
	outputDir string
}

func New(outputDir string) *YtDlp {
	return &YtDlp{outputDir: outputDir}
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
