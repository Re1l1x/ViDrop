package dto

type InfoRequest struct {
	URL string `json:"url"`
}

type DownloadRequest struct {
	URL          string `json:"url"`
	Resolution   int    `json:"resolution"`
	AudioBitrate int    `json:"audio_bitrate"`
}
