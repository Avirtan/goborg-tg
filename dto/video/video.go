package video_dto

import utils_dto "github.com/Avirtan/goborg-tg/dto/utils"

// https://core.telegram.org/bots/api#video
type Video struct {
	FileId       string               `json:"file_id"`
	FileUniqueId string               `json:"file_unique_id"`
	Width        int                  `json:"width"`
	Height       int                  `json:"height"`
	Duration     int                  `json:"duration"`
	Thumbnail    *utils_dto.PhotoSize `json:"thumbnail,omitempty"`
	FileName     string               `json:"file_name,omitempty"`
	MimeType     string               `json:"mime_type,omitempty"`
	FileSize     int64                `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#videonote
type VideoNote struct {
	FileId       string               `json:"file_id"`
	FileUniqueId string               `json:"file_unique_id"`
	Length       int                  `json:"length"`
	Duration     int                  `json:"duration"`
	Thumbnail    *utils_dto.PhotoSize `json:"thumbnail,omitempty"`
	FileSize     int64                `json:"file_size,omitempty"`
}
