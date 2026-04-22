package yt

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

type YtDlp struct {
	outputDir string
}

func New(outputDir string) *YtDlp {
	return &YtDlp{outputDir: outputDir}
}

type VideoInfo struct {
	Title         string `json:"title"`
	Thumbnail     string `json:"thumbnail"`
	Resolutions   []int  `json:"resolutions"`
	AudioBitrates []int  `json:"audio_bitrates"`
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

	resolutions := make([]int, 0, len(resMap))
	audioBitrates := make([]int, 0, len(audioMap))

	for r := range resMap {
		resolutions = append(resolutions, r)
	}
	for a := range audioMap {
		audioBitrates = append(audioBitrates, a)
	}

	sort.Ints(resolutions)
	sort.Ints(audioBitrates)

	return VideoInfo{
		Title:         raw.Title,
		Thumbnail:     raw.Thumbnail,
		Resolutions:   resolutions,
		AudioBitrates: audioBitrates,
	}, nil
}

func (y *YtDlp) Download(url string, resolution int, audioBitrate int) (string, error) {
	cmdID := exec.Command("yt-dlp", "--print", "id", url)
	idOut, err := cmdID.Output()
	if err != nil {
		return "", err
	}

	videoID := strings.TrimSpace(string(idOut))

	format := fmt.Sprintf(
		"bestvideo[height<=%d]+bestaudio[abr<=%d]/best",
		resolution,
		audioBitrate+1,
	)

	outputTemplate := filepath.Join(y.outputDir, videoID+".mp4")

	cmd := exec.Command(
		"yt-dlp",
		"-f", format,
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
