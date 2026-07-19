package job

import (
	"log/slog"
	"time"

	"ViDrop/internal/service"
)

type Runner struct {
	downloader *service.Downloader
	broker     *Broker
}

func NewRunner(d *service.Downloader, b *Broker) *Runner {
	return &Runner{
		downloader: d,
		broker:     b,
	}
}

func (r *Runner) Run(j *DownloadJob) {
	j.Status = Downloading
	j.UpdatedAt = time.Now()
	r.publish(j)

	slog.Debug(
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

			r.publish(j)
		},
	)

	if err != nil {
		j.Status = Error
		j.Error = err.Error()
		j.UpdatedAt = time.Now()

		r.publish(j)

		slog.Error("download job failed", "job_id", j.ID, "error", err)

		return
	}

	j.Status = Done
	j.Progress = 100
	j.FileID = fileID
	j.UpdatedAt = time.Now()

	r.publish(j)

	slog.Debug("download job completed", "job_id", j.ID, "file_id", fileID)
}

func (r *Runner) publish(j *DownloadJob) {
	r.broker.Publish(j.ID, ProgressEvent{
		Status:   j.Status,
		Progress: j.Progress,
		FileID:   j.FileID,
		Error:    j.Error,
	})
}
