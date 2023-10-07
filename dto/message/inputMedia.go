package message_dto

import message_entity_dto "github.com/Avirtan/TGoBot/dto/message_entity"

// https://core.telegram.org/bots/api#inputmedia
type InputMedia interface {
	GetTypeMedia() string
}

// https://core.telegram.org/bots/api#inputmediaphoto
type InputMediaPhoto struct {
	Type            string                             `json:"type"`
	Media           string                             `json:"media"`
	Caption         string                             `json:"caption,omitempty"`
	ParseMode       string                             `json:"parse_mode,omitempty"`
	CaptionEntities []message_entity_dto.MessageEntity `json:"caption_entities,omitempty"`
	HasSpoiler      bool                               `json:"has_spoiler,omitempty"`
}

func (m *InputMediaPhoto) GetTypeMedia() string {
	return "photo"
}

// https://core.telegram.org/bots/api#inputmediavideo
type InputMediaVideo struct {
	Type              string                             `json:"type"`
	Media             string                             `json:"media"`
	Thumbnail         string                             `json:"thumbnail,omitempty"`
	Caption           string                             `json:"caption,omitempty"`
	ParseMode         string                             `json:"parse_mode,omitempty"`
	CaptionEntities   []message_entity_dto.MessageEntity `json:"caption_entities,omitempty"`
	Width             int                                `json:"width,omitempty"`
	Height            int                                `json:"height,omitempty"`
	Duration          int                                `json:"duration,omitempty"`
	SupportsStreaming bool                               `json:"supports_streaming,omitempty"`
	HasSpoiler        bool                               `json:"has_spoiler,omitempty"`
}

func (m *InputMediaVideo) GetTypeMedia() string {
	return "video"
}

// https://core.telegram.org/bots/api#inputmediaanimation
type InputMediaAnimation struct {
	Type            string                             `json:"type"`
	Media           string                             `json:"media"`
	Thumbnail       string                             `json:"thumbnail,omitempty"`
	Caption         string                             `json:"caption,omitempty"`
	ParseMode       string                             `json:"parse_mode,omitempty"`
	CaptionEntities []message_entity_dto.MessageEntity `json:"caption_entities,omitempty"`
	Width           int                                `json:"width,omitempty"`
	Height          int                                `json:"height,omitempty"`
	Duration        int                                `json:"duration,omitempty"`
	HasSpoiler      bool                               `json:"has_spoiler,omitempty"`
}

func (m *InputMediaAnimation) GetTypeMedia() string {
	return "animation"
}

// https://core.telegram.org/bots/api#inputmediaaudio
type InputMediaAudio struct {
	Type            string                             `json:"type"`
	Media           string                             `json:"media"`
	Thumbnail       string                             `json:"thumbnail,omitempty"`
	Caption         string                             `json:"caption,omitempty"`
	ParseMode       string                             `json:"parse_mode,omitempty"`
	CaptionEntities []message_entity_dto.MessageEntity `json:"caption_entities,omitempty"`
	Duration        int                                `json:"duration,omitempty"`
	Performer       string                             `json:"performer,omitempty"`
	Title           string                             `json:"title,omitempty"`
}

func (m *InputMediaAudio) GetTypeMedia() string {
	return "audio"
}

// https://core.telegram.org/bots/api#inputmediadocument
type InputMediaDocument struct {
	Type                        string                             `json:"type"`
	Media                       string                             `json:"media"`
	Thumbnail                   string                             `json:"thumbnail,omitempty"`
	Caption                     string                             `json:"caption,omitempty"`
	ParseMode                   string                             `json:"parse_mode,omitempty"`
	CaptionEntities             []message_entity_dto.MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool                               `json:"disable_content_type_detection,omitempty"`
}

func (m *InputMediaDocument) GetTypeMedia() string {
	return "document"
}
