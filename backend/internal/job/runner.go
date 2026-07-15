package job

import (
	"log/slog"
	"time"

	"ViDrop/internal/service"
)

type Runner struct {
	downloader *service.Downloader
}

func NewRunner(d *service.Downloader) *Runner {
	return &Runner{
		downloader: d,
	}
}

func (r *Runner) Run(j *DownloadJob) {
	j.Status = Downloading
	j.UpdatedAt = time.Now()

	slog.Info(
		"download job started",
		"job_id", j.ID,
		"url", j.URL,
		"resolution", j.Resolution,
		"audio_bitrate", j.AudioBitrate,
		"format", j.Format,
	)

	fileID, err := r.downloader.Download(
		j.URL,
		j.Resolution,
		j.AudioBitrate,
		j.Format,
		func(p int) {
			j.Progress = p
			j.UpdatedAt = time.Now()
		},
	)

	if err != nil {
		j.Status = Error
		j.Error = err.Error()
		j.UpdatedAt = time.Now()

		slog.Error("download job failed", "job_id", j.ID, "error", err)
		return
	}

	j.Status = Done
	j.FileID = fileID
	j.UpdatedAt = time.Now()

	slog.Info("download job completed", "job_id", j.ID, "file_id", fileID)
}
