package dto

// https://core.telegram.org/bots/api#audio
type Audio struct {
	FileId       string    `json:"file_id"`
	FileUniqueId string    `json:"file_unique_id"`
	Duration     int       `json:"duration"`
	Performer    string    `json:"performer,omitempty"`
	Title        string    `json:"title,omitempty"`
	FileName     string    `json:"file_name,omitempty"`
	NimeType     string    `json:"mime_type,omitempty"`
	FileSize     int       `json:"file_size,omitempty"`
	Thumbnail    PhotoSize `json:"thumbnail,omitempty"`
}
