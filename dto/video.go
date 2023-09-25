package dto

// https://core.telegram.org/bots/api#video
type Video struct {
	FileId       string    `json:"file_id"`
	FileUniqueId string    `json:"file_unique_id"`
	Width        int       `json:"width"`
	Height       int       `json:"height"`
	Duration     int       `json:"duration"`
	Thumbnail    PhotoSize `json:"thumbnail,omitempty"`
	FileName     string    `json:"file_name,omitempty"`
	MimeType     string    `json:"mime_type,omitempty"`
	FileSize     string    `json:"file_size,omitempty"`
}