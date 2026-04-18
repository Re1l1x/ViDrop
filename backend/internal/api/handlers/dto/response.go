package dto

type InfoResponse struct {
    Title     string `json:"title"`
	Thumbnail string `json:"thumbnail"`
}

type DownloadResponse struct {
    FileID      string `json:"file_id"`
    DownloadURL string `json:"download_url"`
}