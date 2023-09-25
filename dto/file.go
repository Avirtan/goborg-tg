package dto

//https://core.telegram.org/bots/api#file
type File struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size,omitempty"`
	FilePath     string `json:"file_path,omitempty"`
}

// https://core.telegram.org/bots/api#photosize
type PhotoSize struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int64  `json:"file_size"`
}
