package yt

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
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

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	out, err := cmd.Output()
	if err != nil {
		return VideoInfo{}, fmt.Errorf("yt-dlp: get video info: %w: %s", err, stderr.String())
	}

	var raw ytResponse
	if err := json.Unmarshal(out, &raw); err != nil {
		return VideoInfo{}, fmt.Errorf("yt-dlp: parse response: %w", err)
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

func (y *YtDlp) Download(url string, resolution int, audioBitrate int, format string, fileName string, onProgress func(int)) (string, error) {
	formatSelector := fmt.Sprintf(
		"bestvideo[height<=%d]+bestaudio[abr<=%d]/best",
		resolution,
		audioBitrate+1,
	)

	outputPath := filepath.Join(y.outputDir, fileName)

	cmd := exec.Command(
		"yt-dlp",
		"--newline",
		"--progress",
		"-f", formatSelector,
		"-o", outputPath,
		"--merge-output-format", format,
		url,
	)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("yt-dlp: create stdout pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("yt-dlp: start process: %w", err)
	}

	scanner := bufio.NewScanner(stdout)

	const wideoWeight = 0.9
	const audioWeight = 0.1
	var stage int
	var lastProgress int

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "Destination") {
			stage++
		}

		if strings.Contains(line, "%") {
			percent := parsePercent(line)
			if percent >= 0 {
				var total int

				if stage == 1 {
					total = int(float64(percent) * wideoWeight)
				} else {
					total = 90 + int(float64(percent)*audioWeight)
				}

				if total != lastProgress {
					lastProgress = total
					onProgress(total)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("yt-dlp: read output: %w", err)
	}

	if err := cmd.Wait(); err != nil {
		return "", fmt.Errorf("yt-dlp: download failed: %w", err)
	}

	return outputPath, nil
}

func (y *YtDlp) GetVideoID(url string) (string, error) {
	cmd := exec.Command("yt-dlp", "--print", "id", url)

	id, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("yt-dlp: get video id: %w", err)
	}

	return strings.TrimSpace(string(id)), nil
}

func parsePercent(line string) int {
	re := regexp.MustCompile(`(\d+\.?\d*)%`)
	match := re.FindStringSubmatch(line)

	if len(match) < 2 {
		return -1
	}

	f, err := strconv.ParseFloat(match[1], 64)
	if err != nil {
		return -1
	}

	return int(f)
}
