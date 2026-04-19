package yt

import (
    "fmt"
	"strings"
	"os/exec"
    "encoding/json"
	"path/filepath"
    "sort"
)

type YtDlp struct {
	outputDir string
}

func New(outputDir string) *YtDlp {
	return &YtDlp{outputDir: outputDir}
}

type VideoInfo struct {
	Title          string   `json:"title"`
	Thumbnail      string   `json:"thumbnail"`
	Resolutions    []string `json:"resolutions"`
	AudioBitrates  []string `json:"audio_bitrates"`
}

type ytResponse struct {
	Title     string   `json:"title"`
	Thumbnail string   `json:"thumbnail"`
	Formats   []format `json:"formats"`
}

type format struct {
	Height int     `json:"height"`
	VCodec string  `json:"vcodec"`
	ACodec string  `json:"acodec"`
	ABR    float64 `json:"abr"`
}

func (y *YtDlp) GetInfo(url string) (VideoInfo, error) {
	cmd := exec.Command("yt-dlp", "-j", url)

	out, err := cmd.Output()
	if err != nil {
		return VideoInfo{}, err
	}

	var raw ytResponse
	if err := json.Unmarshal(out, &raw); err != nil {
		return VideoInfo{}, err
	}

	resMap := make(map[int]struct{})
	audioMap := make(map[int]struct{})

	for _, f := range raw.Formats {
		if f.VCodec != "none" && f.ACodec == "none" && f.Height > 0 {
			resMap[f.Height] = struct{}{}
		}

		if f.VCodec == "none" && f.ACodec != "none" && f.ABR > 0 {
			audioMap[int(f.ABR)] = struct{}{}
		}
	}

	return VideoInfo{
		Title:         raw.Title,
		Thumbnail:     raw.Thumbnail,
		Resolutions:   mapToSortedStrings(resMap, "p"),
		AudioBitrates: mapToSortedStrings(audioMap, "kbps"),
	}, nil
}

func mapToSortedStrings(m map[int]struct{}, suffix string) []string {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	result := make([]string, 0, len(keys))
	for _, k := range keys {
		result = append(result, fmt.Sprintf("%d%s", k, suffix))
	}

	return result
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
