package dto

type VideoInfoRequest struct {
	URL string `json:"url"`
}

type StartDownloadRequest struct {
	URL          string `json:"url"`
	Resolution   int    `json:"resolution"`
	AudioBitrate int    `json:"audio_bitrate"`
	Format       string `json:"format"`
}
