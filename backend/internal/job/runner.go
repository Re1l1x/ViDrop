package job

import (
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

	videoID, err := r.downloader.GetVideoID(j.URL)
	if err != nil {
		j.Status = Error
		j.Error = err.Error()
		return
	}

	fileID := service.GenerateFileID(videoID, j.Resolution, j.AudioBitrate, j.Format)
	j.FileID = fileID

	if _, err := r.downloader.GetFilePath(fileID); err == nil {
		j.Status = Done
		return
	}

	fileID, err = r.downloader.Download(
		j.URL,
		j.Resolution,
		j.AudioBitrate,
		j.Format,
	)
	if err != nil {
		j.Status = Error
		j.Error = err.Error()
		return
	}

	j.Status = Done
	j.FileID = fileID
	j.UpdatedAt = time.Now()
}
