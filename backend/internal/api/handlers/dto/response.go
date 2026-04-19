package dto

type InfoResponse struct {
    Title          string   `json:"title"`
    Thumbnail      string   `json:"thumbnail"`
    Resolutions    []string `json:"resolutions"`
    AudioBitrates  []string `json:"audio_bitrates"`
}

type DownloadResponse struct {
    FileID      string `json:"file_id"`
    DownloadURL string `json:"download_url"`
}
