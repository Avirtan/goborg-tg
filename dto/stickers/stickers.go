package stickers_dto

import utils_dto "github.com/Avirtan/goborg-tg/dto/utils"

// https://core.telegram.org/bots/api#stickers
type Sticker struct {
	FileID           string               `json:"file_id"`
	FileUniqueID     string               `json:"file_unique_id"`
	Type             string               `json:"type"`
	Width            int                  `json:"width"`
	Height           int                  `json:"height"`
	IsAnimated       bool                 `json:"is_animated"`
	IsVideo          bool                 `json:"is_video"`
	Thumbnail        *utils_dto.PhotoSize `json:"thumbnail,omitempty"`
	Emoji            string               `json:"emoji,omitempty"`
	SetName          string               `json:"set_name,omitempty"`
	PremiumAnimation *utils_dto.File      `json:"premium_animation,omitempty"`
	MaskPosition     *MaskPosition        `json:"mask_position,omitempty"`
	CustomEmojiID    string               `json:"custom_emoji_id,omitempty"`
	NeedsRepainting  bool                 `json:"needs_repainting,omitempty"`
	FileSize         int                  `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#stickerset
type StickerSet struct {
	Name        string               `json:"name"`
	Title       string               `json:"title"`
	StickerType string               `json:"sticker_type"`
	IsAnimated  bool                 `json:"is_animated"`
	IsVideo     bool                 `json:"is_video"`
	Stickers    []Sticker            `json:"stickers"`
	Thumbnail   *utils_dto.PhotoSize `json:"thumbnail,omitempty"`
}

// https://core.telegram.org/bots/api#maskposition
type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}

// https://core.telegram.org/bots/api#inputsticker
type InputSticker struct {
	Sticker      string        `json:"sticker"`
	EmojiList    []string      `json:"emoji_list,omitempty"`
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	Keywords     []string      `json:"keywords,omitempty"`
}
