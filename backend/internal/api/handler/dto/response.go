package dto

type VideoInfoResponse struct {
	Title         string `json:"title"`
	Thumbnail     string `json:"thumbnail"`
	Resolutions   []int  `json:"resolutions"`
	AudioBitrates []int  `json:"audio_bitrates"`
}

type StartDownloadResponse struct {
	JobID string `json:"job_id"`
}

type DownloadStatusResponse struct {
	Status   string `json:"status"`
	Progress int    `json:"progress"`
	FileID   string `json:"file_id,omitempty"`
	Error    string `json:"error,omitempty"`
}

type DownloadProgressEvent struct {
	Status   string `json:"status"`
	Progress int    `json:"progress"`
	FileID   string `json:"file_id,omitempty"`
	Error    string `json:"error,omitempty"`
}
