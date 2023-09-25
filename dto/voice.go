package dto

// https://core.telegram.org/bots/api#voice
type Voice struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Duration     int64  `json:"duration"`
	MimeType     string `json:"mime_type,omitempty"`
	FileSize     string `json:"file_size,omitempty"`
}
