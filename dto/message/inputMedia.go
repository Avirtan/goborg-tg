package message_dto

// https://core.telegram.org/bots/api#inputmedia
type InputMedia struct {
	Type                        string          `json:"type"`
	Media                       string          `json:"media"`
	Caption                     string          `json:"caption,omitempty"`
	ParseMode                   string          `json:"parse_mode,omitempty"`
	CaptionEntities             []MessageEntity `json:"caption_entities,omitempty"`
	HasSpoiler                  bool            `json:"has_spoiler,omitempty"`
	Thumbnail                   string          `json:"thumbnail,omitempty"`
	Width                       int             `json:"width,omitempty"`
	Height                      int             `json:"height,omitempty"`
	Duration                    int             `json:"duration,omitempty"`
	SupportsStreaming           bool            `json:"supports_streaming,omitempty"`
	Performer                   string          `json:"performer,omitempty"`
	Title                       string          `json:"title,omitempty"`
	DisableContentTypeDetection bool            `json:"disable_content_type_detection,omitempty"`
}
